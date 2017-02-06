package jolokia2

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/influxdata/telegraf/internal/config"
	"github.com/influxdata/telegraf/testutil"
	"github.com/stretchr/testify/assert"

	"github.com/influxdata/toml"
)

func TestJolokia2_ScalarValuesFixture(t *testing.T) {
	runFixture(t, "./testdata/scalar_values.toml")
}

func TestJolokia2_ObjectValuesFixture(t *testing.T) {
	runFixture(t, "./testdata/object_values.toml")
}

func TestJolokia2_TagsAndFieldsFixture(t *testing.T) {
	runFixture(t, "./testdata/tags_and_fields.toml")
}

func TestJolokia2_SubstitutionFixture(t *testing.T) {
	runFixture(t, "./testdata/substitution.toml")
}

func TestJolokia2_CompactionFixture(t *testing.T) {
	runFixture(t, "./testdata/compaction.toml")
}

func TestJolokia2_JvmFixture(t *testing.T) {
	runFixture(t, "./testdata/jvm.toml")
}

func TestJolokia2_KafkaLogFixture(t *testing.T) {
	runFixture(t, "./testdata/kafka_log.toml")
}

func TestJolokia2_KafkaTopicFixture(t *testing.T) {
	runFixture(t, "./testdata/kafka_topic.toml")
}

func runFixture(t *testing.T, path string) {
	fixture := setupFixture(t, path)

	server := setupServer(http.StatusOK, fixture.Response)
	defer server.Close()

	jolokia, err := setupPlugin(fixture.Config, server.URL)
	if err != nil {
		t.Fatalf("Could not setup plugin. %v", err)
	}

	var acc testutil.Accumulator
	err = jolokia.Gather(&acc)

	assert.Nil(t, err)
	for _, expect := range fixture.Expects {
		expect.Tags["jolokia_agent_url"] = server.URL
		acc.AssertContainsTaggedFields(t,
			expect.Measurement, expect.Fields, expect.Tags)
	}
}

type Fixture struct {
	Config   string
	Expects  []Expect
	Response string
}

type Expect struct {
	Measurement string
	Tags        map[string]string
	Fields      map[string]interface{}
}

func setupFixture(t *testing.T, path string) *Fixture {
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatalf("Could not read fixture %s", path)
	}

	var fixture Fixture
	if err := toml.Unmarshal(contents, &fixture); err != nil {
		t.Fatalf("Could not unmarshal fixture. %v", err)
	}

	return &fixture
}

func setupServer(status int, resp string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, resp)
	}))
}

func setupPlugin(conf, url string) (*Jolokia, error) {
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

	jc.Agents.Urls = []string{url}
	return jc, nil
}
