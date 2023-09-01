export class AnnotationTargetBuilder extends OptionsBuilder<AnnotationTarget> {
	internal: AnnotationTarget;

	build(): AnnotationTarget {
		return this.internal;
	}

	// Only required/valid for the grafana datasource...
	// but code+tests is already depending on it so hard to change
	withLimit(limit: number): this {
		
		this.internal.limit = limit;

		return this;
	}

	// Only required/valid for the grafana datasource...
	// but code+tests is already depending on it so hard to change
	withMatchAny(matchAny: boolean): this {
		
		this.internal.matchAny = matchAny;

		return this;
	}

	// Only required/valid for the grafana datasource...
	// but code+tests is already depending on it so hard to change
	withTags(tags: string[]): this {
		
		this.internal.tags = tags;

		return this;
	}

	// Only required/valid for the grafana datasource...
	// but code+tests is already depending on it so hard to change
	withType(typeArg: string): this {
		
		this.internal.type = typeArg;

		return this;
	}

}
