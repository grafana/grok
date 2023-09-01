import * as types from "../specialvaluemap_types_gen";
import { OptionsBuilder } from "../options_builder_gen";

export class SpecialValueMapBuilder implements OptionsBuilder<types.SpecialValueMap> {
	internal: types.SpecialValueMap;

	build(): types.SpecialValueMap {
		return this.internal;
	}

	withType(type: string): this {
		if (!(type == "special")) {
			throw new Error("type must be == special");
		}

		this.internal.type = type;

		return this;
	}

	withOptions(options: {
	// Special value to match against
	match: types.SpecialValueMatch;
	// Config to apply when the value matches the special value
	result: types.ValueMappingResult;
}): this {
		
		this.internal.options = options;

		return this;
	}

}
