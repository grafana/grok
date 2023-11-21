.PHONY: generate
generate:
	MIN_MATURITY=merged GRAFANA_VERSION=v10.2.0 go generate ./
