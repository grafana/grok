import * as types from "../matcherconfig_types_gen";
import { OptionsBuilder } from "../options_builder_gen";

export class MatcherConfigBuilder implements OptionsBuilder<types.MatcherConfig> {
	internal: types.MatcherConfig;

	build(): types.MatcherConfig {
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
