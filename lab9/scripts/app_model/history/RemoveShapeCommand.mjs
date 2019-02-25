export default class RemoveShapeCommand {
    constructor(shape, id) {
        this._shapes = shape;
        this._id = id;
        this._shape = this._shapes.getShape(id);
    }

    execute() {
        this._shapes.removeShape(this._id);
    }

    unexecute() {
        this._shapes.retrieveShape(this._shape, this._id);
    }
}