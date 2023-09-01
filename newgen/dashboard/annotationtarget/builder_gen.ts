import * as types from "../annotationtarget_types_gen";
import { OptionsBuilder } from "../options_builder_gen";

export class AnnotationTargetBuilder implements OptionsBuilder<types.AnnotationTarget> {
	internal: types.AnnotationTarget;

	build(): types.AnnotationTarget {
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
	withType(type: string): this {
		
		this.internal.type = type;

		return this;
	}

}
