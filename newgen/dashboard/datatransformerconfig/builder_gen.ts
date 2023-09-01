export class DataTransformerConfigBuilder extends OptionsBuilder<DataTransformerConfig> {
	internal: DataTransformerConfig;

	build(): DataTransformerConfig {
		return this.internal;
	}

	// Unique identifier of transformer
	withId(id: string): this {
		
		this.internal.id = id;

		return this;
	}

	// Disabled transformations are skipped
	withDisabled(disabled: boolean): this {
		
		this.internal.disabled = disabled;

		return this;
	}

	// Optional frame matcher. When missing it will be applied to all results
	withFilter(filter: MatcherConfig): this {
		
		this.internal.filter = filter;

		return this;
	}

	// Options to be passed to the transformer
	// Valid options depend on the transformer id
	withOptions(options: any): this {
		
		this.internal.options = options;

		return this;
	}

}
