import History from "./history/History.mjs";
import Signal from "../common/Signal.mjs";
import DocumentAppModel from "./DocumentAppModel.mjs";
import AddShapeCommand from "./history/AddShapeCommand.mjs";
import RemoveShapeCommand from "./history/RemoveShapeCommand.mjs";
import Rect from "../common/Rect.mjs";
import Shape from "../model/Shape.mjs";

const maxHistorySize = 64;

export default class AppModel {
    constructor() {
        this._history = new History(maxHistorySize);
        this._document = new DocumentAppModel(this._history);
        this.onHistoryUpdate = new Signal();
    }

    loadFile(file) {
        this._document.reset();

        const data = JSON.parse(file);
        const shapes = data.shapes;

        for (let shape of shapes.values()) {
            this._loadShape(shape);
        }

        this.resetHistory();
    }

    _loadShape(shape) {
        const rectData = shape.rect;
        const rect = new Rect(
            rectData.left,
            rectData.top,
            rectData.width,
            rectData.height
        );

        this.addShape(shape.type, rect);
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
        this._history.addAndExecute(new AddShapeCommand(this._document, type, rect));
        this.onHistoryUpdate.emit();
    }

    removeShape(id) {
        this._history.addAndExecute(new RemoveShapeCommand(this._document, id));
        this.onHistoryUpdate.emit();
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