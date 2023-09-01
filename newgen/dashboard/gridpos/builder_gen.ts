export class GridPosBuilder extends OptionsBuilder<GridPos> {
	internal: GridPos;

	build(): GridPos {
		return this.internal;
	}

	// Panel height. The height is the number of rows from the top edge of the panel.
	withH(h: number): this {
		if (!(h > 0)) {
			throw new Error("h must be > 0");
		}

		this.internal.h = h;

		return this;
	}

	// Panel width. The width is the number of columns from the left edge of the panel.
	withW(w: number): this {
		if (!(w > 0)) {
			throw new Error("w must be > 0");
		}

if (!(w <= 24)) {
			throw new Error("w must be <= 24");
		}

		this.internal.w = w;

		return this;
	}

	// Panel x. The x coordinate is the number of columns from the left edge of the grid
	withX(x: number): this {
		if (!(x >= 0)) {
			throw new Error("x must be >= 0");
		}

if (!(x < 24)) {
			throw new Error("x must be < 24");
		}

		this.internal.x = x;

		return this;
	}

	// Panel y. The y coordinate is the number of rows from the top edge of the grid
	withY(y: number): this {
		if (!(y >= 0)) {
			throw new Error("y must be >= 0");
		}

		this.internal.y = y;

		return this;
	}

	// Whether the panel is fixed within the grid. If true, the panel will not be affected by other panels' interactions
	withStatic(static: boolean): this {
		
		this.internal.static = static;

		return this;
	}

}
