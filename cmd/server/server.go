package main

import (
	"context"
	"fmt"
	"math"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/JuanSuarez-dev/truora-stocks-backend/models"
)


// RegisterRoutes monta todos los endpoints sobre el router r
func RegisterRoutes(r *gin.Engine, pool *pgxpool.Pool) {
	// 1) List all stocks
	r.GET("/api/stocks", func(c *gin.Context) {
		rows, err := pool.Query(context.Background(), `
      SELECT ticker, company, brokerage, action,
             rating_from, rating_to,
             target_from, target_to, time
      FROM stocks
      ORDER BY ticker
    `)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		var list []models.StockItem
		for rows.Next() {
			var it models.StockItem
			rows.Scan(&it.Ticker, &it.Company, &it.Brokerage, &it.Action,
				&it.RatingFrom, &it.RatingTo, &it.TargetFrom, &it.TargetTo, &it.Time)
			list = append(list, it)
		}
		c.JSON(http.StatusOK, list)
	})

	// 2) Best pick
	r.GET("/api/stocks/best", func(c *gin.Context) {
		rows, err := pool.Query(context.Background(), `
      SELECT ticker, target_from, target_to FROM stocks
    `)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		bestTicker := ""
		bestUpside := -math.MaxFloat64
		for rows.Next() {
			var ticker, fromS, toS string
			rows.Scan(&ticker, &fromS, &toS)

			from := parseDollar(fromS)
			to := parseDollar(toS)
			if from == 0 || math.IsNaN(from) {
				continue
			}
			upside := (to - from) / from * 100
			if upside > bestUpside {
				bestUpside = upside
				bestTicker = ticker
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"ticker": bestTicker,
			"upside": fmt.Sprintf("%.1f%%", bestUpside),
		})
	})
}

func main() {
	// Leer .env o env
	dsn := os.Getenv("COCKROACH_DSN")
	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		panic(err)
	}
	defer pool.Close()

	r := gin.Default()
	RegisterRoutes(r, pool)      // monta todos los endpoints
	r.Run(":8080")               // arranca
}
