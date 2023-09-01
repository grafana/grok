export class ThresholdBuilder extends OptionsBuilder<Threshold> {
	internal: Threshold;

	build(): Threshold {
		return this.internal;
	}

	// Value represents a specified metric for the threshold, which triggers a visual change in the dashboard when this value is met or exceeded.
	// Nulls currently appear here when serializing -Infinity to JSON.
	withValue(value: number | null): this {
		
		this.internal.value = value;

		return this;
	}

	// Color represents the color of the visual change that will occur in the dashboard when the threshold value is met or exceeded.
	withColor(color: string): this {
		
		this.internal.color = color;

		return this;
	}

}
