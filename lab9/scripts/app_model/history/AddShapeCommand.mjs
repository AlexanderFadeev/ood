import Shape from "../../model/Shape.mjs";

export default class AddShapeCommand {
    constructor(shapes, type, rect) {
        this._shapes = shapes;
        this._type = type;
        this._rect = rect.clone();
    }

    execute() {
        if (!this._shape) {
            this._id = this._shapes.addShape(new Shape(this._type, this._rect));
            this._shape = this._shapes.getShape(this._id);
        } else {
            this._shapes.retrieveShape(this._shape, this._id);
        }
    }

    unexecute() {
        this._shapes.removeShape(this._id);
    }
}