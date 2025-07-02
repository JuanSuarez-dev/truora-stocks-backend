package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func setupRouter(t *testing.T) *gin.Engine {
	dsn := os.Getenv("TEST_COCKROACH_DSN")
	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		t.Fatal(err)
	}
	r := gin.New()
	RegisterRoutes(r, pool)
	return r
}

func TestListStocks(t *testing.T) {
	r := setupRouter(t)
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/stocks", nil)
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("status %d; want 200", w.Code)
	}
}

func TestBestPick(t *testing.T) {
	r := setupRouter(t)
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/api/stocks/best", nil)
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("status %d; want 200", w.Code)
	}
}
