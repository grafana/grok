.PHONY: generate
generate:
	MIN_MATURITY=merged GRAFANA_VERSION=v10.1.0 go generate ./
