.PHONY: generate
generate:
	MIN_MATURITY=merged GRAFANA_VERSION=v10.4.0 go generate ./
