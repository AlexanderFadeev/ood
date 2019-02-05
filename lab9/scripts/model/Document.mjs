export default class Document {
    constructor() {
        this.shapes = [];

    }

    addShape(shape) {
        this.shapes.push(shape);
    }
}