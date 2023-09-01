import * as types from "../dynamicconfigvalue_types_gen";
import { OptionsBuilder } from "../options_builder_gen";

export class DynamicConfigValueBuilder implements OptionsBuilder<types.DynamicConfigValue> {
	internal: types.DynamicConfigValue;

	build(): types.DynamicConfigValue {
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
