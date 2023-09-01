import * as types from "../fieldcolor_types_gen";
import { OptionsBuilder } from "../options_builder_gen";

export class FieldColorBuilder implements OptionsBuilder<types.FieldColor> {
	internal: types.FieldColor;

	build(): types.FieldColor {
		return this.internal;
	}

	// The main color scheme mode.
	withMode(mode: types.FieldColorModeId): this {
		
		this.internal.mode = mode;

		return this;
	}

	// The fixed color value for fixed or shades color modes.
	withFixedColor(fixedColor: string): this {
		
		this.internal.fixedColor = fixedColor;

		return this;
	}

	// Some visualizations need to know how to assign a series color from by value color schemes.
	withSeriesBy(seriesBy: types.FieldColorSeriesByMode): this {
		
		this.internal.seriesBy = seriesBy;

		return this;
	}

}
