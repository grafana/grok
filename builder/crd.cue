package builder

import	(
	"github.com/grafana/thema"
"github.com/grafana/grafana/pkg/coremodel/dashboard"
)

Sch=_sch: (thema.#Pick & {lin: dashboard, v: [0]}).out
#CRD: {
	apiVersion: string
	kind:       string
	metadata: {
		name: string
		...
	}
	spec: {...}
}

#DashboardCR: #CRD & {
	apiVersion: "grafana.com/v1alpha1"
	kind:       "Dashboard"
	spec:       Sch
}

#PanelCR: #CRD & {
	apiVersion: "grafana.com/v1alpha1"
	kind:       "Panel"
	spec:       Sch.#Panel
}

#TemplateCR: #CRD & {
	apiVersion: "grafana.com/v1alpha1"
	kind:       "Template"
}

#QueryCR:           #CRD & {
	apiVersion: "grafana.com/v1alpha1"
	kind: "Query"
}
#TargetCR:          #CRD & {
	apiVersion: "grafana.com/v1alpha1"
	kind: "Target"
}
#PanelDefaultsCR:   #CRD & {
	apiVersion: "grafana.com/v1alpha1"
	kind: "PanelDefaults"
}
#ComposedDashboardCR: #CRD & {
	apiVersion: "grafana.com/v1alpha1"
	kind: "ComposedDashboard"
}
#FolderCR:            #CRD & {
	apiVersion: "grafana.com/v1alpha1"
	kind: "Folder"
}

// Can we express the G-K as a template? oof
[X= =~"[A-Za-z]*CR$"]: #CRD & {
	apiVersion: "grafana.com/v1alpha1"
}

#PanelReference: {
	kind: "Panel"
	name: string
	overrides: Sch.#Panel
}