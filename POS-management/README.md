# POS Management Monorepo

This repository contains a Go backend and a React (Vite) frontend.

## Structure

- `backend/` Go HTTP server (entry: `cmd/server/main.go`)
- `frontend/` Vite React app
- `docker-compose.yml` Local dev: runs both services

## Quickstart (Windows PowerShell)

Prereqs: Go 1.22+, Node.js 18+ and npm, Docker Desktop 

### Run backend locally

1. Open a terminal and run:

```
cd "./backend"
go run ./cmd/server
```

Server listens on `http://localhost:8080`.

### Run frontend locally

In another terminal:

```
cd "./frontend"
npm install
npm run dev
```

Open `http://localhost:5173`. API calls to `/api/*` are proxied to the backend.

### Run with Docker Compose

```
docker compose up --build
```

- Frontend: http://localhost:5173
- Backend: http://localhost:8080

## API

- GET `/api/health` -> `{ "status": "ok" }`


