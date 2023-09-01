export class TimePickerBuilder extends OptionsBuilder<TimePicker> {
	internal: TimePicker;

	build(): TimePicker {
		return this.internal;
	}

	// Whether timepicker is visible or not.
	withHidden(hidden: boolean): this {
		
		this.internal.hidden = hidden;

		return this;
	}

	// Interval options available in the refresh picker dropdown.
	withRefreshIntervals(refreshIntervals: string[]): this {
		
		this.internal.refresh_intervals = refreshIntervals;

		return this;
	}

	// Whether timepicker is collapsed or not. Has no effect on provisioned dashboard.
	withCollapse(collapse: boolean): this {
		
		this.internal.collapse = collapse;

		return this;
	}

	// Whether timepicker is enabled or not. Has no effect on provisioned dashboard.
	withEnable(enable: boolean): this {
		
		this.internal.enable = enable;

		return this;
	}

	// Selectable options available in the time picker dropdown. Has no effect on provisioned dashboard.
	withTimeOptions(timeOptions: string[]): this {
		
		this.internal.time_options = timeOptions;

		return this;
	}

}
