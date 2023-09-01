import * as types from "../valuemap_types_gen";
import { OptionsBuilder } from "../options_builder_gen";

export class ValueMapBuilder implements OptionsBuilder<types.ValueMap> {
	internal: types.ValueMap;

	build(): types.ValueMap {
		return this.internal;
	}

	withType(type: string): this {
		if (!(type == "value")) {
			throw new Error("type must be == value");
		}

		this.internal.type = type;

		return this;
	}

	// Map with <value_to_match>: ValueMappingResult. For example: { "10": { text: "Perfection!", color: "green" } }
	withOptions(options: Record<string, types.ValueMappingResult>): this {
		
		this.internal.options = options;

		return this;
	}

}
