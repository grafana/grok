import * as types from "../dashboardtemplating_types_gen";
import { OptionsBuilder } from "../options_builder_gen";

export class DashboardTemplatingBuilder implements OptionsBuilder<types.DashboardTemplating> {
	internal: types.DashboardTemplating;

	build(): types.DashboardTemplating {
		return this.internal;
	}

	// List of configured template variables with their saved values along with some other metadata
	withList(list: types.VariableModel[]): this {
		
		this.internal.list = list;

		return this;
	}

}
