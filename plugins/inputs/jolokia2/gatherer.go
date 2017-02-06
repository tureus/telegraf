package jolokia2

import (
	"strings"

	"github.com/influxdata/telegraf"
)

type Metric struct {
	Name           string
	Mbean          string
	Paths          []string
	TagKeys        []string
	UntagKeys      []string
	FieldName      string
	FieldPrefix    string
	FieldSeparator string
	TagPrefix      string
	TagSeparator   string
}

type Gatherer struct {
	metrics     []Metric
	accumulator telegraf.Accumulator
}

type point struct {
	Tags   map[string]string
	Fields map[string]interface{}
}

func NewGatherer(metrics []Metric, acc telegraf.Accumulator) *Gatherer {
	return &Gatherer{
		metrics:     metrics,
		accumulator: acc,
	}
}

// Gather adds points to the accumulator from the ReadResponse objects
// returned by a Jolokia agent.
func (g *Gatherer) Gather(responses []ReadResponse, tags map[string]string) {
	series := make(map[string][]point, 0)

	for _, metric := range g.metrics {
		points, ok := series[metric.Name]
		if !ok {
			points = make([]point, 0)
		}

		for _, point := range g.gatherPoints(metric, responses) {
			points = append(points, point)
		}

		series[metric.Name] = points
	}

	for measurement, points := range series {
		for _, point := range compactPoints(points) {
			g.accumulator.AddFields(measurement,
				point.Fields, gatherTags(point.Tags, tags))
		}
	}
}

// gatherMetric generates points for the supplied metric from the ReadResponse
// objects returned by a Jolokia agent.
func (g *Gatherer) gatherPoints(metric Metric, responses []ReadResponse) []point {
	points := make([]point, 0)

	for _, response := range responses {
		if response.Status != 200 {
			// TODO:
			//acc.AddError(fmt.Errorf("Not expected status value in response body (%s:%s mbean=\"%s\" attribute=\"%s\"): %3.f",
			//	server.Host, server.Port, metrics[i].Mbean, metrics[i].Attribute, status))
			continue
		}

		request := response.Request
		if !metricMatchesRequest(metric, request) {
			continue
		}

		pb := newPointBuilder(metric, response.Request)
		for _, point := range pb.Build(metric.Mbean, response.Value) {
			points = append(points, point)
		}
	}

	return points
}

// gatherTags merges two tag sets into a single tag set.
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

// metricMatchesRequest returns true when the name, attributes, and path
// of a Metric match the corresponding elements in a ReadRequest object
// returned by a Jolokia agent.
func metricMatchesRequest(metric Metric, request ReadRequest) bool {
	if metric.Mbean != request.Mbean {
		return false
	}

	if len(metric.Paths) == 0 {
		return len(request.Attributes) == 0
	}

	for _, fullPath := range metric.Paths {
		segments := strings.SplitN(fullPath, "/", 2)
		attribute := segments[0]

		var path string
		if len(segments) == 2 {
			path = segments[1]
		}

		for _, rattr := range request.Attributes {
			if attribute == rattr {
				return path == request.Path
			}
		}
	}

	return false
}

// compactPoints attepts to remove points by compacting points
// with matching tag sets. When a match is found, the fields from
// one point are moved to another, and the empty point is removed.
func compactPoints(points []point) []point {
	compactedPoints := make([]point, 0)

	for _, sourcePoint := range points {
		keepPoint := true

		for _, compactPoint := range compactedPoints {
			if !tagSetsMatch(sourcePoint.Tags, compactPoint.Tags) {
				continue
			}

			keepPoint = false
			for key, val := range sourcePoint.Fields {
				compactPoint.Fields[key] = val
			}
		}

		if keepPoint {
			compactedPoints = append(compactedPoints, sourcePoint)
		}
	}

	return compactedPoints
}

// tagSetsMatch returns true if two maps are equivalent.
func tagSetsMatch(a, b map[string]string) bool {
	if len(a) != len(b) {
		return false
	}

	for ak, av := range a {
		bv, ok := b[ak]
		if !ok {
			return false
		}
		if av != bv {
			return false
		}
	}

	return true
}
