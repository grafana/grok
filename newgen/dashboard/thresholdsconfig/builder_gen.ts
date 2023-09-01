import * as types from "../thresholdsconfig_types_gen";
import { OptionsBuilder } from "../options_builder_gen";

export class ThresholdsConfigBuilder implements OptionsBuilder<types.ThresholdsConfig> {
	internal: types.ThresholdsConfig;

	build(): types.ThresholdsConfig {
		return this.internal;
	}

	// Thresholds mode.
	withMode(mode: types.ThresholdsMode): this {
		
		this.internal.mode = mode;

		return this;
	}

	// Must be sorted by 'value', first value is always -Infinity
	withSteps(steps: types.Threshold[]): this {
		
		this.internal.steps = steps;

		return this;
	}

}
