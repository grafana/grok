.PHONY: generate
generate:
	MIN_MATURITY=merged GRAFANA_VERSION=v10.0.0 go generate ./
