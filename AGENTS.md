# AGENTS.md

This file provides guidance to coding agents (e.g. Claude Code, claude.ai/code) when working with code in this repository.

## Repository purpose

Go module `go.bytebuilders.dev/catalog` — API types and CRDs for **catalog bindings**: the AppsCode platform's representation of customer-managed DB instances and gateway resources (`DruidBinding`, `ElasticsearchBinding`, `CassandraBinding`, `ClickhouseBinding`, `FerretDBBinding`, …, plus gateway types). Library only; consumed by ACE backend (`b3`) and downstream operators.

The README is a Kubebuilder scaffold stub; treat this file as the source of truth.

## Architecture

- `api/catalog/v1alpha1/`, `api/gateway/v1alpha1/` — Kubebuilder API types and helpers (`*_types.go`, `groupversion_info.go`, generated `zz_generated.deepcopy.go`).
- `crds/` — generated CRD YAMLs (one per kind, under `catalog.appscode.com` and `gateway.appscode.com`).
- `pkg/lib/`:
  - `gateway.go` — accessor / helper for gateway bindings.
- `PROJECT` — Kubebuilder metadata. Multigroup, domain `appscode.com`.
- `hack/`, `Makefile` — AppsCode build harness (codegen).
- `vendor/` — checked-in deps.

This is a **library** — no Docker image, no install chart, no long-running binary.

## Common commands

- `make gen` — regenerate clientset + manifests after API type changes.
- `make manifests` — regenerate CRDs only.
- `make fmt`, `make lint`, `make unit-tests` / `make test` — standard.
- `make verify` — codegen + module-tidy verification.
- `make add-license` / `make check-license` — manage license headers.

## Conventions

- Module path is `go.bytebuilders.dev/catalog` (vanity URL); imports must use that.
- License: `LICENSE`. Sign off commits (`git commit -s`).
- Vendor directory is checked in; keep `go mod tidy && go mod vendor` clean.
- Do not hand-edit `zz_generated.*.go` or `crds/*.yaml` — change `api/<group>/v1alpha1/*_types.go` and re-run `make gen`.
- This is a **Kubebuilder multigroup project** — use `kubebuilder create api` to scaffold new binding kinds. Don't hand-create files that `PROJECT` should track.
- Adding a new DB binding: it's one `kubebuilder create api --group catalog --kind <Name>Binding` plus `make gen`. The actual binding fields go in the generated `<name>binding_types.go`.
