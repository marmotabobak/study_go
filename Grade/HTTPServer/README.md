Build with version: go build -ldflags "-X main.Version=0.0.1_$(date -u +%Y%m%d%H%M%S)" -o httpserver ./cmd/httpserver
