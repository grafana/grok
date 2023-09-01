export class SpecialValueMapBuilder extends OptionsBuilder<SpecialValueMap> {
	internal: SpecialValueMap;

	build(): SpecialValueMap {
		return this.internal;
	}

	withType(typeArg: string): this {
		if (!(typeArg == "special")) {
			throw new Error("typeArg must be == special");
		}

		this.internal.type = typeArg;

		return this;
	}

	withOptions(options: {
	// Special value to match against
	match: SpecialValueMatch;
	// Config to apply when the value matches the special value
	result: ValueMappingResult;
}): this {
		
		this.internal.options = options;

		return this;
	}

}
