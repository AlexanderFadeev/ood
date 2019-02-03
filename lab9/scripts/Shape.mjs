import Util from "./Util.mjs";
import Rect from "./Rect.mjs";

export default class Shape {
    constructor(id, className, parent, onRectUpdate) {
        this._id = id;
        this._parent = parent;
        this._element = document.createElement("div");
        this._onUpdate = onRectUpdate;

        this._element.classList.add(className);
        this._element.id = `shape${id}`;
        this._element.style.width = "42px";
        this._element.style.height = "42px";
        this._parent.appendChild(this._element);

        this._makeDraggable();
        this._makeResizeable();
        this._prevDimensions = this._dimensions;

        this.handleMouseMove = Util.throttle(this.handleMouseMove.bind(this), 25);
    }

    handleMouseMove(pos) {
        if (!this._isDragged || this._isResized) {
            return;
        }

        this._element.style.left = `${this._element.offsetLeft + pos.x - this._prevPos.x}px`;
        this._element.style.top = `${this._element.offsetTop + pos.y - this._prevPos.y}px`;

        let dim = this._dimensions;
        if (dim.left < 0) {
            dim.left = 0
        }
        if (dim.top < 0) {
            dim.top = 0
        }
        if (dim.right > 1) {
            dim.left -= this._dimensions.right - 1;
        }
        if (dim.bottom > 1) {
            dim.top -= this._dimensions.bottom - 1;
        }
        this._dimensions = dim;

        this._prevPos = pos;
    }

    onWindowResize() {
        this._dimensions = this._prevDimensions;
    }

    handleMouseUp(pos) {
        this.handleMouseMove(pos);
        this._isDragged = false;
        this._isResized = false;
    }

    _makeDraggable() {
        this._element.classList.add("draggable");

        this._isDragged = false;
        this._prevPos = null;

        this._element.addEventListener("mousedown", (event) => {
            this._isDragged = true;
            this._prevPos = Util.extractMousePosition(event);
        });
    }

    _makeResizeable() {
        this._element.classList.add("resizeable");

        this._isResized = false;

        let observer = new MutationObserver(this._handleResize.bind(this));

        observer.observe(this._element, {attributes: true});
    }

    _handleResize() {
        if (!this._resized) {
            return;
        }
        this._isResized = true;

        let dim = this._dimensions;
        if (dim.right > 1) {
            dim.width -= this._dimensions.right - 1;
        }
        if (dim.bottom > 1) {
            dim.height -= this._dimensions.bottom - 1;
        }
        this._dimensions = dim;
    }

    get _dimensions() {
        const elemRect = this._element.getBoundingClientRect();
        const parentRect = this._parent.getBoundingClientRect();

        return new Rect(
            (elemRect.left - parentRect.left) / parentRect.width,
            (elemRect.top - parentRect.top) / parentRect.height,
            elemRect.width / parentRect.width,
            elemRect.height / parentRect.height
        );
    }

    set _dimensions(d) {
        this._prevDimensions = this._dimensions;

        const rect = this._parent.getBoundingClientRect();
        const offset = Shape._getOffset(this._parent);

        this._element.style.left = `${offset.left + rect.width * d.left}px`;
        this._element.style.top = `${offset.top + rect.height * d.top}px`;

        this._element.style.width = `${rect.width * d.width - 2}px`;
        this._element.style.height = `${rect.height * d.height - 2}px`;

        this._update();
    }

    get _resized() {
        const dimensions = this._dimensions;

        return dimensions.width !== this._prevDimensions.width ||
            dimensions.height !== this._prevDimensions.height;
    }

    _update() {
        const dimensions = this._dimensions;

        if (dimensions.equal(this._prevDimensions)) {
            return;
        }

        this._onUpdate(this._id, dimensions);
    }

    static _getOffset(elem) {
        const rect = elem.getBoundingClientRect();
        return {
            left: rect.left + window.scrollX,
            top: rect.top + window.scrollY,
        }
    }
}