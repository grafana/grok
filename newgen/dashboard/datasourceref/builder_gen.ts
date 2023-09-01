export class DataSourceRefBuilder extends OptionsBuilder<DataSourceRef> {
	internal: DataSourceRef;

	build(): DataSourceRef {
		return this.internal;
	}

	// The plugin type-id
	withType(typeArg: string): this {
		
		this.internal.type = typeArg;

		return this;
	}

	// Specific datasource instance
	withUid(uid: string): this {
		
		this.internal.uid = uid;

		return this;
	}

}
