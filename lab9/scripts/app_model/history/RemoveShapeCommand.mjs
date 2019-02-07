export default class RemoveShapeCommand {
    constructor(model, document, id) {
        this._model = model;
        this._document = document;
        this._id = id;
        this._shape = this._document.getShape(id);
    }

    execute() {
        this._document.removeShape(this._id);
        this._model.onShapeRemoved.emit(this._id);
    }

    unexecute() {
        this._document.retrieveShape(this._shape, this._id);
        this._model.onShapeAdded.emit(this._id);
    }
}