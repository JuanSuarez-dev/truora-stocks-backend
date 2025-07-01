package fetch

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/JuanSuarez-dev/truora-stocks-backend/models"
)

type APIResponse struct {
	Items    []models.StockItem `json:"items"`
	NextPage string             `json:"next_page"`
}

// sampleJSON es la respuesta de ejemplo
const sampleJSON = `{
  "items": [
    {
      "ticker": "BSBR",
      "target_from": "$4.20",
      "target_to": "$4.70",
      "company": "Banco Santander (Brasil)",
      "action": "upgraded by",
      "brokerage": "The Goldman Sachs Group",
      "rating_from": "Sell",
      "rating_to": "Neutral",
      "time": "2025-01-13T00:30:05.813548892Z"
    },
    {
      "ticker": "VYGR",
      "target_from": "$11.00",
      "target_to": "$9.00",
      "company": "Voyager Therapeutics",
      "action": "reiterated by",
      "brokerage": "Wedbush",
      "rating_from": "Outperform",
      "rating_to": "Outperform",
      "time": "2025-01-14T00:30:05.813548892Z"
    }
  ],
  "next_page": ""
}`

// FetchPage devuelve sampleJSON si token=="bypass-token", sino hace la llamada real.
func FetchPage(token, url string) (APIResponse, error) {
	// stub con bypass-token
	if token == "bypass-token" {
		var resp APIResponse
		if err := json.Unmarshal([]byte(sampleJSON), &resp); err != nil {
			return APIResponse{}, err
		}
		return resp, nil
	}
	// caso real (más adelante)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return APIResponse{}, err
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	if res.StatusCode != http.StatusOK {
		return APIResponse{}, fmt.Errorf("bad status: %d – body: %s", res.StatusCode, string(body))
	}

	var apiResp APIResponse
	if err := json.Unmarshal(body, &apiResp); err != nil {
		return APIResponse{}, err
	}
	return apiResp, nil
}
