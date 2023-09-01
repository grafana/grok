export class RangeMapBuilder extends OptionsBuilder<RangeMap> {
	internal: RangeMap;

	build(): RangeMap {
		return this.internal;
	}

	withType(typeArg: string): this {
		if (!(typeArg == "range")) {
			throw new Error("typeArg must be == range");
		}

		this.internal.type = typeArg;

		return this;
	}

	// Range to match against and the result to apply when the value is within the range
	withOptions(options: {
	// Min value of the range. It can be null which means -Infinity
	from: number | null;
	// Max value of the range. It can be null which means +Infinity
	to: number | null;
	// Config to apply when the value is within the range
	result: ValueMappingResult;
}): this {
		
		this.internal.options = options;

		return this;
	}

}
