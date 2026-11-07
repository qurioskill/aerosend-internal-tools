# ops-dashboard

## Overview

A private Vue 3 and Vite single-page app (`aerosend-ops-dashboard`) that shows RabbitMQ queue health and IP pool health. `index.html` loads `/src/main.js`. Dependencies are `vue` and `vue-router`, and the `package.json` scripts are `dev` and `build`.

## Subdirectories & Boundaries

- src: [./src/AGENTS.md](./src/AGENTS.md) - App bootstrap, router, and the API client for the RabbitMQ management API and core's REST API, with the routed pages under `src/views`.

## Verification

- Run: `npm run build`, after `npm install`
- `package.json` has no test script, so the build is the only check.
- The build currently fails with `Install @vitejs/plugin-vue to handle .vue files`. `@vitejs/plugin-vue` is in `devDependencies`, but there is no `vite.config.*` to register it.
