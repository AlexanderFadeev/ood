export default class ShapePresenter {
    constructor(id, model, documentView) {
        this._id = id;
        this._documentView = documentView;

        this._shape = model.getShape(id);

        documentView.addShape(id, this._shape.type, this._shape.rect);
        this._shapeView = documentView.getShape(id);

        this._loadShapeRect();

        this._isDragged = false;
        this._startMousePos = null;

        this._isResized = false;
        this._ignoreNextResize = false;

        documentView.doOnMouseMove(this._handleMouseMove.bind(this));
        documentView.doOnMouseUp(this._handleMouseUp.bind(this));
        this._shapeView.doOnResize(this._handleResize.bind(this));
        this._shapeView.doOnMouseDown(this._handleMouseDown.bind(this));

        this._onRectChangedConn = this._shape.onRectChanged.connect(this._onRectChanged.bind(this));
    }

    onWindowResize() {
        this._loadShapeRect();
        this._syncView();
    }

    remove() {
        this._onRectChangedConn.disconnect();
        this._documentView.removeShape(this._id);
    }

    _onRectChanged() {
        this._loadShapeRect();
        this._syncView();
    }

    _loadShapeRect() {
        this._rect = this._shape.rect.clone();
    }

    _saveShapeRect() {
        this._shape.rect = this._rect.clone();
    }

    _syncView() {
        this._shapeView.rect = this._rect;
        this._ignoreNextResize = true;
    }

    _handleMouseMove(pos) {
        if (!this._isDragged || this._isResized) {
            return;
        }

        this._loadShapeRect();

        const delta = pos.sub(this._startMousePos);
        this._rect.left += delta.x;
        this._rect.top += delta.y;

        if (this._rect.left < 0) {
            this._rect.left = 0;
        }
        if (this._rect.top < 0) {
            this._rect.top = 0;
        }
        if (this._rect.right > 1) {
            this._rect.left -= this._rect.right - 1;
        }
        if (this._rect.bottom > 1) {
            this._rect.top -= this._rect.bottom - 1;
        }

        this._syncView();
    }

    _handleMouseUp(pos) {
        if (!this._isDragged && !this._isResized) {
            return;
        }

        this._handleMouseMove(pos);
        this._saveShapeRect();
        this._isDragged = false;
        this._isResized = false;
    }

    _handleMouseDown(pos) {
        this._isDragged = true;
        this._startMousePos = pos;
    }

    _handleResize(size) {
        if (this._ignoreNextResize) {
            this._ignoreNextResize = false;
            return;
        }

        this._isResized = true;

        this._rect.width = size.x;
        this._rect.height = size.y;

        if (this._rect.right > 1) {
            this._rect.right = 1;
        }
        if (this._rect.bottom > 1) {
            this._rect.bottom = 1;
        }

        this._syncView();
    }
}