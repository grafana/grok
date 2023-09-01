export class FieldConfigSourceBuilder extends OptionsBuilder<FieldConfigSource> {
	internal: FieldConfigSource;

	build(): FieldConfigSource {
		return this.internal;
	}

	// Defaults are the options applied to all fields.
	withDefaults(defaults: FieldConfig): this {
		
		this.internal.defaults = defaults;

		return this;
	}

	// Overrides are the options applied to specific fields overriding the defaults.
	withOverrides(overrides: FieldConfigSourceOverride[]): this {
		
		this.internal.overrides = overrides;

		return this;
	}

}
