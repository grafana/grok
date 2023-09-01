export class AnnotationQueryBuilder extends OptionsBuilder<AnnotationQuery> {
	internal: AnnotationQuery;

	build(): AnnotationQuery {
		return this.internal;
	}

	// Name of annotation.
	withName(name: string): this {
		
		this.internal.name = name;

		return this;
	}

	// Datasource where the annotations data is
	withDatasource(datasource: DataSourceRef): this {
		
		this.internal.datasource = datasource;

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

	// Filters to apply when fetching annotations
	withFilter(filter: AnnotationPanelFilter): this {
		
		this.internal.filter = filter;

		return this;
	}

	// TODO.. this should just be a normal query target
	withTarget(target: AnnotationTarget): this {
		
		this.internal.target = target;

		return this;
	}

	// TODO -- this should not exist here, it is based on the --grafana-- datasource
	withType(typeArg: string): this {
		
		this.internal.type = typeArg;

		return this;
	}

}
