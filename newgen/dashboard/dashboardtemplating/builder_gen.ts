export class DashboardTemplatingBuilder extends OptionsBuilder<DashboardTemplating> {
	internal: DashboardTemplating;

	build(): DashboardTemplating {
		return this.internal;
	}

	// List of configured template variables with their saved values along with some other metadata
	withList(list: VariableModel[]): this {
		
		this.internal.list = list;

		return this;
	}

}
