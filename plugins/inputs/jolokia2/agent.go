package jolokia2

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

type ReadResponse struct {
	Status  int
	Request ReadRequest
	Value   interface{}
}

type ReadRequest struct {
	Mbean      string
	Attributes []string
	Path       string
}

type Agent struct {
	url      string
	client   *http.Client
	username string
	password string
}

type agentResponse struct {
	Status  int          `json:"status"`
	Request agentRequest `json:"request"`
	Value   interface{}  `json:"value"`
}

type agentRequest struct {
	Type      string      `json:"type"`
	Mbean     string      `json:"mbean"`
	Attribute interface{} `json:"attribute,omitempty"`
	Path      string      `json:"path,omitempty"`
}

func NewAgent(url string, config *remoteConfig) *Agent {
	client := &http.Client{
		Timeout: config.ResponseTimeout,
	}

	return &Agent{
		url:    url,
		client: client,
	}
}

func (a *Agent) Read(requests []ReadRequest) ([]ReadResponse, error) {
	requestObjects := makeRequests(requests)
	requestBody, err := json.Marshal(requestObjects)
	if err != nil {
		return nil, err
	}

	requestUrl, err := makeReadUrl(a.url, a.username, a.password)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", requestUrl, bytes.NewBuffer(requestBody))
	req.Header.Add("Content-type", "application/json")

	resp, err := a.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Response from url \"%s\" has status code %d (%s), expected %d (%s)",
			a.url, resp.StatusCode, http.StatusText(resp.StatusCode), http.StatusOK, http.StatusText(http.StatusOK))
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var responses []agentResponse
	if err = json.Unmarshal([]byte(responseBody), &responses); err != nil {
		return nil, fmt.Errorf("Error decoding JSON response: %s: %s", err, responseBody)
	}

	return makeResponses(responses), nil
}

func makeReadUrl(configUrl, username, password string) (string, error) {
	parsedUrl, err := url.Parse(configUrl)
	if err != nil {
		return "", err
	}

	readUrl := url.URL{
		Host:   parsedUrl.Host,
		Scheme: parsedUrl.Scheme,
	}

	if username != "" || password != "" {
		readUrl.User = url.UserPassword(username, password)
	}

	readUrl.Path = path.Join(parsedUrl.Path, "read")
	return readUrl.String(), nil
}

func makeRequests(requests []ReadRequest) []agentRequest {
	requestObjects := make([]agentRequest, len(requests))
	for i, request := range requests {
		requestObjects[i] = agentRequest{
			Type:  "read",
			Mbean: request.Mbean,
			Path:  request.Path,
		}
		if len(request.Attributes) == 1 {
			requestObjects[i].Attribute = request.Attributes[0]
		}
		if len(request.Attributes) > 1 {
			requestObjects[i].Attribute = request.Attributes
		}
	}

	return requestObjects
}

func makeResponses(responseObjects []agentResponse) []ReadResponse {
	responses := make([]ReadResponse, len(responseObjects))

	for i, object := range responseObjects {
		request := ReadRequest{
			Mbean:      object.Request.Mbean,
			Path:       object.Request.Path,
			Attributes: []string{},
		}

		attrValue := object.Request.Attribute
		if attrValue != nil {
			attribute, ok := attrValue.(string)
			if ok {
				request.Attributes = []string{attribute}
			} else {
				attributes, _ := attrValue.([]interface{})
				request.Attributes = make([]string, len(attributes))
				for i, attr := range attributes {
					request.Attributes[i] = attr.(string)
				}
			}
		}

		responses[i] = ReadResponse{
			Request: request,
			Value:   object.Value,
			Status:  object.Status,
		}
	}

	return responses
}
