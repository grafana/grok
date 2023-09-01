export class FieldConfigSourceOverrideBuilder extends OptionsBuilder<FieldConfigSourceOverride> {
	internal: FieldConfigSourceOverride;

	build(): FieldConfigSourceOverride {
		return this.internal;
	}

	withMatcher(matcher: MatcherConfig): this {
		
		this.internal.matcher = matcher;

		return this;
	}

	withProperties(properties: DynamicConfigValue[]): this {
		
		this.internal.properties = properties;

		return this;
	}

}
