# grok
Grok (**Gr**afana **O**bject development **K**it) is a CLI tool and Go library for working with Grafana objects from code, primarily dashboards.

# Maturity
> _The code in this repository should be considered experimental. Documentation is only
available alongside the code. It comes with no support, but we are keen to receive
feedback on the product and suggestions on how to improve it, though we cannot commit
to resolution of any particular issue. No SLAs are available. It is not meant to be used
in production environments, and the risks are unknown/high._

Grafana Labs defines experimental features as follows:

> Projects and features in the Experimental stage are supported only by the Engineering
teams; on-call support is not available. Documentation is either limited or not provided
outside of code comments. No SLA is provided.
>
> Experimental projects or features are primarily intended for open source engineers who
want to participate in ensuring systems stability, and to gain consensus and approval
for open source governance projects.
>
> Projects and features in the Experimental phase are not meant to be used in production
environments, and the risks are unknown/high.

# Goals
Key goals include:

* [ ] Provide a structured, stable interface to canonical Grafana kind schemas in a preferred language (Go, CUE, Jsonnet)
  * [ ] Go
  * [ ] CUE
  * [ ] Jsonnet
* [ ] Creating a dashboard/object in code, from scratch.
* [ ] Ingesting a UI-created dashboard/object, transforming it, and emitting the result.
* [ ] Composing a dashboard JSON object from a set of convenient, predefined parts (e.g. queries, panels, template variables) expressed in text form (JSON, YAML)
  * [ ] Decomposing a dashboard JSON object into the above parts
* [ ] Extracting a ["partial"](#Partials) from a dashboard/object into a portable, reusable text format ([CUE](https://cuelang.org))
* [ ] Incorporating a partial fragment into another dashboard/object, either as:
  * [ ] Defaults (values in partial fragments may be overwritten)
  * [ ] Policy/Constraints (values in partial fragments are enforced)
