import * as types from "../variablemodel_types_gen";
import { OptionsBuilder } from "../options_builder_gen";

export class VariableModelBuilder implements OptionsBuilder<types.VariableModel> {
	internal: types.VariableModel;

	build(): types.VariableModel {
		return this.internal;
	}

	// Unique numeric identifier for the variable.
	withId(id: string): this {
		
		this.internal.id = id;

		return this;
	}

	// Type of variable
	withType(type: types.VariableType): this {
		
		this.internal.type = type;

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
	withHide(hide: types.VariableHide): this {
		
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

	withDatasource(builder: OptionsBuilder<types.DataSourceRef>): this {
		this.internal.datasource = builder.build();

		return this;
	}

	// Format to use while fetching all values from data source, eg: wildcard, glob, regex, pipe, etc.
	withAllFormat(allFormat: string): this {
		
		this.internal.allFormat = allFormat;

		return this;
	}

	withCurrent(builder: OptionsBuilder<types.VariableOption>): this {
		this.internal.current = builder.build();

		return this;
	}

	// Whether multiple values can be selected or not from variable value list
	withMulti(multi: boolean): this {
		
		this.internal.multi = multi;

		return this;
	}

	// Options that can be selected for a variable.
	withOptions(options: types.VariableOption[]): this {
		
		this.internal.options = options;

		return this;
	}

	withRefresh(refresh: types.VariableRefresh): this {
		
		this.internal.refresh = refresh;

		return this;
	}

}
