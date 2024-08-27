package indexer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/RyeHarvestProtocol/programmable-layer/pkg/utils/utils"
)

type GraphQLClient struct {
	URL     string
	Retries int
}

type GraphQLQuery struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables"`
}

func NewGraphQLClient(url string, retries int) *GraphQLClient {
	return &GraphQLClient{
		URL:     url,
		Retries: retries,
	}
}

func (client *GraphQLClient) GetBTCPrice(poolID string) (float64, error) {
	query := `
  query ($poolID: String!) {
    pools(where: {
      id: $poolID
    }) {
      id,
      token1Price,
    }
  }`

	variables := map[string]interface{}{
		"poolID": poolID,
	}

	res, err := client.ExecuteQuery(query, variables)
	if err != nil {
		return 0.0, fmt.Errorf("failed to execute query: %v", err)
	}

	token1Price, err := extractToken1Price(res)
	if err != nil {
		return 0.0, fmt.Errorf("failed to extract token1 price: %v", err)
	}

	// convert string to float64
	token1PriceFloat64, err := utils.StrToFloat64(token1Price)
	if err != nil {
		return 0.0, fmt.Errorf("failed to convert token1 price to float64: %v", err)
	}

	return token1PriceFloat64, nil
}

func extractToken1Price(poolInfo map[string]interface{}) (string, error) {
	data, ok := poolInfo["data"].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("data not in expected format")
	}

	pools, ok := data["pools"].([]interface{})
	if !ok || len(pools) == 0 {
		return "", fmt.Errorf("pools not found or empty")
	}

	pool, ok := pools[0].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("pool data not in expected format")
	}

	token1Price, ok := pool["token1Price"].(string)
	if !ok {
		return "", fmt.Errorf("token1Price not found or not a string")
	}

	return token1Price, nil
}

func (client *GraphQLClient) ExecuteQuery(query string, variables map[string]interface{}) (map[string]interface{}, error) {
	payload := GraphQLQuery{
		Query:     query,
		Variables: variables,
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal query: %v", err)
	}

	for i := 0; i < client.Retries; i++ {
		req, err := http.NewRequest("POST", client.URL, bytes.NewBuffer(payloadBytes))
		if err != nil {
			return nil, fmt.Errorf("failed to create request: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			if i == client.Retries-1 {
				return nil, fmt.Errorf("request failed: %v", err)
			}
			continue
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read response body: %v", err)
		}

		if resp.StatusCode != http.StatusOK {
			if i == client.Retries-1 {
				return nil, fmt.Errorf("non-200 response: %d %s", resp.StatusCode, string(body))
			}
			continue
		}

		var response map[string]interface{}
		err = json.Unmarshal(body, &response)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal response: %v", err)
		}

		return response, nil
	}

	return nil, fmt.Errorf("all retries failed")
}
