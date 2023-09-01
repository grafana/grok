import * as types from "../variableoption_types_gen";
import { OptionsBuilder } from "../options_builder_gen";

export class VariableOptionBuilder implements OptionsBuilder<types.VariableOption> {
	internal: types.VariableOption;

	build(): types.VariableOption {
		return this.internal;
	}

	// Whether the option is selected or not
	withSelected(selected: boolean): this {
		
		this.internal.selected = selected;

		return this;
	}

	// Text to be displayed for the option
	withText(text: string | string[]): this {
		
		this.internal.text = text;

		return this;
	}

	// Value of the option
	withValue(value: string | string[]): this {
		
		this.internal.value = value;

		return this;
	}

}
