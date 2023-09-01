import * as types from "../regexmap_types_gen";
import { OptionsBuilder } from "../options_builder_gen";

export class RegexMapBuilder implements OptionsBuilder<types.RegexMap> {
	internal: types.RegexMap;

	build(): types.RegexMap {
		return this.internal;
	}

	withType(type: string): this {
		if (!(type == "regex")) {
			throw new Error("type must be == regex");
		}

		this.internal.type = type;

		return this;
	}

	// Regular expression to match against and the result to apply when the value matches the regex
	withOptions(options: {
	// Regular expression to match against
	pattern: string;
	// Config to apply when the value matches the regex
	result: types.ValueMappingResult;
}): this {
		
		this.internal.options = options;

		return this;
	}

}
