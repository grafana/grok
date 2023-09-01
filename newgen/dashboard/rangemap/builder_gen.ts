import * as types from "../rangemap_types_gen";
import { OptionsBuilder } from "../options_builder_gen";

export class RangeMapBuilder implements OptionsBuilder<types.RangeMap> {
	internal: types.RangeMap;

	build(): types.RangeMap {
		return this.internal;
	}

	withType(type: string): this {
		if (!(type == "range")) {
			throw new Error("type must be == range");
		}

		this.internal.type = type;

		return this;
	}

	// Range to match against and the result to apply when the value is within the range
	withOptions(options: {
	// Min value of the range. It can be null which means -Infinity
	from: number | null;
	// Max value of the range. It can be null which means +Infinity
	to: number | null;
	// Config to apply when the value is within the range
	result: types.ValueMappingResult;
}): this {
		
		this.internal.options = options;

		return this;
	}

}
