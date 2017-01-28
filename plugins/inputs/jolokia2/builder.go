package jolokia2

import "strings"

type pointBuilder struct {
	metric  Metric
	request ReadRequest
}

type point struct {
	Tags   map[string]string
	Fields map[string]interface{}
}

func newPointBuilder(m Metric, r ReadRequest) *pointBuilder {
	return &pointBuilder{
		metric:  m,
		request: r,
	}
}

func (pb *pointBuilder) Build(mbean string, value interface{}) *point {
	tags := pb.extractTags(mbean)
	fields := pb.extractFields(value)

	return &point{
		Tags:   tags,
		Fields: fields,
	}
}

func (pb *pointBuilder) extractTags(mbean string) map[string]string {
	tagMap := make(map[string]string)

	objectPair := strings.SplitN(mbean, ":", 2)
	domainName := objectPair[0]

	if domainName != "" && len(objectPair) == 2 {
		propertyList := objectPair[1]

		for _, objectProperty := range strings.Split(propertyList, ",") {
			propertyPair := strings.SplitN(objectProperty, "=", 2)
			if len(propertyPair) != 2 {
				continue
			}

			propertyName := propertyPair[0]
			if propertyName != "" && pb.includeTag(propertyName) {
				tagName := pb.formatTagName(propertyName)
				tagMap[tagName] = propertyPair[1]
			}
		}
	}

	return tagMap
}

func (pb *pointBuilder) includeTag(tagName string) bool {
	for _, t := range pb.metric.TagKeys {
		if tagName == t {
			return true
		}
	}

	for _, t := range pb.metric.UntagKeys {
		if tagName == t {
			return false
		}
	}

	if len(pb.metric.TagKeys) == 0 {
		// Implicitly expose all mbean properties as tags.
		return true
	}

	return false
}

func (pb *pointBuilder) formatTagName(tagName string) string {
	if tagName == "" {
		return ""
	}

	if tagPrefix := pb.metric.TagPrefix; tagPrefix != "" {
		return tagPrefix + pb.metric.TagSeparator + tagName
	}

	return tagName
}

func (pb *pointBuilder) extractFields(value interface{}) map[string]interface{} {
	attributes := pb.request.Attributes
	path := pb.request.Path

	fieldMap := make(map[string]interface{})
	valueMap, ok := value.(map[string]interface{})

	if ok {
		// complex value
		if len(attributes) == 0 {
			// if there were no attributes requested,
			// then the keys are attributes
			fieldName := pb.metric.FieldPrefix
			pb.fillFields(fieldName, valueMap, fieldMap)

		} else if len(attributes) == 1 {
			// if there was a single attribute requested,
			// then the keys are the attribute's properties
			fieldName := pb.formatFieldName(attributes[0], path)
			pb.fillFields(fieldName, valueMap, fieldMap)

		} else {
			// if there were multiple attributes requested,
			// then the keys are the attribute names
			for _, attribute := range attributes {
				fieldName := pb.formatFieldName(attribute, path)
				pb.fillFields(fieldName, valueMap[attribute], fieldMap)
			}
		}
	} else {
		// scalar value
		var fieldName string
		if len(attributes) == 0 {
			fieldName = pb.formatFieldName("value", path)
		} else {
			fieldName = pb.formatFieldName(attributes[0], path)
		}

		fieldMap[fieldName] = value
	}

	return fieldMap
}

func (pb *pointBuilder) formatFieldName(attribute, path string) string {
	fieldName := attribute
	fieldPrefix := pb.metric.FieldPrefix
	fieldSeparator := pb.metric.FieldSeparator

	if fieldPrefix != "" {
		fieldName = fieldPrefix + fieldSeparator + fieldName
	}

	if path != "" {
		fieldName = fieldName + fieldSeparator + strings.Replace(path, "/", fieldSeparator, -1)
	}

	return fieldName
}

func (pb *pointBuilder) fillFields(name string, value interface{}, fieldMap map[string]interface{}) {
	if valueMap, ok := value.(map[string]interface{}); ok {
		separator := pb.metric.FieldSeparator

		// keep going until we get to something that is not a map
		for key, innerValue := range valueMap {
			var innerName string

			if name == "" {
				innerName = key
			} else {
				innerName = name + separator + key
			}

			pb.fillFields(innerName, innerValue, fieldMap)
		}

		return
	}

	if name == "" {
		name = "value"
	}

	fieldMap[name] = value
}
