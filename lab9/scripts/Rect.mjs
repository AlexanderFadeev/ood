export default class Rect {
    constructor(left, top, width, height) {
        this.left = left;
        this.top = top;
        this.width = width;
        this.height = height;
    }

    equal(rect) {
        return this.left === rect.left &&
            this.top === rect.top &&
            this.width === rect.width &&
            this.height === rect.height;
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
}