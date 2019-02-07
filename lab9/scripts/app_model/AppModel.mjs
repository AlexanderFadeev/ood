import History from "./history/History.mjs";
import Signal from "../common/Signal.mjs";
import DocumentAppModel from "./DocumentAppModel.mjs";
import ShapeAppModel from "./ShapeAppModel.mjs";
import AddShapeCommand from "./history/AddShapeCommand.mjs";
import RemoveShapeCommand from "./history/RemoveShapeCommand.mjs";

const maxHistorySize = 64;

export default class AppModel {
    constructor() {
        this._history = new History(maxHistorySize);
        this._document = new DocumentAppModel(this._history);

        this.onShapeAdded = new Signal();
        this.onShapeRemoved = new Signal();
        this.onHistoryUpdate = new Signal();
    }

    undo() {
        this._history.undo();
        this.onHistoryUpdate.emit();
    }

    redo() {
        this._history.redo();
        this.onHistoryUpdate.emit();
    }

    canUndo() {
        return this._history.canUndo();
    }

    canRedo() {
        return this._history.canRedo();
    }

    resetHistory() {
        this._history.reset();
        this.onHistoryUpdate.emit();
    }

    addShape(type, rect) {
        this._history.addAndExecute(new AddShapeCommand(this, this._document, type, rect));
        this.onHistoryUpdate.emit();
    }

    removeShape(id) {
        this._history.addAndExecute(new RemoveShapeCommand(this, this._document, id));
        this.onHistoryUpdate.emit();
    }

    getShape(id) {
        return this._document.getShape(id);
    }

    serialize() {
        return this._document.serialize()
    }
}