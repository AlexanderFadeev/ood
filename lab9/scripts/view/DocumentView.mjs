import ShapeView from "./ShapeView.mjs";
import Util from "./Util.mjs";

export default class DocumentView {
    constructor() {
        this._element = document.getElementById("document");
        this._shapeViews = new Map();
        this.onWindowResize();
    }

    doOnMouseMove(cb) {
        document.addEventListener("mousemove", (event) => {
            const pos = Util.extractMousePosition(event, Util.extractElementSize(this._element));
            cb(pos);
        });
    }

    doOnMouseUp(cb) {
        document.addEventListener("mouseup", (event) => {
            const pos = Util.extractMousePosition(event, Util.extractElementSize(this._element));
            cb(pos);
        })
    }

    onWindowResize() {
        const w = this._element.clientWidth;
        const h = w * 9 / 16;
        this._element.style.height = `${h}px`;
    }

    addShape(id, type, rect) {
        let shapeView = new ShapeView(id, type, this._element);
        shapeView.rect = rect;
        this._shapeViews.set(id, shapeView);
    }

    getShape(id) {
        return this._shapeViews.get(id);
    }

    removeShape(id) {
        this._shapeViews.get(id).remove();
        this._shapeViews.delete(id);
    }

    doOnShapeMouseDown(handler) {
        this._element.addEventListener("mousedown", (event) => {
            if (event.target.id === "editorSpace") {
                handler(null);
                return;
            }

            const id = parseInt(event.target.id.substr(5));
            handler(id);
        })
    }
}