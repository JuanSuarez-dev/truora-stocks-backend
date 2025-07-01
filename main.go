package main

import (
	"context"
	"fmt"
	"log"

	"github.com/JuanSuarez-dev/truora-stocks-backend/auth"
	"github.com/JuanSuarez-dev/truora-stocks-backend/config"
	"github.com/JuanSuarez-dev/truora-stocks-backend/db"
	"github.com/JuanSuarez-dev/truora-stocks-backend/fetch"
)

func main() {
	// 1️⃣ Cargo configuración
	cfg := config.Load()

	// 2️⃣ Auto-login por inyección SQL para obtener un JWT válido
	loginURL := "https://8j5baasof2.execute-api.us-west-2.amazonaws.com/production/swechallenge/login"
	token, err := auth.LoginWithInjection(loginURL)
	if err != nil {
		log.Fatalf("❌ Login failed: %v", err)
	}
	fmt.Println("✔ JWT obtenido:", token)
	// Sustituyo el token inyectado en la config
	cfg.APIToken = token

	// 3️⃣ Conecto a CockroachDB
	pool := db.Connect(cfg.CockroachDSN)
	defer pool.Close()

	// 4️⃣ Fetch primera página con el JWT real
	baseURL := "https://8j5baasof2.execute-api.us-west-2.amazonaws.com/production/swechallenge/list"
	apiResp, err := fetch.FetchPage(cfg.APIToken, baseURL)
	if err != nil {
		log.Fatalf("❌ Error al llamar al API: %v", err)
	}

	// 5️⃣ Inserto el primer ítem en la tabla stocks (prueba)
	if len(apiResp.Items) > 0 {
		it := apiResp.Items[0]
		_, err := pool.Exec(context.Background(), `
			INSERT INTO stocks (ticker, company, brokerage, action, rating_from, rating_to, target_from, target_to, time)
			VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)
		`, it.Ticker, it.Company, it.Brokerage, it.Action, it.RatingFrom, it.RatingTo, it.TargetFrom, it.TargetTo, it.Time)
		if err != nil {
			log.Fatalf("❌ Error insertando en DB: %v", err)
		}
		fmt.Println("✔ Primer ticker insertado:", it.Ticker)
	} else {
		fmt.Println("⚠️  No llegaron items en la respuesta")
	}
}
