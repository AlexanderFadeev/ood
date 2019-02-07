import Shape from "../../model/Shape.mjs";

export default class AddShapeCommand {
    constructor(model, document, type, rect) {
        this._model = model;
        this._document = document;
        this._type = type;
        this._rect = rect.clone();
    }

    execute() {
        if (!this._shape) {
            this._id = this._document.addShape(new Shape(this._type, this._rect));
            this._shape = this._document.getShape(this._id);
        } else {
            this._document.retrieveShape(this._shape, this._id);
        }

        this._model.onShapeAdded.emit(this._id);
    }

    unexecute() {
        this._document.removeShape(this._id);
        this._model.onShapeRemoved.emit(this._id);
    }
}