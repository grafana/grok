export class FieldColorBuilder extends OptionsBuilder<FieldColor> {
	internal: FieldColor;

	build(): FieldColor {
		return this.internal;
	}

	// The main color scheme mode.
	withMode(mode: FieldColorModeId): this {
		
		this.internal.mode = mode;

		return this;
	}

	// The fixed color value for fixed or shades color modes.
	withFixedColor(fixedColor: string): this {
		
		this.internal.fixedColor = fixedColor;

		return this;
	}

	// Some visualizations need to know how to assign a series color from by value color schemes.
	withSeriesBy(seriesBy: FieldColorSeriesByMode): this {
		
		this.internal.seriesBy = seriesBy;

		return this;
	}

}
