package config

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/kelseyhightower/envconfig"
)

// Config holds global test configuration
type Config struct {
	Timeout           float64  `envconfig:"TIMEOUT" default:"600"`
	SkipTest          bool     `envconfig:"SKIP_TEST" default:"false"`
	ClusterName       string   `envconfig:"CLUSTER_NAME"`                                                   // ClusterName allows you to set the name of a cluster already created
	Location          string   `envconfig:"LOCATION"`                                                       // Location where you want to create the cluster
	Regions           []string `envconfig:"REGIONS"`                                                        // A whitelist of available regions
	ClusterConfigPath string   `envconfig:"CLUSTERCONFIG_PATH" default:"config/base/v1alpha1_cluster.yaml"` // path to the YAML for the cluster we're creating
	MachineConfigPath string   `envconfig:"MACHINECONFIG_PATH" default:"config/base/v1alpha1_machine.yaml"` // path to the YAML describing the machines we're creating
	CleanUpOnExit     bool     `envconfig:"CLEANUP_ON_EXIT" default:"true"`                                 // if true the tests will clean up rgs when tests finish
	// CleanUpIfFail     bool     `envconfig:"CLEANUP_IF_FAIL" default:"true"`
	GinkgoFocus       string `envconfig:"GINKGO_FOCUS"`
	GinkgoSkip        string `envconfig:"GINKGO_SKIP"`
	ClientID          string `envconfig:"CLIENT_ID" required:"true"`
	ClientSecret      string `envconfig:"CLIENT_SECRET" required:"true"`
	PublicSSHKey      string `envconfig:"PUBLIC_SSH_KEY"`
	SubscriptionID    string `envconfig:"SUBSCRIPTION_ID" required:"true"`
	TenantID          string `envconfig:"TENANT_ID" required:"true"`
	KubernetesVersion string `envconfig:"KUBERNETES_VERSION" required:"true"`
}

// ParseConfig will parse needed environment variables for running the tests
func ParseConfig() (*Config, error) {
	c := new(Config)
	if err := envconfig.Process("config", c); err != nil {
		return nil, err
	}
	if c.Location == "" {
		c.SetRandomRegion()
	}
	return c, nil
}

// SetRandomRegion sets Location to a random region
func (c *Config) SetRandomRegion() {
	var regions []string
	if c.Regions == nil || len(c.Regions) == 0 {
		regions = []string{"eastus", "southcentralus", "southeastasia", "westus2", "westeurope"}
	} else {
		regions = c.Regions
	}
	log.Printf("Picking Random Region from list %s\n", regions)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	c.Location = regions[r.Intn(len(regions))]
	os.Setenv("LOCATION", c.Location)
	log.Printf("Picked Random Region:%s\n", c.Location)
}
