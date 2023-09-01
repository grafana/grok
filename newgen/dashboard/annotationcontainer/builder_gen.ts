import * as types from "../annotationcontainer_types_gen";
import { OptionsBuilder } from "../options_builder_gen";

export class AnnotationContainerBuilder implements OptionsBuilder<types.AnnotationContainer> {
	internal: types.AnnotationContainer;

	build(): types.AnnotationContainer {
		return this.internal;
	}

	// List of annotations
	withList(list: types.AnnotationQuery[]): this {
		
		this.internal.list = list;

		return this;
	}

}
