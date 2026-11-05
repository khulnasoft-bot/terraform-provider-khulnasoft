package client

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/pkg/errors"
)

type SecureAIProvider struct {
	ApiKey  string `json:"api_key,omitempty"`
	ID      string `json:"id,omitempty"`
	Enabled bool   `json:"enabled"`
	Mode    string `json:"mode"`
}

type SecureAIIntegration struct {
	SelectedUI string            `json:"selected_ui"`
	Akamai     *SecureAIProvider `json:"akamai,omitempty"`
	Khulnasoft       *SecureAIProvider `json:"khulnasoft,omitempty"`
}

func (cli *Client) GetSecureAIIntegration() (*SecureAIIntegration, error) {
	var err error
	var response SecureAIIntegration
	request := cli.gorequest
	apiPath := "/api/v1/settings/integrations/secure_ai"
	err = cli.limiter.Wait(context.Background())
	if err != nil {
		return nil, err
	}
	resp, body, errs := request.Clone().Set("Authorization", "Bearer "+cli.token).Get(cli.url + apiPath).End()
	if errs != nil {
		return nil, fmt.Errorf("error calling %s", apiPath)
	}
	if resp == nil {
		return nil, fmt.Errorf("no response from %s", apiPath)
	}
	if resp.StatusCode == 200 {
		err = json.Unmarshal([]byte(body), &response)
		if err != nil {
			log.Printf("Error calling func GetSecureAIIntegration from %s%s, %v", cli.url, apiPath, err)
			return nil, err
		}
		return &response, nil
	}
	if resp.StatusCode == 404 {
		return nil, nil
	}
	return nil, fmt.Errorf("GetSecureAIIntegration: unexpected status %d from %s: %s", resp.StatusCode, apiPath, body)
}

func (cli *Client) SaveSecureAIIntegration(integration SecureAIIntegration) error {
	payload, err := json.Marshal(integration)
	if err != nil {
		return err
	}
	request := cli.gorequest
	apiPath := "/api/v1/settings/integrations/secure_ai"
	err = cli.limiter.Wait(context.Background())
	if err != nil {
		return err
	}
	resp, body, errs := request.Clone().Set("Authorization", "Bearer "+cli.token).Put(cli.url + apiPath).Send(string(payload)).End()
	if errs != nil {
		return errors.Errorf("error calling %s", apiPath)
	}
	if resp == nil {
		return fmt.Errorf("SaveSecureAIIntegration: no HTTP response (nil)")
	}
	if resp.StatusCode != 200 && resp.StatusCode != 201 && resp.StatusCode != 204 {
		return errors.Errorf("SaveSecureAIIntegration: unexpected status %d: %s", resp.StatusCode, body)
	}
	return nil
}
