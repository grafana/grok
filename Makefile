.PHONY: generate
generate:
	MIN_MATURITY=merged GRAFANA_VERSION=v9.4.0 go generate ./
