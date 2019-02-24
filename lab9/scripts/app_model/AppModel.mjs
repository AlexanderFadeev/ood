import History from "./history/History.mjs";
import DocumentAppModel from "./DocumentAppModel.mjs";
import AddShapeCommand from "./history/AddShapeCommand.mjs";
import RemoveShapeCommand from "./history/RemoveShapeCommand.mjs";
import Rect from "../common/Rect.mjs";

const maxHistorySize = 64;

export default class AppModel {
    constructor() {
        this._history = new History(maxHistorySize);
        this._document = new DocumentAppModel(this._history);
    }

    loadFile(file) {
        this._document.reset();

        const data = JSON.parse(file);

        data.shapes.forEach((shape) => {
            this._loadShape(shape);
        });

        this.resetHistory();
    }

    _loadShape(shape) {
        Object.setPrototypeOf(shape.rect, Rect.prototype);
        this.addShape(shape.type, shape.rect);
    }

    undo() {
        this._history.undo();
    }

    redo() {
        this._history.redo();
    }

    canUndo() {
        return this._history.canUndo();
    }

    canRedo() {
        return this._history.canRedo();
    }

    resetHistory() {
        this._history.reset();
    }

    addShape(type, rect) {
        this._history.addAndExecute(new AddShapeCommand(this._document, type, rect));
    }

    removeShape(id) {
        this._history.addAndExecute(new RemoveShapeCommand(this._document, id));
    }

    get onHistoryUpdate() {
        return this._history.onUpdate;
    }

    get onShapeAdded() {
        return this._document.onShapeAdded;
    }

    get onShapeRemoved() {
        return this._document.onShapeRemoved;
    }

    getShape(id) {
        return this._document.getShape(id);
    }

    serialize() {
        return this._document.serialize()
    }
}
