export class AnnotationContainerBuilder extends OptionsBuilder<AnnotationContainer> {
	internal: AnnotationContainer;

	build(): AnnotationContainer {
		return this.internal;
	}

	// List of annotations
	withList(list: AnnotationQuery[]): this {
		
		this.internal.list = list;

		return this;
	}

}
