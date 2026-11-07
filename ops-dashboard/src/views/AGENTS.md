# views

## Overview

The two routed pages of the ops dashboard. Each is a Vue 3 `<script setup>` component that loads data through `../api` and renders it.

## Interface

- `QueueHealth.vue`, routed at `/queues`: a table of queue name, depth and consumers from `queues()`.
- `IpPoolHealth.vue`, routed at `/ip-pools`: one block per pool from `ipPools()`.

## Invariants

- `QueueHealth` reloads every 5000 ms and must clear its timer in `onUnmounted`. `IpPoolHealth` loads once on mount and does not refresh.
- A queue row gets the `warn` class when `depth > 1000` or `consumers === 0`. No stylesheet or `<style>` block defines `warn`, so it has no visual effect today.
- There is no DLQ-specific logic. A dead-letter queue is treated like any other row.
- `IpPoolHealth` expects each pool to have `id`, `name`, `ips`, `avgHealth` and `unhealthy`.
- Neither view handles a failed load.

## Side Effects

- Sends HTTP requests through `../api` on mount, and every 5 seconds for `QueueHealth`.

## Verification

- Run: `npm run build` from `repos/internal-tools/ops-dashboard`
- The project has no test script. See [the ops-dashboard contract](../../AGENTS.md) for why the build currently fails.
