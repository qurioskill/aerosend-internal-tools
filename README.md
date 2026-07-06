# aerosend-internal-tools

Tools the engineering/ops team uses day to day. Not customer-facing.

| | |
|---|---|
| Category | internal tools |
| Change frequency | **medium** |
| Stack | **Go** (`aeroctl` CLI), **Vue 3** + Vite (`ops-dashboard`) |
| Owners | Platform team |

## aeroctl (Go)

```bash
cd aeroctl && go build -o aeroctl .
./aeroctl dns check getacme.com          # verify SPF/DKIM/DMARC resolve as core generated them
./aeroctl ip rotate --pool pool-us-east-3 --below 70   # rotate unhealthy IPs via core API
```

## ops-dashboard (Vue)

Live view of RabbitMQ queue depths (incl. DLQ), Debezium lag, and IP pool health.

```bash
cd ops-dashboard && npm install && npm run dev
```
