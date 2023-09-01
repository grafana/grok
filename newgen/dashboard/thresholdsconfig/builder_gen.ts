export class ThresholdsConfigBuilder extends OptionsBuilder<ThresholdsConfig> {
	internal: ThresholdsConfig;

	build(): ThresholdsConfig {
		return this.internal;
	}

	// Thresholds mode.
	withMode(mode: ThresholdsMode): this {
		
		this.internal.mode = mode;

		return this;
	}

	// Must be sorted by 'value', first value is always -Infinity
	withSteps(steps: Threshold[]): this {
		
		this.internal.steps = steps;

		return this;
	}

}
