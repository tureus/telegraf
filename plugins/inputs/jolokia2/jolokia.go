package jolokia2

import (
	"time"

	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"
)

type Jolokia struct {
	Agents                agentsConfig
	Proxy                 proxyConfig
	Metrics               []metricConfig `toml:"metric"`
	DefaultFieldPrefix    string         `toml:"default_field_prefix"`
	DefaultFieldSeparator string         `toml:"default_field_separator"`
	DefaultTagPrefix      string         `toml:"default_tag_prefix"`
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
	URLs []string `toml:"urls"`
}

type proxyConfig struct {
	remoteConfig
	URL                   string `toml:"url"`
	DefaultTargetPassword string `toml:"default_target_username"`
	DefaultTargetUsername string `toml:"default_target_password"`

	Targets []proxyTargetConfig `toml:"target"`
}

type proxyTargetConfig struct {
	URL      string `toml:"url"`
	Username string
	Password string
}

type metricConfig struct {
	Name           string
	Mbean          string
	Paths          []string
	FieldName      *string  `toml:"field_name"`
	FieldPrefix    *string  `toml:"field_prefix"`
	FieldSeparator *string  `toml:"field_separator"`
	TagPrefix      *string  `toml:"tag_prefix"`
	TagKeys        []string `toml:"tag_keys"`
}

func (jc *Jolokia) SampleConfig() string {
	return `
  # default_tag_prefix      = ""
  # default_tag_separator   = "_"
  # default_field_separator = "."

  # Add agents to query
  [inputs.jolokia2.agents]
    urls     = ["http://127.0.0.1:8080/jolokia"]
    #username = ""
    #password = ""
    #ssl_ca   = "/var/private/ca.pem"
    #ssl_cert = "/var/private/client.pem"
    #ssl_key  = "/var/private/client-key.pem"

  [[inputs.jolokia2.metric]]
    name  = "jvm_runtime"
    mbean = "java.lang:type=Runtime"
    paths = ["Uptime"]
`
}

func (jc *Jolokia) Description() string {
	return "Read JMX metrics from a Jolokia REST endpoint"
}

func (jc *Jolokia) Gather(acc telegraf.Accumulator) error {
	var metrics []Metric

	for _, config := range jc.Metrics {
		metrics = append(metrics, jc.newMetric(config))
	}

	gatherer := NewGatherer(metrics, acc)
	requests := RequestPayload(metrics)

	for _, url := range jc.Agents.URLs {
		agent := NewAgent(url, &AgentConfig{
			Username:        jc.Agents.Username,
			Password:        jc.Agents.Password,
			ResponseTimeout: jc.Agents.ResponseTimeout,
		})

		responses, err := agent.Read(requests)
		if err != nil {
			return err
		}

		tags := map[string]string{"jolokia_agent_url": url}
		gatherer.Gather(responses, tags)
	}

	if jc.Proxy.URL != "" {
		proxyConfig := &ProxyConfig{
			DefaultTargetUsername: jc.Proxy.DefaultTargetUsername,
			DefaultTargetPassword: jc.Proxy.DefaultTargetPassword,
		}

		for _, target := range jc.Proxy.Targets {
			proxyConfig.Targets = append(proxyConfig.Targets, ProxyTargetConfig{
				URL:      target.URL,
				Username: target.Username,
				Password: target.Password,
			})
		}

		agent := NewAgent(jc.Proxy.URL, &AgentConfig{
			Username:        jc.Proxy.Username,
			Password:        jc.Proxy.Password,
			ResponseTimeout: jc.Proxy.ResponseTimeout,
			ProxyConfig:     proxyConfig,
		})

		responses, err := agent.Read(requests)
		if err != nil {
			return err
		}

		tags := map[string]string{"jolokia_proxy_url": jc.Proxy.URL}
		gatherer.Gather(responses, tags)
	}

	return nil
}

func (jc *Jolokia) newMetric(config metricConfig) Metric {
	metric := Metric{
		Name:    config.Name,
		Mbean:   config.Mbean,
		Paths:   config.Paths,
		TagKeys: config.TagKeys,
	}

	if config.FieldName != nil {
		metric.FieldName = *config.FieldName
	}

	if config.FieldPrefix == nil {
		metric.FieldPrefix = jc.DefaultFieldPrefix
	} else {
		metric.FieldPrefix = *config.FieldPrefix
	}

	if config.FieldSeparator == nil {
		metric.FieldSeparator = jc.DefaultFieldSeparator
	} else {
		metric.FieldSeparator = *config.FieldSeparator
	}

	if config.TagPrefix == nil {
		metric.TagPrefix = jc.DefaultTagPrefix
	} else {
		metric.TagPrefix = *config.TagPrefix
	}

	return metric
}

func init() {
	inputs.Add("jolokia2", func() telegraf.Input {
		return &Jolokia{
			Metrics:               []metricConfig{},
			DefaultFieldPrefix:    "",
			DefaultFieldSeparator: ".",
			DefaultTagPrefix:      "",
		}
	})
}
