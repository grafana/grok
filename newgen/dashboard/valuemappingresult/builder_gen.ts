import * as types from "../valuemappingresult_types_gen";
import { OptionsBuilder } from "../options_builder_gen";

export class ValueMappingResultBuilder implements OptionsBuilder<types.ValueMappingResult> {
	internal: types.ValueMappingResult;

	build(): types.ValueMappingResult {
		return this.internal;
	}

	// Text to display when the value matches
	withText(text: string): this {
		
		this.internal.text = text;

		return this;
	}

	// Text to use when the value matches
	withColor(color: string): this {
		
		this.internal.color = color;

		return this;
	}

	// Icon to display when the value matches. Only specific visualizations.
	withIcon(icon: string): this {
		
		this.internal.icon = icon;

		return this;
	}

	// Position in the mapping array. Only used internally.
	withIndex(index: number): this {
		
		this.internal.index = index;

		return this;
	}

}
