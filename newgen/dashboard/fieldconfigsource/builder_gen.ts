import * as types from "../fieldconfigsource_types_gen";
import { OptionsBuilder } from "../options_builder_gen";

export class FieldConfigSourceBuilder implements OptionsBuilder<types.FieldConfigSource> {
	internal: types.FieldConfigSource;

	build(): types.FieldConfigSource {
		return this.internal;
	}

	withDefaults(builder: OptionsBuilder<types.FieldConfig>): this {
		this.internal.defaults = builder.build();

		return this;
	}

	// Overrides are the options applied to specific fields overriding the defaults.
	withOverrides(overrides: types.FieldConfigSourceOverride[]): this {
		
		this.internal.overrides = overrides;

		return this;
	}

}
