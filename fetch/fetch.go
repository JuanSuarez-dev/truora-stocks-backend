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

func FetchPage(token, url string) (APIResponse, error) {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return APIResponse{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return APIResponse{}, err
	}

	// ————————— Debug aquí —————————
	fmt.Printf("DEBUG: GET %s\n", url)
	fmt.Printf("DEBUG: Status %d\n", resp.StatusCode)
	fmt.Printf("DEBUG: Body   %s\n\n", string(body))
	// ————————————————————————————————

	if resp.StatusCode != http.StatusOK {
		return APIResponse{}, fmt.Errorf("bad status: %d", resp.StatusCode)
	}

	var apiResp APIResponse
	if err := json.Unmarshal(body, &apiResp); err != nil {
		return APIResponse{}, err
	}
	return apiResp, nil
}
