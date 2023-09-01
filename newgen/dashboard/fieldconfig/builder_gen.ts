import * as types from "../fieldconfig_types_gen";
import { OptionsBuilder } from "../options_builder_gen";

export class FieldConfigBuilder implements OptionsBuilder<types.FieldConfig> {
	internal: types.FieldConfig;

	build(): types.FieldConfig {
		return this.internal;
	}

	// The display value for this field.  This supports template variables blank is auto
	withDisplayName(displayName: string): this {
		
		this.internal.displayName = displayName;

		return this;
	}

	// This can be used by data sources that return and explicit naming structure for values and labels
	// When this property is configured, this value is used rather than the default naming strategy.
	withDisplayNameFromDS(displayNameFromDS: string): this {
		
		this.internal.displayNameFromDS = displayNameFromDS;

		return this;
	}

	// Human readable field metadata
	withDescription(description: string): this {
		
		this.internal.description = description;

		return this;
	}

	// An explicit path to the field in the datasource.  When the frame meta includes a path,
	// This will default to `${frame.meta.path}/${field.name}
	// 
	// When defined, this value can be used as an identifier within the datasource scope, and
	// may be used to update the results
	withPath(path: string): this {
		
		this.internal.path = path;

		return this;
	}

	// True if data source can write a value to the path. Auth/authz are supported separately
	withWriteable(writeable: boolean): this {
		
		this.internal.writeable = writeable;

		return this;
	}

	// True if data source field supports ad-hoc filters
	withFilterable(filterable: boolean): this {
		
		this.internal.filterable = filterable;

		return this;
	}

	// Unit a field should use. The unit you select is applied to all fields except time.
	// You can use the units ID availables in Grafana or a custom unit.
	// Available units in Grafana: https://github.com/grafana/grafana/blob/main/packages/grafana-data/src/valueFormats/categories.ts
	// As custom unit, you can use the following formats:
	// `suffix:<suffix>` for custom unit that should go after value.
	// `prefix:<prefix>` for custom unit that should go before value.
	// `time:<format>` For custom date time formats type for example `time:YYYY-MM-DD`.
	// `si:<base scale><unit characters>` for custom SI units. For example: `si: mF`. This one is a bit more advanced as you can specify both a unit and the source data scale. So if your source data is represented as milli (thousands of) something prefix the unit with that SI scale character.
	// `count:<unit>` for a custom count unit.
	// `currency:<unit>` for custom a currency unit.
	withUnit(unit: string): this {
		
		this.internal.unit = unit;

		return this;
	}

	// Specify the number of decimals Grafana includes in the rendered value.
	// If you leave this field blank, Grafana automatically truncates the number of decimals based on the value.
	// For example 1.1234 will display as 1.12 and 100.456 will display as 100.
	// To display all decimals, set the unit to `String`.
	withDecimals(decimals: number): this {
		
		this.internal.decimals = decimals;

		return this;
	}

	// The minimum value used in percentage threshold calculations. Leave blank for auto calculation based on all series and fields.
	withMin(min: number): this {
		
		this.internal.min = min;

		return this;
	}

	// The maximum value used in percentage threshold calculations. Leave blank for auto calculation based on all series and fields.
	withMax(max: number): this {
		
		this.internal.max = max;

		return this;
	}

	// Convert input values into a display string
	withMappings(mappings: types.ValueMap | types.RangeMap | types.RegexMap | types.SpecialValueMap[]): this {
		
		this.internal.mappings = mappings;

		return this;
	}

	withThresholds(builder: OptionsBuilder<types.ThresholdsConfig>): this {
		this.internal.thresholds = builder.build();

		return this;
	}

	withColor(builder: OptionsBuilder<types.FieldColor>): this {
		this.internal.color = builder.build();

		return this;
	}

	// The behavior when clicking on a result
	withLinks(links: any[]): this {
		
		this.internal.links = links;

		return this;
	}

	// Alternative to empty string
	withNoValue(noValue: string): this {
		
		this.internal.noValue = noValue;

		return this;
	}

	// custom is specified by the FieldConfig field
	// in panel plugin schemas.
	withCustom(custom: any): this {
		
		this.internal.custom = custom;

		return this;
	}

}
