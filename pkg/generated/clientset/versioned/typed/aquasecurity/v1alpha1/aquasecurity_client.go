// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"net/http"

	v1alpha1 "github.com/aquasecurity/trivy-operator/pkg/apis/aquasecurity/v1alpha1"
	"github.com/aquasecurity/trivy-operator/pkg/generated/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type AquasecurityV1alpha1Interface interface {
	RESTClient() rest.Interface
	ClusterComplianceDetailReportsGetter
	ClusterComplianceReportsGetter
	ClusterConfigAuditReportsGetter
	ClusterVulnerabilityReportsGetter
	ConfigAuditReportsGetter
	VulnerabilityReportsGetter
}

// AquasecurityV1alpha1Client is used to interact with features provided by the aquasecurity.github.io group.
type AquasecurityV1alpha1Client struct {
	restClient rest.Interface
}

func (c *AquasecurityV1alpha1Client) ClusterComplianceDetailReports(namespace string) ClusterComplianceDetailReportInterface {
	return newClusterComplianceDetailReports(c, namespace)
}

func (c *AquasecurityV1alpha1Client) ClusterComplianceReports(namespace string) ClusterComplianceReportInterface {
	return newClusterComplianceReports(c, namespace)
}

func (c *AquasecurityV1alpha1Client) ClusterConfigAuditReports() ClusterConfigAuditReportInterface {
	return newClusterConfigAuditReports(c)
}

func (c *AquasecurityV1alpha1Client) ClusterVulnerabilityReports() ClusterVulnerabilityReportInterface {
	return newClusterVulnerabilityReports(c)
}

func (c *AquasecurityV1alpha1Client) ConfigAuditReports(namespace string) ConfigAuditReportInterface {
	return newConfigAuditReports(c, namespace)
}

func (c *AquasecurityV1alpha1Client) VulnerabilityReports(namespace string) VulnerabilityReportInterface {
	return newVulnerabilityReports(c, namespace)
}

// NewForConfig creates a new AquasecurityV1alpha1Client for the given config.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*AquasecurityV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	httpClient, err := rest.HTTPClientFor(&config)
	if err != nil {
		return nil, err
	}
	return NewForConfigAndClient(&config, httpClient)
}

// NewForConfigAndClient creates a new AquasecurityV1alpha1Client for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
func NewForConfigAndClient(c *rest.Config, h *http.Client) (*AquasecurityV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientForConfigAndClient(&config, h)
	if err != nil {
		return nil, err
	}
	return &AquasecurityV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new AquasecurityV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *AquasecurityV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new AquasecurityV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *AquasecurityV1alpha1Client {
	return &AquasecurityV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *AquasecurityV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
