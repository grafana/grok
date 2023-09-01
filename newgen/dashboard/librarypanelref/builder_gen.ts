export class LibraryPanelRefBuilder extends OptionsBuilder<LibraryPanelRef> {
	internal: LibraryPanelRef;

	build(): LibraryPanelRef {
		return this.internal;
	}

	// Library panel name
	withName(name: string): this {
		
		this.internal.name = name;

		return this;
	}

	// Library panel uid
	withUid(uid: string): this {
		
		this.internal.uid = uid;

		return this;
	}

}
