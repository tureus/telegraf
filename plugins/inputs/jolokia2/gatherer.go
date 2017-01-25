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

		if metric.Mbean != request.Mbean {
			continue
		}

		if !hasPattern {
			fieldMap := make(map[string]interface{})
			extractFieldsFromValue(request.Attributes, request.Path,
				metric.FieldPrefix, metric.FieldDelimiter, response.Value, fieldMap)

			tagMap := extractTagsFromName(request.Mbean,
				metric.AllowTags, metric.DenyTags, metric.TagPrefix, metric.TagDelimiter, tags)

			g.accumulator.AddFields(metric.Name, fieldMap, tagMap)

		} else {
			valueMap, ok := response.Value.(map[string]interface{})

			if !ok {
				panic(fmt.Sprintf("FIXME! There should be a map here for %s!\n", request.Mbean))
			}

			for mbeanName, mbeanValue := range valueMap {
				fieldMap := make(map[string]interface{})
				extractFieldsFromValue(request.Attributes, request.Path,
					metric.FieldPrefix, metric.FieldDelimiter, mbeanValue, fieldMap)

				tagMap := extractTagsFromName(mbeanName,
					metric.AllowTags, metric.DenyTags, metric.TagPrefix, metric.TagDelimiter, tags)

				g.accumulator.AddFields(metric.Name, fieldMap, tagMap)
			}
		}
	}
}

func extractTagsFromName(name string, allowTags, denyTags []string, tagPrefix, tagDelimiter string, appendTags map[string]string) map[string]string {
	tagMap := make(map[string]string)

	object := strings.SplitN(name, ":", 2)
	domain := object[0]
	if domain != "" && len(object) == 2 {
		properties := object[1]

		for _, property := range strings.Split(properties, ",") {
			propertyPair := strings.SplitN(property, "=", 2)
			if len(propertyPair) != 2 {
				continue
			}

			propertyName := propertyPair[0]
			if propertyName == "" {
				continue
			}

			if tagCanBeExtracted(propertyName, allowTags, denyTags) {
				if tagPrefix != "" {
					propertyName = tagPrefix + tagDelimiter + propertyName
				}

				tagMap[propertyName] = propertyPair[1]
			}
		}
	}

	for tagKey, tagValue := range appendTags {
		tagMap[tagKey] = tagValue
	}

	return tagMap
}

func tagCanBeExtracted(name string, allowTags, denyTags []string) bool {
	for _, t := range allowTags {
		if name == t {
			return true
		}
	}

	for _, t := range denyTags {
		if name == t {
			return false
		}
	}

	if len(allowTags) == 0 {
		return true
	}

	return false
}

func extractFieldsFromValue(attributes []string, path, fieldPrefix, fieldDelimiter string, value interface{}, fieldMap map[string]interface{}) {
	valueMap, ok := value.(map[string]interface{})
	if ok {
		// complex value
		if len(attributes) == 0 {
			// if there were no attributes requested,
			// then the keys are attributes
			fieldName := fieldPrefix
			extractInnerFieldsFromValue(fieldName, fieldDelimiter, valueMap, fieldMap)

		} else if len(attributes) == 1 {
			// if there was a single attribute requested,
			// then the keys are the attribute's properties
			fieldName := joinFieldName(attributes[0], path, fieldPrefix, fieldDelimiter)
			extractInnerFieldsFromValue(fieldName, fieldDelimiter, valueMap, fieldMap)

		} else {
			// if there were multiple attributes requested,
			// then the keys are the attribute names
			for _, attribute := range attributes {
				fieldName := joinFieldName(attribute, path, fieldPrefix, fieldDelimiter)
				extractInnerFieldsFromValue(fieldName, fieldDelimiter, valueMap[attribute], fieldMap)
			}
		}
	} else {
		// scalar value
		var fieldName string
		if len(attributes) == 0 {
			fieldName = joinFieldName("value", path, fieldPrefix, fieldDelimiter)
		} else {
			fieldName = joinFieldName(attributes[0], path, fieldPrefix, fieldDelimiter)
		}

		fieldMap[fieldName] = value
	}
}

func joinFieldName(attribute, path, prefix, delimiter string) string {
	fieldName := attribute
	if prefix != "" {
		fieldName = prefix + delimiter + fieldName
	}

	if path != "" {
		fieldName = fieldName + delimiter + strings.Replace(path, "/", delimiter, -1)
	}

	return fieldName
}

func extractInnerFieldsFromValue(name, delimiter string, value interface{}, fieldMap map[string]interface{}) {
	if valueMap, ok := value.(map[string]interface{}); ok {
		// keep going until we get to something that is not a map
		for key, innerValue := range valueMap {
			var innerName string

			if name == "" {
				innerName = key
			} else {
				innerName = name + delimiter + key
			}

			extractInnerFieldsFromValue(innerName, delimiter, innerValue, fieldMap)
		}

		return
	}

	if name == "" {
		name = "value"
	}

	fieldMap[name] = value
}
