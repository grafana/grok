export class MatcherConfigBuilder extends OptionsBuilder<MatcherConfig> {
	internal: MatcherConfig;

	build(): MatcherConfig {
		return this.internal;
	}

	// The matcher id. This is used to find the matcher implementation from registry.
	withId(id: string): this {
		
		this.internal.id = id;

		return this;
	}

	// The matcher options. This is specific to the matcher implementation.
	withOptions(options: any): this {
		
		this.internal.options = options;

		return this;
	}

}
