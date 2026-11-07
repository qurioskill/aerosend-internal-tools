# src

## Overview

The Vue 3 application source: the bootstrap and router, the root layout, and the API client that the views share.

## Interface

- `main.js`: creates the app and a `createWebHistory` router, then mounts on `#app`. Routes: `/` redirects to `/queues`, `/queues` renders `QueueHealth`, `/ip-pools` renders `IpPoolHealth`.
- `App.vue`: the nav links and `<RouterView />`.
- `api.js`:
  - `queues()`: calls `GET ${RABBIT}/queues` and maps each entry to `{ name, depth, consumers }` from `name`, `messages` and `consumers`.
  - `ipPools()`: calls `GET ${CORE}/ip-pools` and returns the parsed JSON unchanged.
- Page components: [./views/AGENTS.md](./views/AGENTS.md).

## Invariants

- `RABBIT` comes from `VITE_RABBIT_API`, defaulting to `http://localhost:15672/api`. `CORE` comes from `VITE_CORE_API`, defaulting to `http://localhost:8081/api/v1`.
- All backend access goes through `api.js`.
- `queues()` always sends hardcoded `guest:guest` Basic auth, whichever host `VITE_RABBIT_API` names.
- Neither function checks `res.ok`, so an HTTP error surfaces only if the body fails to parse or has the wrong shape.
- `App.vue` labels the first link "Queues & CDC", but nothing here fetches Debezium or CDC data.
- As of this writing, `repos/core` has no `/api/v1/ip-pools` route (only `DomainController`), so `ipPools()` has no backend in that checkout.

## Side Effects

- Sends HTTP `GET` requests from the browser to the RabbitMQ management API and to core. Nothing here writes to either.

## Verification

- Run: `npm run build` from `repos/internal-tools/ops-dashboard`
- The project has no test script. See [the ops-dashboard contract](../AGENTS.md) for why the build currently fails.
