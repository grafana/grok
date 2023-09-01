import * as types from "../fieldconfigsourceoverride_types_gen";
import { OptionsBuilder } from "../options_builder_gen";

export class FieldConfigSourceOverrideBuilder implements OptionsBuilder<types.FieldConfigSourceOverride> {
	internal: types.FieldConfigSourceOverride;

	build(): types.FieldConfigSourceOverride {
		return this.internal;
	}

	withMatcher(builder: OptionsBuilder<types.MatcherConfig>): this {
		this.internal.matcher = builder.build();

		return this;
	}

	withProperties(properties: types.DynamicConfigValue[]): this {
		
		this.internal.properties = properties;

		return this;
	}

}
