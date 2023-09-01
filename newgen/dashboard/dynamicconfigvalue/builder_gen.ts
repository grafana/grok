export class DynamicConfigValueBuilder extends OptionsBuilder<DynamicConfigValue> {
	internal: DynamicConfigValue;

	build(): DynamicConfigValue {
		return this.internal;
	}

	withId(id: string): this {
		
		this.internal.id = id;

		return this;
	}

	withValue(value: any): this {
		
		this.internal.value = value;

		return this;
	}

}
