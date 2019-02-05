import Vector from "./Vector.mjs";

export default class Rect {
    constructor(left, top, width, height) {
        this.left = left;
        this.top = top;
        this.width = width;
        this.height = height;
    }

    equals(rect) {
        return this.left === rect.left &&
            this.top === rect.top &&
            this.width === rect.width &&
            this.height === rect.height;
    }

    clone() {
        return new Rect(this.left, this.top, this.width, this.height);
    }

    get right() {
        return this.left + this.width;
    }

    set right(right) {
        this.width = right - this.left;
    }

    get bottom() {
        return this.top + this.height;
    }

    set bottom(bottom) {
        this.height = bottom - this.top;
    }

    get size() {
        return new Vector(this.width, this.height);
    }
}