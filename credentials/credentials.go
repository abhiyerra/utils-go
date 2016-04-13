package fugue_credentials

import (
	"fmt"

	"github.com/acksin/bridge/api"
)

const (
	CredentialsService = "fugue.auth.credentials"

	BridgeAPIURL = "https://bridge-api.acksin.com/lambda"
)

func GetSessionID(apiKey string) (string, error) {
	resp, err := bridge.Request{
		Service: CredentialsService,
		Action:  "Get",
		Method:  "POST",
		Version: "1.0",
		Async:   false,
		Payload: struct {
			APIKey string
		}{apiKey},
	}.POST(BridgeAPIURL)

	if err != nil {
		return "", err
	}

	if resp.Error != "" {
		return "", fmt.Errorf("%s", resp.Error)
	}

	return resp.Payload.(map[string]interface{})["SessionID"].(string), nil
}
