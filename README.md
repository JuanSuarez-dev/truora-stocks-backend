# Truora Stocks Challenge

**Project Overview**
A full-stack application that retrieves stock recommendation data from an external API, stores it in CockroachDB, exposes a Go/Gin REST API, and provides a Vue 3 + TypeScript + Pinia frontend to display and recommend the best stocks.

---

## Tech Stack

* **Backend**: Go, Gin, pgxpool (PostgreSQL driver)
* **Database**: CockroachDB
* **Frontend**: Vue 3, TypeScript, Pinia, Axios, Tailwind CSS
* **Testing**: Go test, Vitest + Vue Test Utils
* **Dev Tools**: Docker (optional), Vite

---

## Prerequisites

* Go ≥ 1.20
* Node.js ≥ 16
* Docker (optional, for CockroachDB local)
* Git

---

## Setup & Run Locally

1. **Clone the repo**

   ```bash
   git clone https://github.com/JuanSuarez-dev/truora-stocks-backend.git
   cd truora-stocks-backend
   ```

2. **Configure CockroachDB**

   * Export DSN (PowerShell example):

     ```powershell
     $Env:COCKROACH_DSN='postgresql://<user>:<pass>@<host>:26257/defaultdb?sslmode=verify-full&sslrootcert=C:\Users\<you>\AppData\Roaming\postgresql\root.crt'
     ```

3. **Populate Database** (Part 1)

   ```bash
   go run ./cmd/ingest
   ```

4. **Start API Server** (Part 2)

   ```bash
   go run ./cmd/server
   ```

5. **Serve Frontend** (Part 2)

   ```bash
   cd frontend
   npm install
   npm run dev
   ```

6. **Open** `http://localhost:5173`

---

## API Endpoints

* **GET** `/api/stocks` — list all stored stocks
* **GET** `/api/stocks/best` — returns `{ ticker, upside }` for the highest-upside stock

---

## Running Tests

* **Backend**:

  ```bash
  go test ./...
  ```

* **Frontend** (from `frontend/`):

  ```bash
  npm run test
  ```

---
