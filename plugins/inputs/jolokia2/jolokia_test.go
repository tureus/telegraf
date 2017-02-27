package jolokia2

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/influxdata/telegraf/internal/config"
	"github.com/influxdata/telegraf/testutil"
	"github.com/stretchr/testify/assert"
)

func TestJolokia2_ScalarValues(t *testing.T) {
	config := `
	[[inputs.jolokia2]]
		[inputs.jolokia2.agents]
			urls = ["%s"]

		[[inputs.jolokia2.metric]]
		  name  = "scalar_without_attribute"
		  mbean = "scalar_without_attribute"

		[[inputs.jolokia2.metric]]
		  name  = "scalar_with_attribute"
		  mbean = "scalar_with_attribute"
		  paths = ["biz"]

		[[inputs.jolokia2.metric]]
		  name  = "scalar_with_attribute_and_path"
		  mbean = "scalar_with_attribute_and_path"
		  paths = ["biz/baz"]

		# This should return multiple series with different test tags.
		[[inputs.jolokia2.metric]]
		  name     = "scalar_with_key_pattern"
		  mbean    = "scalar_with_key_pattern:test=*"
		  tag_keys = ["test"]`

	response := `[{
		"request": {
			"mbean": "scalar_without_attribute",
			"type": "read"
		},
		"value": 123,
		"status": 200
	  }, {
		"request": {
			"mbean": "scalar_with_attribute",
			"attribute": "biz",
			"type": "read"
		},
		"value": 456,
		"status": 200
	  }, {
		"request": {
			"mbean": "scalar_with_attribute_and_path",
			"attribute": "biz",
			"path": "baz",
			"type": "read"
		},
		"value": 789,
		"status": 200
	  }, {
		"request": {
			"mbean": "scalar_with_key_pattern:test=*",
			"type": "read"
		},
		"value": {
			"scalar_with_key_pattern:test=foo": 123,
			"scalar_with_key_pattern:test=bar": 456
		},
		"status": 200
	  }]`

	server := setupServer(http.StatusOK, response)
	defer server.Close()

	jolokia, err := setupPlugin(fmt.Sprintf(config, server.URL))
	if err != nil {
		t.Fatalf("Could not setup plugin. %v", err)
	}

	var acc testutil.Accumulator
	assert.NoError(t, jolokia.Gather(&acc))

	acc.AssertContainsTaggedFields(t, "scalar_without_attribute", map[string]interface{}{
		"value": 123.0,
	}, map[string]string{
		"jolokia_agent_url": server.URL,
	})

	acc.AssertContainsTaggedFields(t, "scalar_with_attribute", map[string]interface{}{
		"biz": 456.0,
	}, map[string]string{
		"jolokia_agent_url": server.URL,
	})

	acc.AssertContainsTaggedFields(t, "scalar_with_attribute_and_path", map[string]interface{}{
		"biz.baz": 789.0,
	}, map[string]string{
		"jolokia_agent_url": server.URL,
	})

	acc.AssertContainsTaggedFields(t, "scalar_with_key_pattern", map[string]interface{}{
		"value": 123.0,
	}, map[string]string{
		"jolokia_agent_url": server.URL,
		"test":              "foo",
	})
	acc.AssertContainsTaggedFields(t, "scalar_with_key_pattern", map[string]interface{}{
		"value": 456.0,
	}, map[string]string{
		"jolokia_agent_url": server.URL,
		"test":              "bar",
	})
}

func TestJolokia2_ObjectValues(t *testing.T) {
	config := `
	[[inputs.jolokia2]]
		[inputs.jolokia2.agents]
			urls = ["%s"]

		[[inputs.jolokia2.metric]]
			name     = "object_without_attribute"
			mbean    = "object_without_attribute"
			tag_keys = ["foo"]

		[[inputs.jolokia2.metric]]
			name  = "object_with_attribute"
			mbean = "object_with_attribute"
			paths = ["biz"]

		[[inputs.jolokia2.metric]]
			name  = "object_with_attribute_and_path"
			mbean = "object_with_attribute_and_path"
			paths = ["biz/baz"]

		# This will generate two separate request objects.
		[[inputs.jolokia2.metric]]
			name  = "object_with_branching_paths"
			mbean = "object_with_branching_paths"
			paths = ["foo/fiz", "foo/faz"]

		# This should return multiple series with different test tags.
		[[inputs.jolokia2.metric]]
			name     = "object_with_key_pattern"
			mbean    = "object_with_key_pattern:test=*"
			tag_keys = ["test"]`

	response := `[{
		"request": {
			"mbean": "object_without_attribute",
			"type": "read"
		},
		"value": {
			"biz": 123,
			"baz": 456
		},
		"status": 200
	}, {
		"request": {
			"mbean": "object_with_attribute",
			"attribute": "biz",
			"type": "read"
		},
		"value": {
			"fiz": 123,
			"faz": 456
		},
		"status": 200
	}, {
		"request": {
			"mbean": "object_with_branching_paths",
			"attribute": "foo",
			"path": "fiz",
			"type": "read"
		},
		"value": {
			"bing": 123
		},
		"status": 200
	}, {
		"request": {
			"mbean": "object_with_branching_paths",
			"attribute": "foo",
			"path": "faz",
			"type": "read"
		},
		"value": {
			"bang": 456
		},
		"status": 200
	}, {
		"request": {
			"mbean": "object_with_attribute_and_path",
			"attribute": "biz",
			"path": "baz",
			"type": "read"
		},
		"value": {
			"bing": 123,
			"bang": 456
		},
		"status": 200
	}, {
		"request": {
			"mbean": "object_with_key_pattern:test=*",
			"type": "read"
		},
		"value": {
			"object_with_key_pattern:test=foo": {
				"fiz": 123
			},
			"object_with_key_pattern:test=bar": {
				"biz": 456
			}
		},
		"status": 200
	}]`

	server := setupServer(http.StatusOK, response)
	defer server.Close()

	jolokia, err := setupPlugin(fmt.Sprintf(config, server.URL))
	if err != nil {
		t.Fatalf("Could not setup plugin. %v", err)
	}

	var acc testutil.Accumulator
	assert.NoError(t, jolokia.Gather(&acc))

	acc.AssertContainsTaggedFields(t, "object_without_attribute", map[string]interface{}{
		"biz": 123.0,
		"baz": 456.0,
	}, map[string]string{
		"jolokia_agent_url": server.URL,
	})

	acc.AssertContainsTaggedFields(t, "object_with_attribute", map[string]interface{}{
		"biz.fiz": 123.0,
		"biz.faz": 456.0,
	}, map[string]string{
		"jolokia_agent_url": server.URL,
	})

	acc.AssertContainsTaggedFields(t, "object_with_attribute_and_path", map[string]interface{}{
		"biz.baz.bing": 123.0,
		"biz.baz.bang": 456.0,
	}, map[string]string{
		"jolokia_agent_url": server.URL,
	})

	acc.AssertContainsTaggedFields(t, "object_with_branching_paths", map[string]interface{}{
		"foo.fiz.bing": 123.0,
		"foo.faz.bang": 456.0,
	}, map[string]string{
		"jolokia_agent_url": server.URL,
	})

	acc.AssertContainsTaggedFields(t, "object_with_key_pattern", map[string]interface{}{
		"fiz": 123.0,
	}, map[string]string{
		"test":              "foo",
		"jolokia_agent_url": server.URL,
	})

	acc.AssertContainsTaggedFields(t, "object_with_key_pattern", map[string]interface{}{
		"biz": 456.0,
	}, map[string]string{
		"test":              "bar",
		"jolokia_agent_url": server.URL,
	})
}

func TestJolokia2_TagRenaming(t *testing.T) {
	config := `
	[[inputs.jolokia2]]
		default_tag_prefix      = "DEFAULT_PREFIX_"

		[inputs.jolokia2.agents]
			urls = ["%s"]

		[[inputs.jolokia2.metric]]
			name     = "default_tag_prefix"
			mbean    = "default_tag_prefix:biz=baz,fiz=faz"
			tag_keys = ["biz", "fiz"]

		[[inputs.jolokia2.metric]]
			name       = "custom_tag_prefix"
			mbean      = "custom_tag_prefix:biz=baz,fiz=faz"
			tag_keys   = ["biz", "fiz"]
			tag_prefix = "CUSTOM_PREFIX_"`

	response := `[{
		"request": {
			"mbean": "default_tag_prefix:biz=baz,fiz=faz",
			"type": "read"
		},
		"value": 123,
		"status": 200
	}, {
		"request": {
			"mbean": "custom_tag_prefix:biz=baz,fiz=faz",
			"type": "read"
		},
		"value": 123,
		"status": 200
	}]`

	server := setupServer(http.StatusOK, response)
	defer server.Close()

	jolokia, err := setupPlugin(fmt.Sprintf(config, server.URL))
	if err != nil {
		t.Fatalf("Could not setup plugin. %v", err)
	}

	var acc testutil.Accumulator
	assert.NoError(t, jolokia.Gather(&acc))

	acc.AssertContainsTaggedFields(t, "default_tag_prefix", map[string]interface{}{
		"value": 123.0,
	}, map[string]string{
		"DEFAULT_PREFIX_biz": "baz",
		"DEFAULT_PREFIX_fiz": "faz",
		"jolokia_agent_url":  server.URL,
	})

	acc.AssertContainsTaggedFields(t, "custom_tag_prefix", map[string]interface{}{
		"value": 123.0,
	}, map[string]string{
		"CUSTOM_PREFIX_biz": "baz",
		"CUSTOM_PREFIX_fiz": "faz",
		"jolokia_agent_url": server.URL,
	})
}

func TestJolokia2_FieldRenaming(t *testing.T) {
	config := `
	[[inputs.jolokia2]]
		default_field_prefix    = "DEFAULT_PREFIX_"
		default_field_separator = "_DEFAULT_SEPARATOR_"

		[inputs.jolokia2.agents]
			urls = ["%s"]

		[[inputs.jolokia2.metric]]
			name  = "default_field_modifiers"
			mbean = "default_field_modifiers"

		[[inputs.jolokia2.metric]]
			name            = "custom_field_modifiers"
			mbean           = "custom_field_modifiers"
			field_prefix    = "CUSTOM_PREFIX_"
			field_separator = "_CUSTOM_SEPARATOR_"

		[[inputs.jolokia2.metric]]
			name            = "field_prefix_substitution"
			mbean           = "field_prefix_substitution:foo=*"
			field_prefix    = "$1_"

		[[inputs.jolokia2.metric]]
			name         = "field_name_substitution"
			mbean        = "field_name_substitution:foo=*"
			field_prefix = ""
			field_name   = "$1"`

	response := `[{
		"request": {
			"mbean": "default_field_modifiers",
			"type": "read"
		},
		"value": {
			"hello": { "world": 123 }
		},
		"status": 200
	}, {
		"request": {
			"mbean": "custom_field_modifiers",
			"type": "read"
		},
		"value": {
			"hello": { "world": 123 }
		},
		"status": 200
	}, {
		"request": {
			"mbean": "field_prefix_substitution:foo=*",
			"type": "read"
		},
		"value": {
			"field_prefix_substitution:foo=biz": 123,
			"field_prefix_substitution:foo=baz": 456
		},
		"status": 200
	}, {
		"request": {
			"mbean": "field_name_substitution:foo=*",
			"type": "read"
		},
		"value": {
			"field_name_substitution:foo=biz": 123,
			"field_name_substitution:foo=baz": 456
		},
		"status": 200
	}]`

	server := setupServer(http.StatusOK, response)
	defer server.Close()

	jolokia, err := setupPlugin(fmt.Sprintf(config, server.URL))
	if err != nil {
		t.Fatalf("Could not setup plugin. %v", err)
	}

	var acc testutil.Accumulator
	assert.NoError(t, jolokia.Gather(&acc))

	acc.AssertContainsTaggedFields(t, "default_field_modifiers", map[string]interface{}{
		"DEFAULT_PREFIX_hello_DEFAULT_SEPARATOR_world": 123.0,
	}, map[string]string{
		"jolokia_agent_url": server.URL,
	})

	acc.AssertContainsTaggedFields(t, "custom_field_modifiers", map[string]interface{}{
		"CUSTOM_PREFIX_hello_CUSTOM_SEPARATOR_world": 123.0,
	}, map[string]string{
		"jolokia_agent_url": server.URL,
	})

	acc.AssertContainsTaggedFields(t, "field_prefix_substitution", map[string]interface{}{
		"biz_value": 123.0,
		"baz_value": 456.0,
	}, map[string]string{
		"jolokia_agent_url": server.URL,
	})

	acc.AssertContainsTaggedFields(t, "field_name_substitution", map[string]interface{}{
		"biz": 123.0,
		"baz": 456.0,
	}, map[string]string{
		"jolokia_agent_url": server.URL,
	})
}

func TestJolokia2_MetricCompaction(t *testing.T) {
	config := `
	[[inputs.jolokia2]]
		[inputs.jolokia2.agents]
			urls = ["%s"]
		[[inputs.jolokia2.metric]]
			name     = "compact_metric"
			mbean    = "scalar_value:flavor=chocolate"
			tag_keys = ["flavor"]

		[[inputs.jolokia2.metric]]
			name     = "compact_metric"
			mbean    = "scalar_value:flavor=vanilla"
			tag_keys = ["flavor"]

		[[inputs.jolokia2.metric]]
			name     = "compact_metric"
			mbean    = "object_value1:flavor=chocolate"
			tag_keys = ["flavor"]

		[[inputs.jolokia2.metric]]
			name     = "compact_metric"
			mbean    = "object_value2:flavor=chocolate"
			tag_keys = ["flavor"]`

	response := `[{
		"request": {
			"mbean": "scalar_value:flavor=chocolate",
			"type": "read"
		},
		"value": 123,
		"status": 200
	}, {
		"request": {
			"mbean": "scalar_value:flavor=vanilla",
			"type": "read"
		},
		"value": 999,
		"status": 200
	}, {
		"request": {
			"mbean": "object_value1:flavor=chocolate",
			"type": "read"
		},
		"value": {
			"foo": 456
		},
		"status": 200
	}, {
		"request": {
			"mbean": "object_value2:flavor=chocolate",
			"type": "read"
		},
		"value": {
			"bar": 789
		},
		"status": 200
	}]`

	server := setupServer(http.StatusOK, response)
	defer server.Close()

	jolokia, err := setupPlugin(fmt.Sprintf(config, server.URL))
	if err != nil {
		t.Fatalf("Could not setup plugin. %v", err)
	}

	var acc testutil.Accumulator
	assert.NoError(t, jolokia.Gather(&acc))

	acc.AssertContainsTaggedFields(t, "compact_metric", map[string]interface{}{
		"value": 123.0,
		"foo":   456.0,
		"bar":   789.0,
	}, map[string]string{
		"flavor":            "chocolate",
		"jolokia_agent_url": server.URL,
	})

	acc.AssertContainsTaggedFields(t, "compact_metric", map[string]interface{}{
		"value": 999.0,
	}, map[string]string{
		"flavor":            "vanilla",
		"jolokia_agent_url": server.URL,
	})
}

func TestJolokia2_ProxyTargets(t *testing.T) {
	config := `
	[[inputs.jolokia2]]
		[inputs.jolokia2.proxy]
			url = "%s"

			[[inputs.jolokia2.proxy.target]]
				url = "service:jmx:rmi:///jndi/rmi://target1:9010/jmxrmi"
			[[inputs.jolokia2.proxy.target]]
				url = "service:jmx:rmi:///jndi/rmi://target2:9010/jmxrmi"

		[[inputs.jolokia2.metric]]
			name  = "hello"
			mbean = "hello:foo=bar"`

	response := `[{
		"request": {
			"type": "read",
			"mbean": "hello:foo=bar",
			"target": {
				"url": "service:jmx:rmi:///jndi/rmi://target1:9010/jmxrmi"
			}
		},
		"value": 123,
		"status": 200
	}, {
		"request": {
			"type": "read",
			"mbean": "hello:foo=bar",
			"target": {
				"url": "service:jmx:rmi:///jndi/rmi://target2:9010/jmxrmi"
			}
		},
		"value": 456,
		"status": 200
	}]`

	server := setupServer(http.StatusOK, response)
	defer server.Close()

	jolokia, err := setupPlugin(fmt.Sprintf(config, server.URL))
	if err != nil {
		t.Fatalf("Could not setup plugin. %v", err)
	}

	var acc testutil.Accumulator
	assert.NoError(t, jolokia.Gather(&acc))

	acc.AssertContainsTaggedFields(t, "hello", map[string]interface{}{
		"value": 123.0,
	}, map[string]string{
		"jolokia_proxy_url":  server.URL,
		"jolokia_target_url": "service:jmx:rmi:///jndi/rmi://target1:9010/jmxrmi",
	})
	acc.AssertContainsTaggedFields(t, "hello", map[string]interface{}{
		"value": 456.0,
	}, map[string]string{
		"jolokia_proxy_url":  server.URL,
		"jolokia_target_url": "service:jmx:rmi:///jndi/rmi://target2:9010/jmxrmi",
	})
}

func setupServer(status int, resp string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		//body, err := ioutil.ReadAll(r.Body)
		//if err == nil {
		//	fmt.Println(string(body))
		//}

		fmt.Fprintln(w, resp)
	}))
}

func setupPlugin(conf string) (*Jolokia, error) {
	c := config.NewConfig()
	r := strings.NewReader(conf)
	err := c.ParseConfig(r)
	if err != nil {
		return nil, err
	}

	jc := c.Inputs[0].Input.(*Jolokia)
	if jc == nil {
		return nil, errors.New("Missing jolokia2 from config")
	}

	//jc.Agents.URLs = []string{url}
	return jc, nil
}

func parseConfig(t *testing.T, conf string) *Jolokia {
	c := config.NewConfig()
	r := strings.NewReader(conf)

	err := c.ParseConfig(r)
	if err != nil {
		t.Fatalf("Could not parse config! %v", err)
	}

	return c.Inputs[0].Input.(*Jolokia)
}
