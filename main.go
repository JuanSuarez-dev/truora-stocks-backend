package main

import (
	"context"
	"fmt"
	"log"

	"github.com/JuanSuarez-dev/truora-stocks-backend/config"
	"github.com/JuanSuarez-dev/truora-stocks-backend/db"
	"github.com/JuanSuarez-dev/truora-stocks-backend/fetch"
)

func main() {
	// Carga .env
	cfg := config.Load()

	// Conecta a CockroachDB
	pool := db.Connect(cfg.CockroachDSN)
	defer pool.Close()

	// Paginación + upsert
	next := "https://8j5baasof2.execute-api.us-west-2.amazonaws.com/production/swechallenge/list"
	for next != "" {
		resp, err := fetch.FetchPage(cfg.APIToken, next)
		if err != nil {
			log.Fatalf("❌ FetchPage error: %v", err)
		}
		fmt.Printf("→ %d items, next_page=%q\n", len(resp.Items), resp.NextPage)

		for _, it := range resp.Items {
			_, err := pool.Exec(context.Background(), `
        INSERT INTO stocks (
          ticker, company, brokerage, action,
          rating_from, rating_to, target_from, target_to, time
        ) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)
        ON CONFLICT (ticker) DO UPDATE SET
          company     = EXCLUDED.company,
          brokerage   = EXCLUDED.brokerage,
          action      = EXCLUDED.action,
          rating_from = EXCLUDED.rating_from,
          rating_to   = EXCLUDED.rating_to,
          target_from = EXCLUDED.target_from,
          target_to   = EXCLUDED.target_to,
          time        = EXCLUDED.time
      `, it.Ticker, it.Company, it.Brokerage, it.Action,
				it.RatingFrom, it.RatingTo, it.TargetFrom, it.TargetTo, it.Time)
			if err != nil {
				log.Printf("⚠️ upsert %s: %v", it.Ticker, err)
			}
		}

		next = resp.NextPage
	}

	fmt.Println("✅ Part 1 done: all data stored in CockroachDB")
}
