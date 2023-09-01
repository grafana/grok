export class RegexMapBuilder extends OptionsBuilder<RegexMap> {
	internal: RegexMap;

	build(): RegexMap {
		return this.internal;
	}

	withType(typeArg: string): this {
		if (!(typeArg == "regex")) {
			throw new Error("typeArg must be == regex");
		}

		this.internal.type = typeArg;

		return this;
	}

	// Regular expression to match against and the result to apply when the value matches the regex
	withOptions(options: {
	// Regular expression to match against
	pattern: string;
	// Config to apply when the value matches the regex
	result: ValueMappingResult;
}): this {
		
		this.internal.options = options;

		return this;
	}

}
