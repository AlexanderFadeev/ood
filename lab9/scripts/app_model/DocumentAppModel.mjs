import Document from "../model/Document.mjs";
import ShapeAppModel from "./ShapeAppModel.mjs";
import Signal from "../common/Signal.mjs";

export default class DocumentAppModel {
    constructor(history) {
        this._history = history;

        this._document = new Document();
        this._idMap = new Map();
        this._shapes = new Map();
        this._nextShapeID = 0;

        this.onShapeAdded = new Signal();
        this.onShapeRemoved = new Signal();
    }

    addShape(shape) {
        return this.retrieveShape(new ShapeAppModel(shape, this._history), this._nextShapeID++);
    }

    retrieveShape(shape, id) {
        this._document.shapes.push(shape.impl);

        if (id === undefined) {
            id = this._nextShapeID++;
        }

        this._idMap.set(id, this._document.shapes.length - 1);
        this._shapes.set(id, shape);

        this.onShapeAdded.emit(id);
        return id;
    }

    getShape(id) {
        return this._shapes.get(id);
    }

    removeShape(id) {
        this._document.shapes.splice(this._idMap.get(id), 1);

        for (let key of this._idMap.keys()) {
            if (key > id) {
                this._decrement(key);
            }
        }

        this._idMap.delete(id);
        this._shapes.delete(id);
        this.onShapeRemoved.emit(id);
    }

    _decrement(id) {
        const value = this._idMap.get(id);
        this._idMap.set(id, value - 1);
    }

    reset() {
        for (let key of this._idMap.keys()) {
            this.removeShape(key);
        }
    }

    serialize() {
        return JSON.stringify(this._document);
    }
}