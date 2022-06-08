package builder

import (
	"github.com/grafana/grafana/pkg/coremodel/dashboard"
	"github.com/grafana/grok/dashboard-templates/examples/raw-yaml:base"
	"github.com/grafana/thema"
)

Sch=_sch: (thema.#Pick & {lin: dashboard, v: [0]}).out

#DashboardBuilder: {
	#AddPanel: {
		args: {
			name: string
			type: string
		}
		out: Sch.#Panel & args
	}

	panels: [N=string]: #AddPanel & {
		args: name: N
	}

	// Output is an instance of the dashboard schema
	out: Sch
}

#Universe: [Kind=string]: [Name=string]: {kind: Kind, metadata: name: Name}
#Universe: {
	Dashboard: [string]: #DashboardCR
	Query: [string]: #QueryCR
	Target: [string]: #TargetCR
	Panel: [string]: #PanelCR
	Template: [string]: #TemplateCR
	PanelDefaults: [string]: #PanelDefaultsCR
	ComposedDashboard: [string]: #ComposedDashboardCR
	Folder: [string]: #FolderCR
	for x in (base & [...#CRD]) let X = x.kind {(X): (x.metadata.name): x}
}

#assembleReference: {
	p: #PanelReference
	let x = #Universe.Panel[p.name]
	out: p.params | *x
//	out: p.params & x
}