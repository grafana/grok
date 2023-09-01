import * as types from "../datasourceref_types_gen";
import { OptionsBuilder } from "../options_builder_gen";

export class DataSourceRefBuilder implements OptionsBuilder<types.DataSourceRef> {
	internal: types.DataSourceRef;

	build(): types.DataSourceRef {
		return this.internal;
	}

	// The plugin type-id
	withType(type: string): this {
		
		this.internal.type = type;

		return this;
	}

	// Specific datasource instance
	withUid(uid: string): this {
		
		this.internal.uid = uid;

		return this;
	}

}
