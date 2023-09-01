export class VariableModelBuilder extends OptionsBuilder<VariableModel> {
	internal: VariableModel;

	build(): VariableModel {
		return this.internal;
	}

	// Unique numeric identifier for the variable.
	withId(id: string): this {
		
		this.internal.id = id;

		return this;
	}

	// Type of variable
	withType(typeArg: VariableType): this {
		
		this.internal.type = typeArg;

		return this;
	}

	// Name of variable
	withName(name: string): this {
		
		this.internal.name = name;

		return this;
	}

	// Optional display name
	withLabel(label: string): this {
		
		this.internal.label = label;

		return this;
	}

	// Visibility configuration for the variable
	withHide(hide: VariableHide): this {
		
		this.internal.hide = hide;

		return this;
	}

	// Whether the variable value should be managed by URL query params or not
	withSkipUrlSync(skipUrlSync: boolean): this {
		
		this.internal.skipUrlSync = skipUrlSync;

		return this;
	}

	// Description of variable. It can be defined but `null`.
	withDescription(description: string): this {
		
		this.internal.description = description;

		return this;
	}

	// Query used to fetch values for a variable
	withQuery(query: any): this {
		
		this.internal.query = query;

		return this;
	}

	// Data source used to fetch values for a variable. It can be defined but `null`.
	withDatasource(datasource: DataSourceRef): this {
		
		this.internal.datasource = datasource;

		return this;
	}

	// Format to use while fetching all values from data source, eg: wildcard, glob, regex, pipe, etc.
	withAllFormat(allFormat: string): this {
		
		this.internal.allFormat = allFormat;

		return this;
	}

	// Shows current selected variable text/value on the dashboard
	withCurrent(current: VariableOption): this {
		
		this.internal.current = current;

		return this;
	}

	// Whether multiple values can be selected or not from variable value list
	withMulti(multi: boolean): this {
		
		this.internal.multi = multi;

		return this;
	}

	// Options that can be selected for a variable.
	withOptions(options: VariableOption[]): this {
		
		this.internal.options = options;

		return this;
	}

	withRefresh(refresh: VariableRefresh): this {
		
		this.internal.refresh = refresh;

		return this;
	}

}
