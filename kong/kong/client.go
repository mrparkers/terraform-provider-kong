package kong

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	rbacHeader = "Kong-Admin-Token"
)

type KongConfig struct {
	AdminApiUrl string
	RbacToken   string
}

type KongClient struct {
	Config KongConfig

	client *http.Client
}

func NewKongClient(config KongConfig) (*KongClient, error) {
	httpClient := &http.Client{
		Timeout: time.Second * 60,
	}

	client := &KongClient{Config: config, client: httpClient}

	status, err := client.GetStatus()

	if err != nil {
		return nil, fmt.Errorf("error initializing Kong client: %s", err)
	}

	if !status.Database.Reachable {
		return nil, fmt.Errorf("/status check returned, but Kong indicated that the database isn't reachable. This client is unlikely to work")
	}

	return client, nil
}

func (kongClient *KongClient) post(path string, payload interface{}, responseResource interface{}) error {
	return kongClient.request("POST", path, payload, responseResource)
}

func (kongClient *KongClient) put(path string, payload interface{}) error {
	return kongClient.request("PUT", path, payload, nil)
}

func (kongClient *KongClient) get(path string, responseResource interface{}) error {
	return kongClient.request("GET", path, nil, responseResource)
}

func (kongClient *KongClient) delete(path string) error {
	return kongClient.request("DELETE", path, nil, nil)
}

func (kongClient *KongClient) patch(path string, payload interface{}) error {
	return kongClient.request("PATCH", path, payload, nil)
}

func (kongClient *KongClient) request(method string, path string, payload interface{}, responseResource interface{}) error {
	endpoint := kongClient.Config.AdminApiUrl + path

	serializedPayload, err := serialize(method, payload)

	if err != nil {
		return err
	}

	request, err := http.NewRequest(method, endpoint, serializedPayload)

	if err != nil {
		return err
	}

	kongClient.addDefaultHeaders(request)

	if methodShouldHavePayload(method) {
		request.Header.Set("Content-Type", "application/json")
	}

	response, err := kongClient.client.Do(request)

	if err != nil {
		return err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return err
	}

	if response.StatusCode >= 400 {
		errorMessage := string(body[:])

		return &HttpError{StatusCode: response.StatusCode, Message: errorMessage}
	}

	if responseResource != nil {
		return json.Unmarshal(body, responseResource)
	}

	return nil
}

func serialize(method string, payload interface{}) (io.Reader, error) {
	if methodShouldHavePayload(method) && payload != nil {
		serializedPayload, err := json.Marshal(payload)

		if err != nil {
			return nil, err
		}

		return bytes.NewReader(serializedPayload), nil
	}

	return nil, nil
}

func methodShouldHavePayload(method string) bool {
	return method == "PUT" || method == "POST" || method == "PATCH"
}

func (kongClient *KongClient) addDefaultHeaders(request *http.Request) {
	if kongClient.Config.RbacToken != "" {
		request.Header.Set(rbacHeader, kongClient.Config.RbacToken)
	}

	request.Header.Set("Accept", "application/json")
}
