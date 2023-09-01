import * as types from "../librarypanelref_types_gen";
import { OptionsBuilder } from "../options_builder_gen";

export class LibraryPanelRefBuilder implements OptionsBuilder<types.LibraryPanelRef> {
	internal: types.LibraryPanelRef;

	build(): types.LibraryPanelRef {
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
