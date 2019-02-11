export default class RemoveShapeCommand {
    constructor(document, id) {
        this._document = document;
        this._id = id;
        this._shape = this._document.getShape(id);
    }

    execute() {
        this._document.removeShape(this._id);
    }

    unexecute() {
        this._document.retrieveShape(this._shape, this._id);
    }
}