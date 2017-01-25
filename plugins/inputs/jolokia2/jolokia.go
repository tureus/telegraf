package jolokia2

import (
	"fmt"
	"time"

	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"
)

type Jolokia struct {
	Agents                agentsConfig
	Proxy                 proxyConfig
	Metrics               []metricConfig `toml:"metric"`
	DefaultFieldPrefix    string         `toml:"default_field_prefix"`
	DefaultFieldDelimiter string         `toml:"default_field_delimiter"`
	DefaultTagPrefix      string         `toml:"default_tag_prefix"`
	DefaultTagDelimiter   string         `toml:"default_tag_delimiter"`
}

type remoteConfig struct {
	ResponseTimeout    time.Duration `toml:"response_timeout"`
	Username           string
	Password           string
	SSLCA              string `toml:"ssl_ca"`
	SSLCert            string `toml:"ssl_cert"`
	SSLKey             string `toml:"ssl_key"`
	InsecureSkipVerify bool   `toml:"insecure_skip_verify"`
}

type agentsConfig struct {
	remoteConfig
	Urls []string
}

type proxyConfig struct {
	remoteConfig
	Url                   string
	DefaultTargetPassword string `toml:"default_target_username"`
	DefaultTargetUsername string `toml:"default_target_password"`

	Targets []proxyTargetConfig
}

type proxyTargetConfig struct {
	Url      string
	Username string
	Password string
}

type metricConfig struct {
	Name           string
	Mbean          string
	Paths          []string
	FieldName      string   `toml:"field_name"`
	FieldPrefix    *string  `toml:"field_prefix"`
	FieldDelimiter *string  `toml:"field_delimiter"`
	TagPrefix      *string  `toml:"tag_prefix"`
	TagDelimiter   *string  `toml:"tag_delimiter"`
	TagInclude     []string `toml:"taginclude"`
	TagExclude     []string `toml:"tagexclude"`
}

func (jc *Jolokia) SampleConfig() string {
	return fmt.Sprintf(`
# %s

[[inputs.jolokia2]]
  # Add a metric name prefix
  #name_prefix = "example_"

  # Add agents to query
  [inputs.jolokia2.agents]
    urls     = ["http://kafka:8080/jolokia"]
    #username = ""
    #password = ""
    #ssl_ca   = "/var/private/ca.pem"
    #ssl_cert = "/var/private/client.pem"
    #ssl_key  = "/var/private/client-key.pem"
    #insecure_skip_verify = false

  [[inputs.jolokia2.metric]]
    name  = "jvm_runtime"
    mbean = "java.lang:type=Runtime"
    paths = ["Uptime"]
`, jc.Description())
}

func (jc *Jolokia) Description() string {
	return "Read JMX metrics from a Jolokia REST endpoint"
}

func (jc *Jolokia) Gather(acc telegraf.Accumulator) error {
	var metrics []Metric

	for _, config := range jc.Metrics {
		metric := Metric{
			Name:      config.Name,
			Mbean:     config.Mbean,
			Paths:     config.Paths,
			AllowTags: config.TagInclude,
			DenyTags:  config.TagExclude,
		}

		if config.FieldPrefix == nil {
			metric.FieldPrefix = jc.DefaultFieldPrefix
		} else {
			metric.FieldPrefix = *config.FieldPrefix
		}

		if config.FieldDelimiter == nil {
			metric.FieldDelimiter = jc.DefaultFieldDelimiter
		} else {
			metric.FieldDelimiter = *config.FieldDelimiter
		}

		if config.TagPrefix == nil {
			metric.TagPrefix = jc.DefaultTagPrefix
		} else {
			metric.TagPrefix = *config.TagPrefix
		}

		if config.TagDelimiter == nil {
			metric.TagDelimiter = jc.DefaultTagDelimiter
		} else {
			metric.TagDelimiter = *config.TagDelimiter
		}

		metrics = append(metrics, metric)
	}

	gatherer := NewGatherer(metrics, acc)
	requests := RequestPayload(metrics)

	// for each remote config...
	for _, url := range jc.Agents.Urls {
		agent := NewAgent(url, &jc.Agents.remoteConfig)
		tags := map[string]string{"jolokia_agent_url": agent.url}

		responses, err := agent.Read(requests)
		if err != nil {
			return err
		}

		gatherer.Gather(responses, tags)
	}

	return nil
}

func init() {
	inputs.Add("jolokia2", func() telegraf.Input {
		return &Jolokia{
			Metrics:               []metricConfig{},
			DefaultFieldPrefix:    "",
			DefaultFieldDelimiter: ".",
			DefaultTagPrefix:      "mbean",
			DefaultTagDelimiter:   "_",
		}
	})
}
