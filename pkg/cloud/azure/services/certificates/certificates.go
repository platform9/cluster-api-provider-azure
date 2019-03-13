/*
Copyright 2019 The Kubernetes Authors.

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

package certificates

import (
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	clientcertutil "k8s.io/client-go/util/cert"
	bootstraputil "k8s.io/cluster-bootstrap/token/util"
	"k8s.io/klog"
	kubeadmapi "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm"
	kubeadmscheme "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/scheme"
	kubeadmv1beta1 "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/v1beta1"
	kubeadmconstants "k8s.io/kubernetes/cmd/kubeadm/app/constants"
	tokenphase "k8s.io/kubernetes/cmd/kubeadm/app/phases/bootstraptoken/node"
	certsphase "k8s.io/kubernetes/cmd/kubeadm/app/phases/certs"
	kubeconfigphase "k8s.io/kubernetes/cmd/kubeadm/app/phases/kubeconfig"
	kubeconfigutil "k8s.io/kubernetes/cmd/kubeadm/app/util/kubeconfig"
	"k8s.io/kubernetes/cmd/kubeadm/app/util/pubkeypin"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/apis/azureprovider/v1alpha1"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

// ReconcileCertificates generate certificates if none exists.
func (s *Service) ReconcileCertificates() error {
	klog.V(2).Infof("Reconciling certificates for cluster %q", s.scope.Cluster.Name)

	if err := CreateOrUpdateCertificates(s.scope.ClusterConfig, s.scope.Cluster.Name); err != nil {
		return err
	}

	return nil
}

// CreateOrUpdateCertificates Helper function so this can be unittested
func CreateOrUpdateCertificates(clusterConfig *v1alpha1.AzureClusterProviderSpec, clusterName string) error {
	tmpDirName := "/tmp/cluster-api/" + clusterName
	dnsName := fmt.Sprintf("%s-api", clusterName)

	defer os.RemoveAll(tmpDirName)

	v1beta1cfg := &kubeadmv1beta1.InitConfiguration{}
	kubeadmscheme.Scheme.Default(v1beta1cfg)
	v1beta1cfg.CertificatesDir = tmpDirName + "/certs"
	v1beta1cfg.Etcd.Local = &kubeadmv1beta1.LocalEtcd{}
	// 10.0.0.1 is fake api server address, since this is also generated on masters
	v1beta1cfg.LocalAPIEndpoint = kubeadmv1beta1.APIEndpoint{AdvertiseAddress: "10.0.0.1", BindPort: 6443}
	v1beta1cfg.ControlPlaneEndpoint = fmt.Sprintf("%s:6443", dnsName)
	v1beta1cfg.APIServer.CertSANs = []string{}
	// require a fake node name for now, this will be regenerated when it runs on node anyways
	v1beta1cfg.NodeRegistration.Name = "fakenode" + clusterName
	cfg := &kubeadmapi.InitConfiguration{}
	kubeadmscheme.Scheme.Default(cfg)
	kubeadmscheme.Scheme.Convert(v1beta1cfg, cfg, nil)

	if err := CreatePKICertificates(cfg); err != nil {
		return errors.Wrapf(err, "Failed to generate pki certs: %q", err)
	}

	if err := CreateSACertificates(cfg); err != nil {
		return errors.Wrapf(err, "Failed to generate sa certs: %q", err)
	}

	kubeConfigDir := tmpDirName + "/kubeconfigs"
	if err := CreateKubeconfigs(cfg, kubeConfigDir); err != nil {
		return errors.Wrapf(err, "Failed to generate kubeconfigs: %q", err)
	}

	if err := updateClusterConfigKeyPairs(clusterConfig, tmpDirName); err != nil {
		return errors.Wrapf(err, "Failed to update certificates: %q", err)
	}

	return nil
}

// CreatePKICertificates creates base pki assets in cfg.CertDir directory
func CreatePKICertificates(cfg *kubeadmapi.InitConfiguration) error {
	if err := certsphase.CreatePKIAssets(cfg); err != nil {
		return err
	}
	return nil
}

// CreateSACertificates creates sa certificates in cfg.CertDir directory
func CreateSACertificates(cfg *kubeadmapi.InitConfiguration) error {
	if err := certsphase.CreateServiceAccountKeyAndPublicKeyFiles(cfg); err != nil {
		return err
	}

	return nil
}

// GetDiscoveryHashes returns discovery hashes from a given kubeconfig file
func GetDiscoveryHashes(kubeConfigFile string) ([]string, error) {
	// load the kubeconfig file to get the CA certificate and endpoint
	config, err := clientcmd.LoadFromFile(kubeConfigFile)
	if err != nil {
		return nil, err
	}

	// load the default cluster config
	clusterConfig := kubeconfigutil.GetClusterFromKubeConfig(config)
	if clusterConfig == nil {
		return nil, fmt.Errorf("failed to get default cluster config")
	}

	// load CA certificates from the kubeconfig (either from PEM data or by file path)
	var caCerts []*x509.Certificate
	if clusterConfig.CertificateAuthorityData != nil {
		caCerts, err = clientcertutil.ParseCertsPEM(clusterConfig.CertificateAuthorityData)
		if err != nil {
			return nil, err
		}
	} else if clusterConfig.CertificateAuthority != "" {
		caCerts, err = clientcertutil.CertsFromFile(clusterConfig.CertificateAuthority)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("no CA certificates found in kubeconfig")
	}

	// hash all the CA certs and include their public key pins as trusted values
	publicKeyPins := make([]string, 0, len(caCerts))
	for _, caCert := range caCerts {
		publicKeyPins = append(publicKeyPins, pubkeypin.Hash(caCert))
	}
	return publicKeyPins, nil
}

func CreateNewBootstrapToken() (string, error) {
	token, err := bootstraputil.GenerateBootstrapToken()
	if err != nil {
		return token, err
	}

	cfg, err := config.GetConfig()
	if err != nil {
		return token, err
	}

	kclientset, err := clientset.NewForConfig(cfg)
	if err != nil {
		return token, err
	}

	tokenString, err := kubeadmapi.NewBootstrapTokenString(token)
	if err != nil {
		return token, err
	}

	bootstrapTokens := []kubeadmapi.BootstrapToken{
		kubeadmapi.BootstrapToken{
			Token:  tokenString,
			TTL:    &metav1.Duration{Duration: 1 * time.Hour},
			Groups: []string{"system:bootstrappers:kubeadm:default-node-token"},
			Usages: []string{"signing", "authentication"},
		},
	}

	if err := tokenphase.CreateNewTokens(kclientset, bootstrapTokens); err != nil {
		return token, err
	}

	return token, nil
}

// CreateKubeconfigs creates kubeconfigs for all profiles
func CreateKubeconfigs(cfg *kubeadmapi.InitConfiguration, kubeConfigDir string) error {
	if err := kubeconfigphase.CreateKubeConfigFile(kubeadmconstants.AdminKubeConfigFileName, kubeConfigDir, cfg); err != nil {
		return err
	}
	// if err := kubeconfigphase.CreateKubeConfigFile(kubeadmconstants.KubeletKubeConfigFileName, kubeConfigDir, cfg); err != nil {
	// 	return err
	// }
	// if err := kubeconfigphase.CreateKubeConfigFile(kubeadmconstants.ControllerManagerKubeConfigFileName, kubeConfigDir, cfg); err != nil {
	// 	return err
	// }
	// if err := kubeconfigphase.CreateKubeConfigFile(kubeadmconstants.SchedulerKubeConfigFileName, kubeConfigDir, cfg); err != nil {
	// 	return err
	// }
	return nil
}

// updateClusterConfigKeyPairs populates clusterConfig with all the requisite certs
func updateClusterConfigKeyPairs(clusterConfig *v1alpha1.AzureClusterProviderSpec, tmpDirName string) error {
	certsDir := tmpDirName + "/certs"
	if len(clusterConfig.CAKeyPair.Cert) <= 0 {
		buf, err := ioutil.ReadFile(certsDir + "/ca.crt")
		if err != nil {
			return err
		}
		clusterConfig.CAKeyPair.Cert = buf
	}
	if len(clusterConfig.CAKeyPair.Key) <= 0 {
		buf, err := ioutil.ReadFile(certsDir + "/ca.key")
		if err != nil {
			return err
		}
		clusterConfig.CAKeyPair.Key = buf
	}
	if len(clusterConfig.SAKeyPair.Key) <= 0 {
		buf, err := ioutil.ReadFile(certsDir + "/sa.key")
		if err != nil {
			return err
		}
		clusterConfig.SAKeyPair.Key = buf
	}
	if len(clusterConfig.SAKeyPair.Cert) <= 0 {
		buf, err := ioutil.ReadFile(certsDir + "/sa.pub")
		if err != nil {
			return err
		}
		clusterConfig.SAKeyPair.Cert = buf
	}
	if len(clusterConfig.FrontProxyCAKeyPair.Cert) <= 0 {
		buf, err := ioutil.ReadFile(certsDir + "/front-proxy-ca.crt")
		if err != nil {
			return err
		}
		clusterConfig.FrontProxyCAKeyPair.Cert = buf
	}

	if len(clusterConfig.FrontProxyCAKeyPair.Key) <= 0 {
		buf, err := ioutil.ReadFile(certsDir + "/front-proxy-ca.key")
		if err != nil {
			return err
		}
		clusterConfig.FrontProxyCAKeyPair.Key = buf
	}

	if len(clusterConfig.EtcdCAKeyPair.Cert) <= 0 {
		buf, err := ioutil.ReadFile(certsDir + "/etcd/ca.crt")
		if err != nil {
			return err
		}
		clusterConfig.EtcdCAKeyPair.Cert = buf
	}

	if len(clusterConfig.EtcdCAKeyPair.Key) <= 0 {
		buf, err := ioutil.ReadFile(certsDir + "/etcd/ca.key")
		if err != nil {
			return err
		}
		clusterConfig.EtcdCAKeyPair.Key = buf
	}

	kubeConfigsDir := tmpDirName + "/kubeconfigs"

	if len(clusterConfig.AdminKubeconfig) <= 0 {
		buf, err := ioutil.ReadFile(kubeConfigsDir + "/admin.conf")
		if err != nil {
			return err
		}
		clusterConfig.AdminKubeconfig = string(buf)
	}

	// // Discovery hashes typically never changes
	if len(clusterConfig.DiscoveryHashes) <= 0 {
		discoveryHashes, err := GetDiscoveryHashes(kubeConfigsDir + "/admin.conf")
		if err != nil {
			return err
		}
		clusterConfig.DiscoveryHashes = discoveryHashes
	}
	return nil
}
