/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	clusterv1 "sigs.k8s.io/cluster-api/api/v1alpha3"
	clusterv1exp "sigs.k8s.io/cluster-api/exp/api/v1alpha3"
	"sigs.k8s.io/cluster-api/util"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/apiutil"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	infrav1 "sigs.k8s.io/cluster-api-provider-azure/api/v1alpha3"
	"sigs.k8s.io/cluster-api-provider-azure/controllers"
	infrav1exp "sigs.k8s.io/cluster-api-provider-azure/exp/api/v1alpha3"
	"sigs.k8s.io/cluster-api-provider-azure/util/reconciler"
)

// AzureClusterToAzureMachinePoolsMapper creates a mapping handler to transform AzureClusters into AzureMachinePools. The transform
// requires AzureCluster to map to the owning Cluster, then from the Cluster, collect the MachinePools belonging to the cluster,
// then finally projecting the infrastructure reference to the AzureMachinePool.
func AzureClusterToAzureMachinePoolsMapper(c client.Client, scheme *runtime.Scheme, log logr.Logger) (handler.Mapper, error) {
	gvk, err := apiutil.GVKForObject(new(infrav1exp.AzureMachinePool), scheme)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to find GVK for AzureMachinePool")
	}

	return handler.ToRequestsFunc(func(o handler.MapObject) []ctrl.Request {
		ctx, cancel := context.WithTimeout(context.Background(), reconciler.DefaultMappingTimeout)
		defer cancel()

		azCluster, ok := o.Object.(*infrav1.AzureCluster)
		if !ok {
			log.Error(errors.Errorf("expected an AzureCluster, got %T instead", o.Object), "failed to map AzureCluster")
			return nil
		}

		log = log.WithValues("AzureCluster", azCluster.Name, "Namespace", azCluster.Namespace)

		// Don't handle deleted AzureClusters
		if !azCluster.ObjectMeta.DeletionTimestamp.IsZero() {
			log.V(4).Info("AzureCluster has a deletion timestamp, skipping mapping.")
			return nil
		}

		clusterName, ok := controllers.GetOwnerClusterName(azCluster.ObjectMeta)
		if !ok {
			log.V(4).Info("unable to get the owner cluster")
			return nil
		}

		machineList := &clusterv1exp.MachinePoolList{}
		// list all of the requested objects within the cluster namespace with the cluster name label
		if err := c.List(ctx, machineList, client.InNamespace(azCluster.Namespace), client.MatchingLabels{clusterv1.ClusterLabelName: clusterName}); err != nil {
			return nil
		}

		mapFunc := MachinePoolToInfrastructureMapFunc(gvk, log)
		var results []ctrl.Request
		for _, machine := range machineList.Items {
			m := machine
			azureMachines := mapFunc.Map(handler.MapObject{
				Object: &m,
			})
			results = append(results, azureMachines...)
		}

		return results
	}), nil
}

// AzureMachinePoolMachineMapper creates a mapping handler to transform AzureMachinePoolMachine to AzureMachinePools
func AzureMachinePoolMachineMapper(scheme *runtime.Scheme, log logr.Logger) handler.ToRequestsFunc {
	return func(o handler.MapObject) []ctrl.Request {
		gvk, err := apiutil.GVKForObject(new(infrav1exp.AzureMachinePool), scheme)
		if err != nil {
			log.Error(errors.WithStack(err), "failed to find GVK for AzureMachinePool")
			return nil
		}

		azureMachinePoolMachine, ok := o.Object.(*infrav1exp.AzureMachinePoolMachine)
		if !ok {
			log.Error(errors.Errorf("expected an AzureCluster, got %T instead", o.Object), "failed to map AzureMachinePoolMachine")
			return nil
		}

		log = log.WithValues("AzureMachinePoolMachine", azureMachinePoolMachine.Name, "Namespace", azureMachinePoolMachine.Namespace)
		for _, ref := range azureMachinePoolMachine.OwnerReferences {
			if ref.Kind != gvk.Kind {
				continue
			}

			gv, err := schema.ParseGroupVersion(ref.APIVersion)
			if err != nil {
				log.Error(errors.WithStack(err), "unable to parse group version", "APIVersion", ref.APIVersion)
				return nil
			}

			if gv.Group == gvk.Group {
				return []ctrl.Request{
					{
						NamespacedName: types.NamespacedName{
							Name:      ref.Name,
							Namespace: azureMachinePoolMachine.Namespace,
						},
					},
				}
			}
		}

		return nil
	}
}

// AzureManagedClusterToAzureManagedMachinePoolsMapper creates a mapping handler to transform AzureManagedClusters into
// AzureManagedMachinePools. The transform requires AzureManagedCluster to map to the owning Cluster, then from the
// Cluster, collect the MachinePools belonging to the cluster, then finally projecting the infrastructure reference
// to the AzureManagedMachinePools.
func AzureManagedClusterToAzureManagedMachinePoolsMapper(c client.Client, scheme *runtime.Scheme, log logr.Logger) (handler.Mapper, error) {
	gvk, err := apiutil.GVKForObject(new(infrav1exp.AzureManagedMachinePool), scheme)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to find GVK for AzureManagedMachinePool")
	}

	return handler.ToRequestsFunc(func(o handler.MapObject) []ctrl.Request {
		ctx, cancel := context.WithTimeout(context.Background(), reconciler.DefaultMappingTimeout)
		defer cancel()

		azCluster, ok := o.Object.(*infrav1exp.AzureManagedCluster)
		if !ok {
			log.Error(errors.Errorf("expected an AzureManagedCluster, got %T instead", o.Object), "failed to map AzureManagedCluster")
			return nil
		}

		log = log.WithValues("AzureCluster", azCluster.Name, "Namespace", azCluster.Namespace)

		// Don't handle deleted AzureClusters
		if !azCluster.ObjectMeta.DeletionTimestamp.IsZero() {
			log.V(4).Info("AzureManagedCluster has a deletion timestamp, skipping mapping.")
			return nil
		}

		clusterName, ok := controllers.GetOwnerClusterName(azCluster.ObjectMeta)
		if !ok {
			log.V(4).Info("unable to get the owner cluster")
			return nil
		}

		machineList := &clusterv1exp.MachinePoolList{}
		// list all of the requested objects within the cluster namespace with the cluster name label
		if err := c.List(ctx, machineList, client.InNamespace(azCluster.Namespace), client.MatchingLabels{clusterv1.ClusterLabelName: clusterName}); err != nil {
			return nil
		}

		mapFunc := MachinePoolToInfrastructureMapFunc(gvk, log)
		var results []ctrl.Request
		for _, machine := range machineList.Items {
			m := machine
			azureMachines := mapFunc.Map(handler.MapObject{
				Object: &m,
			})
			results = append(results, azureMachines...)
		}

		return results
	}), nil
}

// AzureManagedClusterToAzureManagedControlPlaneMapper creates a mapping handler to transform AzureManagedClusters into
// AzureManagedControlPlane. The transform requires AzureManagedCluster to map to the owning Cluster, then from the
// Cluster, collect the control plane infrastructure reference.
func AzureManagedClusterToAzureManagedControlPlaneMapper(c client.Client, log logr.Logger) (handler.Mapper, error) {
	return handler.ToRequestsFunc(func(o handler.MapObject) []ctrl.Request {
		ctx, cancel := context.WithTimeout(context.Background(), reconciler.DefaultMappingTimeout)
		defer cancel()

		azCluster, ok := o.Object.(*infrav1exp.AzureManagedCluster)
		if !ok {
			log.Error(errors.Errorf("expected an AzureManagedCluster, got %T instead", o.Object), "failed to map AzureManagedCluster")
			return nil
		}

		log = log.WithValues("AzureCluster", azCluster.Name, "Namespace", azCluster.Namespace)

		// Don't handle deleted AzureClusters
		if !azCluster.ObjectMeta.DeletionTimestamp.IsZero() {
			log.V(4).Info("AzureManagedCluster has a deletion timestamp, skipping mapping.")
			return nil
		}

		cluster, err := util.GetOwnerCluster(ctx, c, azCluster.ObjectMeta)
		if err != nil {
			log.Error(err, "failed to get the owning cluster")
			return nil
		}

		if cluster == nil {
			log.Error(err, "cluster has not set owner ref yet")
			return nil
		}

		ref := cluster.Spec.ControlPlaneRef
		if ref == nil || ref.Name == "" {
			return nil
		}

		return []ctrl.Request{
			{
				NamespacedName: types.NamespacedName{
					Namespace: ref.Namespace,
					Name:      ref.Name,
				},
			},
		}
	}), nil
}

// MachinePoolToInfrastructureMapFunc returns a handler.ToRequestsFunc that watches for
// MachinePool events and returns reconciliation requests for an infrastructure provider object.
func MachinePoolToInfrastructureMapFunc(gvk schema.GroupVersionKind, log logr.Logger) handler.ToRequestsFunc {
	return func(o handler.MapObject) []reconcile.Request {
		m, ok := o.Object.(*clusterv1exp.MachinePool)
		if !ok {
			log.V(4).Info("attempt to map incorrect type", "type", fmt.Sprintf("%T", o.Object))
			return nil
		}

		gk := gvk.GroupKind()
		ref := m.Spec.Template.Spec.InfrastructureRef
		// Return early if the GroupKind doesn't match what we expect.
		infraGK := ref.GroupVersionKind().GroupKind()
		if gk != infraGK {
			log.V(4).Info("gk does not match", "gk", gk, "infraGK", infraGK)
			return nil
		}

		return []reconcile.Request{
			{
				NamespacedName: client.ObjectKey{
					Namespace: m.Namespace,
					Name:      ref.Name,
				},
			},
		}
	}
}

// AzureClusterToAzureMachinePoolsFunc is a handler.ToRequestsFunc to be used to enqueue
// requests for reconciliation of AzureMachinePools.
func AzureClusterToAzureMachinePoolsFunc(kClient client.Client, log logr.Logger) handler.ToRequestsFunc {
	return func(o handler.MapObject) []reconcile.Request {
		ctx, cancel := context.WithTimeout(context.Background(), reconciler.DefaultMappingTimeout)
		defer cancel()

		c, ok := o.Object.(*infrav1.AzureCluster)
		if !ok {
			log.Error(errors.Errorf("expected a AzureCluster but got a %T", o.Object), "failed to get AzureCluster")
			return nil
		}
		logWithValues := log.WithValues("AzureCluster", c.Name, "Namespace", c.Namespace)

		cluster, err := util.GetOwnerCluster(ctx, kClient, c.ObjectMeta)
		switch {
		case apierrors.IsNotFound(err) || cluster == nil:
			logWithValues.V(4).Info("owning cluster not found")
			return nil
		case err != nil:
			logWithValues.Error(err, "failed to get owning cluster")
			return nil
		}

		labels := map[string]string{clusterv1.ClusterLabelName: cluster.Name}
		ampl := &infrav1exp.AzureMachinePoolList{}
		if err := kClient.List(ctx, ampl, client.InNamespace(c.Namespace), client.MatchingLabels(labels)); err != nil {
			logWithValues.Error(err, "failed to list AzureMachinePools")
			return nil
		}

		var result []reconcile.Request
		for _, m := range ampl.Items {
			result = append(result, reconcile.Request{
				NamespacedName: client.ObjectKey{
					Namespace: m.Namespace,
					Name:      m.Name,
				},
			})
		}

		return result
	}
}

// AzureMachinePoolToAzureMachinePoolMachines maps an AzureMachinePool to it's child AzureMachinePoolMachines through
// Cluster and MachinePool labels
func AzureMachinePoolToAzureMachinePoolMachines(kClient client.Client, log logr.Logger) handler.ToRequestsFunc {
	return func(o handler.MapObject) []reconcile.Request {
		ctx, cancel := context.WithTimeout(context.Background(), reconciler.DefaultMappingTimeout)
		defer cancel()

		amp, ok := o.Object.(*infrav1exp.AzureMachinePool)
		if !ok {
			log.Error(errors.Errorf("expected a AzureMachinePool but got a %T", o.Object), "failed to get AzureMachinePool")
			return nil
		}
		logWithValues := log.WithValues("AzureMachinePool", amp.Name, "Namespace", amp.Namespace)

		labels := map[string]string{
			clusterv1.ClusterLabelName:      amp.Labels[clusterv1.ClusterLabelName],
			infrav1exp.MachinePoolNameLabel: amp.Name,
		}
		ampml := &infrav1exp.AzureMachinePoolMachineList{}
		if err := kClient.List(ctx, ampml, client.InNamespace(amp.Namespace), client.MatchingLabels(labels)); err != nil {
			logWithValues.Error(err, "failed to list AzureMachinePoolMachines")
			return nil
		}

		logWithValues.Info("mapping from AzureMachinePool", "count", len(ampml.Items))
		var result []reconcile.Request
		for _, m := range ampml.Items {
			result = append(result, reconcile.Request{
				NamespacedName: client.ObjectKey{
					Namespace: m.Namespace,
					Name:      m.Name,
				},
			})
		}

		return result
	}
}

// MachinePoolModelHasChanged predicates any events based on changes to the AzureMachinePool model
func MachinePoolModelHasChanged(logger logr.Logger) predicate.Funcs {
	return predicate.Funcs{
		UpdateFunc: func(e event.UpdateEvent) bool {
			log := logger.WithValues("predicate", "MachinePoolModelHasChanged", "eventType", "update")

			oldAmp, ok := e.ObjectOld.(*infrav1exp.AzureMachinePool)
			if !ok {

				log.V(4).Info("Expected AzureMachinePool", "type", e.ObjectOld.GetObjectKind().GroupVersionKind().String())
				return false
			}
			log = log.WithValues("namespace", oldAmp.Namespace, "azureMachinePool", oldAmp.Name)

			newAmp := e.ObjectNew.(*infrav1exp.AzureMachinePool)

			// if any of these are not equal, run the update
			shouldUpdate := !cmp.Equal(oldAmp.Spec.Identity, newAmp.Spec.Identity) ||
				!cmp.Equal(oldAmp.Spec.Template, newAmp.Spec.Template) ||
				!cmp.Equal(oldAmp.Spec.UserAssignedIdentities, newAmp.Spec.UserAssignedIdentities)

			//if shouldUpdate {
			log.Info("machine pool predicate", "shouldUpdate", shouldUpdate)
			//}
			return shouldUpdate
		},
		CreateFunc:  func(e event.CreateEvent) bool { return false },
		DeleteFunc:  func(e event.DeleteEvent) bool { return false },
		GenericFunc: func(e event.GenericEvent) bool { return false },
	}
}

// MachinePoolMachineHasStateOrVersionChange predicates any events based on changes to the AzureMachinePoolMachine status
// relevant for the AzureMachinePool controller
func MachinePoolMachineHasStateOrVersionChange(logger logr.Logger) predicate.Funcs {
	return predicate.Funcs{
		UpdateFunc: func(e event.UpdateEvent) bool {
			log := logger.WithValues("predicate", "MachinePoolModelHasChanged", "eventType", "update")

			oldAmp, ok := e.ObjectOld.(*infrav1exp.AzureMachinePoolMachine)
			if !ok {

				log.V(4).Info("Expected AzureMachinePoolMachine", "type", e.ObjectOld.GetObjectKind().GroupVersionKind().String())
				return false
			}
			log = log.WithValues("namespace", oldAmp.Namespace, "machinePoolMachine", oldAmp.Name)

			newAmp := e.ObjectNew.(*infrav1exp.AzureMachinePoolMachine)

			// if any of these are not equal, run the update
			shouldUpdate := oldAmp.Status.LatestModelApplied != newAmp.Status.LatestModelApplied ||
				oldAmp.Status.Version != newAmp.Status.Version ||
				oldAmp.Status.ProvisioningState != newAmp.Status.ProvisioningState ||
				oldAmp.Status.Ready != newAmp.Status.Ready

			if shouldUpdate {
				log.Info("machine pool machine predicate", "shouldUpdate", shouldUpdate)
			}
			return shouldUpdate
		},
		CreateFunc:  func(e event.CreateEvent) bool { return false },
		DeleteFunc:  func(e event.DeleteEvent) bool { return false },
		GenericFunc: func(e event.GenericEvent) bool { return false },
	}
}
