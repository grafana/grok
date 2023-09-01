import * as types from "../datatransformerconfig_types_gen";
import { OptionsBuilder } from "../options_builder_gen";

export class DataTransformerConfigBuilder implements OptionsBuilder<types.DataTransformerConfig> {
	internal: types.DataTransformerConfig;

	build(): types.DataTransformerConfig {
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

	withFilter(builder: OptionsBuilder<types.MatcherConfig>): this {
		this.internal.filter = builder.build();

		return this;
	}

	// Options to be passed to the transformer
	// Valid options depend on the transformer id
	withOptions(options: any): this {
		
		this.internal.options = options;

		return this;
	}

}
