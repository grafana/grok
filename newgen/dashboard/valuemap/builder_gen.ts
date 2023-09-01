export class ValueMapBuilder extends OptionsBuilder<ValueMap> {
	internal: ValueMap;

	build(): ValueMap {
		return this.internal;
	}

	withType(typeArg: string): this {
		if (!(typeArg == "value")) {
			throw new Error("typeArg must be == value");
		}

		this.internal.type = typeArg;

		return this;
	}

	// Map with <value_to_match>: ValueMappingResult. For example: { "10": { text: "Perfection!", color: "green" } }
	withOptions(options: Record<&{string []}, ValueMappingResult>): this {
		
		this.internal.options = options;

		return this;
	}

}
