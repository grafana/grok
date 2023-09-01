import * as types from "../timeinterval_types_gen";
import { OptionsBuilder } from "../options_builder_gen";

export class TimeIntervalBuilder implements OptionsBuilder<types.TimeInterval> {
	internal: types.TimeInterval;

	build(): types.TimeInterval {
		return this.internal;
	}

	withFrom(from: string): this {
		
		this.internal.from = from;

		return this;
	}

	withTo(to: string): this {
		
		this.internal.to = to;

		return this;
	}

}
