export class TimeIntervalBuilder extends OptionsBuilder<TimeInterval> {
	internal: TimeInterval;

	build(): TimeInterval {
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
