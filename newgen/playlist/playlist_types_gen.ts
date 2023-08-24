export enum PlaylistItemType {
	Dashboard_by_tag = "dashboard_by_tag",
	Dashboard_by_uid = "dashboard_by_uid",
}

export interface PlaylistItem {
	type: PlaylistItemType;
	value: string;
}

export interface playlist {
	interval: string;
	items: PlaylistItem[];
	name: string;
	xxx: string;
}

