package jolokia2

import (
	"sort"
	"strings"
)

func RequestPayload(metrics []Metric) []ReadRequest {
	var requests []ReadRequest
	for _, metric := range metrics {

		if len(metric.Paths) == 0 {
			requests = append(requests, ReadRequest{
				Mbean:      metric.Mbean,
				Attributes: []string{},
			})
		} else {
			attributes := make(map[string][]string)

			for _, path := range metric.Paths {
				segments := strings.Split(path, "/")
				attribute := segments[0]

				if _, ok := attributes[attribute]; !ok {
					attributes[attribute] = make([]string, 0)
				}

				if len(segments) > 1 {
					paths := attributes[attribute]
					attributes[attribute] = append(paths, strings.Join(segments[1:], "/"))
				}
			}

			rootAttributes := payloadAttributesWithoutPaths(attributes)
			if len(rootAttributes) > 0 {
				requests = append(requests, ReadRequest{
					Mbean:      metric.Mbean,
					Attributes: rootAttributes,
				})
			}

			for _, deepAttribute := range payloadAttributesWithPaths(attributes) {
				for _, path := range attributes[deepAttribute] {
					requests = append(requests, ReadRequest{
						Mbean:      metric.Mbean,
						Attributes: []string{deepAttribute},
						Path:       path,
					})
				}
			}
		}
	}

	return requests
}

func payloadAttributesWithoutPaths(attributes map[string][]string) []string {
	results := make([]string, 0)
	for attr, paths := range attributes {
		if len(paths) == 0 {
			results = append(results, attr)
		}
	}

	sort.Strings(results)
	return results
}

func payloadAttributesWithPaths(attributes map[string][]string) []string {
	results := make([]string, 0)
	for attr, paths := range attributes {
		if len(paths) != 0 {
			results = append(results, attr)
		}
	}

	sort.Strings(results)
	return results
}
