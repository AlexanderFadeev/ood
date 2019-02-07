import LambdaCommand from "./history/LambdaCommand.mjs";
import Signal from "../common/Signal.mjs";

export default class ShapeAppModel {
    constructor(shape, history) {
        this.impl = shape;
        this._history = history;
        this.onRectChanged = new Signal();
    }

    get rect() {
        return this.impl.rect;
    }

    set rect(rect) {
        rect = rect.clone();
        const oldRect = this.impl.rect.clone();
        this._history.addAndExecute(new LambdaCommand(() => {
            this.impl.rect = rect.clone();
            this.onRectChanged.emit();
        }, () => {
            this.impl.rect = oldRect.clone();
            this.onRectChanged.emit();
        }))
    }

    get type() {
        return this.impl.type;
    }

    set type(type) {
        this.impl.type = type;
    }
}