# grok
Grok (**Gr**afana **O**bject development **K**it) is a CLI tool and Go library for working with Grafana objects from code, primarily dashboards.

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
