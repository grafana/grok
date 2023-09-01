export class VariableOptionBuilder extends OptionsBuilder<VariableOption> {
	internal: VariableOption;

	build(): VariableOption {
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
