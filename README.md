# Distributed Rate Limiter

Redis-backed rate limiter in Go with sliding window algorithm.

## Run

```bash
cd deployments
docker compose up --build
```

## Test

```bash
# Hit the limit (10 req/min)
for i in {1..12}; do curl localhost:8080/test; echo ""; done
```

## Config

| Variable | Default | Description |
|----------|---------|-------------|
| `REDIS_ADDR` | localhost:6379 | Redis address |
| `RATE_LIMIT` | 10 | Requests per window |
| `RATE_WINDOW_SECONDS` | 60 | Window in seconds |

## Endpoints

- `GET /health` - Health check
- `GET /test` - Rate limited endpoint


