import * as types from "../annotationquery_types_gen";
import { OptionsBuilder } from "../options_builder_gen";

export class AnnotationQueryBuilder implements OptionsBuilder<types.AnnotationQuery> {
	internal: types.AnnotationQuery;

	build(): types.AnnotationQuery {
		return this.internal;
	}

	// Name of annotation.
	withName(name: string): this {
		
		this.internal.name = name;

		return this;
	}

	withDatasource(builder: OptionsBuilder<types.DataSourceRef>): this {
		this.internal.datasource = builder.build();

		return this;
	}

	// When enabled the annotation query is issued with every dashboard refresh
	withEnable(enable: boolean): this {
		
		this.internal.enable = enable;

		return this;
	}

	// Annotation queries can be toggled on or off at the top of the dashboard.
	// When hide is true, the toggle is not shown in the dashboard.
	withHide(hide: boolean): this {
		
		this.internal.hide = hide;

		return this;
	}

	// Color to use for the annotation event markers
	withIconColor(iconColor: string): this {
		
		this.internal.iconColor = iconColor;

		return this;
	}

	withFilter(builder: OptionsBuilder<types.AnnotationPanelFilter>): this {
		this.internal.filter = builder.build();

		return this;
	}

	withTarget(builder: OptionsBuilder<types.AnnotationTarget>): this {
		this.internal.target = builder.build();

		return this;
	}

	// TODO -- this should not exist here, it is based on the --grafana-- datasource
	withType(type: string): this {
		
		this.internal.type = type;

		return this;
	}

}
