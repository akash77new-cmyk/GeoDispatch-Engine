.PHONY: run-backend build-backend test-backend run-frontend build-frontend install-frontend

# Backend
run-backend:
	cd backend && go run ./cmd/server

build-backend:
	cd backend && go build -o bin/server ./cmd/server

test-backend:
	cd backend && go test ./...

# Frontend
install-frontend:
	cd frontend && npm install

run-frontend:
	cd frontend && npm run dev

build-frontend:
	cd frontend && npm run build
