import ShapeView from "./ShapeView.mjs";
import Util from "./Util.mjs";

export default class DocumentView {
    constructor() {
        this._shapeViews = new Map();
        this.onWindowResize();
    }

    doOnMouseMove(cb) {
        document.addEventListener("mousemove", (event) => {
            const pos = Util.extractMousePosition(event, Util.extractElementSize(this._editorSpace));
            cb(pos);
        });
    }

    doOnMouseUp(cb) {
        document.addEventListener("mouseup", (event) => {
            const pos = Util.extractMousePosition(event, Util.extractElementSize(this._editorSpace));
            cb(pos);
        })
    }

    onWindowResize() {
        let editorSpace = document.getElementById("editorSpace");
        const w = editorSpace.clientWidth;
        const h = w * 9 / 16;
        editorSpace.style.height = `${h}px`;
    }

    addShape(id, type, rect) {
        let shapeView = new ShapeView(id, type, this._editorSpace);
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
        this._editorSpace.addEventListener("mousedown", (event) => {
            if (event.target.id === "editorSpace") {
                handler(null);
                return;
            }

            const id = parseInt(event.target.id.substr(5));
            handler(id);
        })
    }

    get _editorSpace() {
        return document.getElementById("editorSpace");
    }
}