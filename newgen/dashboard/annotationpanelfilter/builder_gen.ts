export class AnnotationPanelFilterBuilder extends OptionsBuilder<AnnotationPanelFilter> {
	internal: AnnotationPanelFilter;

	build(): AnnotationPanelFilter {
		return this.internal;
	}

	// Should the specified panels be included or excluded
	withExclude(exclude: boolean): this {
		
		this.internal.exclude = exclude;

		return this;
	}

	// Panel IDs that should be included or excluded
	withIds(ids: number[]): this {
		
		this.internal.ids = ids;

		return this;
	}

}
