import * as types from "../dashboardlink_types_gen";
import { OptionsBuilder } from "../options_builder_gen";

export class DashboardLinkBuilder implements OptionsBuilder<types.DashboardLink> {
	internal: types.DashboardLink;

	build(): types.DashboardLink {
		return this.internal;
	}

	// Title to display with the link
	withTitle(title: string): this {
		
		this.internal.title = title;

		return this;
	}

	// Link type. Accepted values are dashboards (to refer to another dashboard) and link (to refer to an external resource)
	withType(type: types.DashboardLinkType): this {
		
		this.internal.type = type;

		return this;
	}

	// Icon name to be displayed with the link
	withIcon(icon: string): this {
		
		this.internal.icon = icon;

		return this;
	}

	// Tooltip to display when the user hovers their mouse over it
	withTooltip(tooltip: string): this {
		
		this.internal.tooltip = tooltip;

		return this;
	}

	// Link URL. Only required/valid if the type is link
	withUrl(url: string): this {
		
		this.internal.url = url;

		return this;
	}

	// List of tags to limit the linked dashboards. If empty, all dashboards will be displayed. Only valid if the type is dashboards
	withTags(tags: string[]): this {
		
		this.internal.tags = tags;

		return this;
	}

	// If true, all dashboards links will be displayed in a dropdown. If false, all dashboards links will be displayed side by side. Only valid if the type is dashboards
	withAsDropdown(asDropdown: boolean): this {
		
		this.internal.asDropdown = asDropdown;

		return this;
	}

	// If true, the link will be opened in a new tab
	withTargetBlank(targetBlank: boolean): this {
		
		this.internal.targetBlank = targetBlank;

		return this;
	}

	// If true, includes current template variables values in the link as query params
	withIncludeVars(includeVars: boolean): this {
		
		this.internal.includeVars = includeVars;

		return this;
	}

	// If true, includes current time range in the link as query params
	withKeepTime(keepTime: boolean): this {
		
		this.internal.keepTime = keepTime;

		return this;
	}

}
