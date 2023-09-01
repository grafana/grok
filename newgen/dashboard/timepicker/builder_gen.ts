import * as types from "../timepicker_types_gen";
import { OptionsBuilder } from "../options_builder_gen";

export class TimePickerBuilder implements OptionsBuilder<types.TimePicker> {
	internal: types.TimePicker;

	build(): types.TimePicker {
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
