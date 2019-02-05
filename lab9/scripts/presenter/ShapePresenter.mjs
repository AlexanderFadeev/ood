import Shape from "../model/Shape.mjs";
import Rect from "../common/Rect.mjs";

export default class ShapePresenter {
    constructor(type, model, view) {
        this._model = model;
        this._view = view;

        this._shapeModel = new Shape(type, new Rect(1 / 4, 1 / 4, 1 / 2, 1 / 2));
        this._shapeView = this._view.newShapeView(type);

        this._model.addShape(this._shapeModel);

        this._isDragged = false;
        this._isResized = false;
        this._prevMousePos = null;

        this._shapeView.doOnResize(this._handleResize.bind(this));
        this._view.doOnMouseMove(this._handleMouseMove.bind(this));
        this._view.doOnMouseUp(this._handleMouseUp.bind(this));
        this._shapeView.doOnMouseDown(this._handleMouseDown.bind(this));
        this._sync();
    }

    onWindowResize() {
        this._sync();
    }

    _handleMouseMove(pos) {
        if (!this._isDragged || this._isResized) {
            return;
        }

        let rect = this._shapeModel.rect;

        const delta = pos.sub(this._prevMousePos);
        rect.left += delta.x;
        rect.top += delta.y;

        if (rect.left < 0) {
            rect.left = 0
        }
        if (rect.top < 0) {
            rect.top = 0
        }
        if (rect.right > 1) {
            rect.left -= rect.right - 1;
        }
        if (rect.bottom > 1) {
            rect.top -= rect.bottom - 1;
        }

        this._sync();

        this._prevMousePos = pos;
    }

    _handleMouseUp(pos) {
        this._handleMouseMove(pos);
        this._isDragged = false;
        this._isResized = false;
    }

    _handleMouseDown(pos) {
        this._isDragged = true;
        this._prevMousePos = pos;
    }

    _handleResize(size) {
        this._isResized = true;

        let rect = this._shapeModel.rect;

        rect.width = size.x;
        rect.height = size.y;

        if (rect.right > 1) {
            rect.right = 1;
        }
        if (rect.bottom > 1) {
            rect.bottom = 1;
        }
        this._sync();
    }

    _sync() {
        this._shapeView.rect = this._shapeModel.rect;
    }
}