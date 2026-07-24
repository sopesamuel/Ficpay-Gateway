# mart-gateway

## PostgreSQL setup

Use the provided Docker Compose service to start PostgreSQL locally.

1. Copy `.env.example` to `.env` or set `DATABASE_URL` in your environment.
2. Run:

```bash
docker-compose up -d
```

3. The database will be available at:

```text
postgres://pguser:pgpass@localhost:5432/payment_gateway_db?sslmode=disable
```

4. Start the Go app with `DATABASE_URL` set.

```bash
export DATABASE_URL="postgres://pguser:pgpass@localhost:5432/payment_gateway_db?sslmode=disable"
go run ./cmd/gateway
```
