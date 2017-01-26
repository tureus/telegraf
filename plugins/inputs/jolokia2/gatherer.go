package jolokia2

import (
	"fmt"
	"strings"

	"github.com/influxdata/telegraf"
)

type Metric struct {
	Name           string
	Mbean          string
	Paths          []string
	AllowTags      []string
	DenyTags       []string
	FieldPrefix    string
	FieldDelimiter string
	TagPrefix      string
	TagDelimiter   string
}

type Gatherer struct {
	metrics     []Metric
	accumulator telegraf.Accumulator
}

func NewGatherer(metrics []Metric, acc telegraf.Accumulator) *Gatherer {
	return &Gatherer{
		metrics:     metrics,
		accumulator: acc,
	}
}

func (g *Gatherer) Gather(responses []ReadResponse, tags map[string]string) {
	for _, metric := range g.metrics {
		g.gatherMetric(metric, responses, tags)
	}
}

func (g *Gatherer) gatherMetric(metric Metric, responses []ReadResponse, tags map[string]string) {
	hasPattern := strings.Contains(metric.Mbean, "*")

	for _, response := range responses {
		request := response.Request

		// FIXME: it takes more to correlate metrics and
		// responses than an object name/pattern.
		if metric.Mbean != request.Mbean {
			continue
		}

		pb := newPointBuilder(metric, response.Request)
		if !hasPattern {
			point := pb.Build(metric.Mbean, response.Value)

			g.accumulator.AddFields(metric.Name,
				point.Fields, gatherTags(point.Tags, tags))

		} else {
			valueMap, ok := response.Value.(map[string]interface{})

			if !ok {
				// FIXME: log it and move on.
				panic(fmt.Sprintf("There should be a map here for %s!\n", request.Mbean))
			}

			for mbean, value := range valueMap {
				point := pb.Build(mbean, value)

				g.accumulator.AddFields(metric.Name,
					point.Fields, gatherTags(point.Tags, tags))
			}
		}
	}
}

func gatherTags(metricTags, outerTags map[string]string) map[string]string {
	tags := make(map[string]string)
	for k, v := range outerTags {
		tags[k] = v
	}
	for k, v := range metricTags {
		tags[k] = v
	}

	return tags
}
