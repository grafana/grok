import * as types from "../annotationpanelfilter_types_gen";
import { OptionsBuilder } from "../options_builder_gen";

export class AnnotationPanelFilterBuilder implements OptionsBuilder<types.AnnotationPanelFilter> {
	internal: types.AnnotationPanelFilter;

	build(): types.AnnotationPanelFilter {
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
