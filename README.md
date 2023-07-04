# Sample project for building a Vue3 + Go Fiber app

This project contains a `config.yaml` sample which is loaded via `server/config/config.go` and available via the `config.Cfg` type throughout the project

Supabase is initialised in `server/supa/supabase.go` and made available by passing the `*supabase.Client` to the handlers
Get started using Supabase by creating a project and supplying the url and key in the `.local/config.yaml`

## Notable changes to the default Vue setup

1. TailwindCSS has been added
2. Axios has been added
3. The `vite.config.ts` has been modified to proxy the requests to the go server (useful in development).
