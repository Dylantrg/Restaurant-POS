# POS Management Monorepo

This repository contains a Go backend and a React (Vite) frontend.

## Structure

- `backend/` Go HTTP server (entry: `cmd/server/main.go`)
- `frontend/` Vite React app
- `docker-compose.yml` Local dev: runs both services

## Quickstart (Windows PowerShell)

Prereqs: Go 1.22+, Node.js 18+ and npm, Docker Desktop (optional for compose)

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
# If PowerShell blocks npm.ps1, run the next line once as Administrator:
# Set-ExecutionPolicy -Scope CurrentUser -ExecutionPolicy RemoteSigned
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

## Next steps

- Add real models, repositories, and services
- Add auth middleware, structured logging, and config via env vars
- DB included: Postgres (Docker). Connection via `DATABASE_URL` env var.

## Database

Docker Compose now starts a Postgres 16 instance:

- Host: `localhost:5432`
- User: `pos`
- Password: `pos`
- DB: `posdb`

Backend env var example (already set in docker-compose):

```
DATABASE_URL=postgres://pos:pos@localhost:5432/posdb?sslmode=disable
```

Migrations:
- Place SQL files under `backend/migrations` (created `0001_init` already).
- You can use a tool like `golang-migrate` or run files manually with `psql`.

## WSL (Ubuntu 22.04) instructions

If you prefer running everything inside WSL:

1) Install prerequisites in WSL

```
sudo apt update
sudo apt install -y curl git build-essential

# Install Go 1.22+ (example using snap or tarball - choose your preferred method)
# sudo snap install go --classic

# Install Node.js 20.x
curl -fsSL https://deb.nodesource.com/setup_20.x | sudo -E bash -
sudo apt install -y nodejs

# Install Docker inside Windows and enable WSL integration (recommended),
# or install Docker in WSL if you know what you’re doing.
```

2) Run backend in WSL

```
cd /mnt/c/Users/ducdu/OneDrive/Desktop/Restaurant-POS/POS-management/backend
go run ./cmd/server
```

3) Run frontend in WSL

```
cd /mnt/c/Users/ducdu/OneDrive/Desktop/Restaurant-POS/POS-management/frontend
npm install
npm run dev
```

Notes:
- `vite --host` is already configured so the dev server binds to 0.0.0.0.
- Access from Windows: use the WSL IP or http://localhost:5173 if using the VS Code server/port forwarding. If you see access issues, try `npm run dev -- --host 0.0.0.0`.
- The Vite proxy forwards `/api` to the backend on `http://localhost:8080` by default. If backend runs in a different host/port in WSL, set `VITE_BACKEND_URL` accordingly, e.g. `VITE_BACKEND_URL=http://127.0.0.1:8080 npm run dev`.

4) Docker with WSL

If Docker Desktop is installed with WSL integration, from WSL you can run:

```
cd /mnt/c/Users/ducdu/OneDrive/Desktop/Restaurant-POS/POS-management
docker compose up --build
```

Then open http://localhost:5173 and http://localhost:8080 from Windows or WSL.
