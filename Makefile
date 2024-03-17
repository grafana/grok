.PHONY: generate
generate:
	MIN_MATURITY=merged GRAFANA_VERSION=v10.3.0 go generate ./
