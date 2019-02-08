import Util from "../common/Util.mjs";
import Rect from "../common/Rect.mjs";

export default class ShapeView {
    constructor(id, className, parent) {
        this._parent = parent;
        this._element = document.createElement("div");
        this._element.id = `shape${id}`;

        this._element.classList.add("shape");
        this._element.classList.add(className);
        this._element.style.zIndex = id.toString();
        this._rect = null;

        parent.appendChild(this._element);
    }

    remove() {
        this._element.remove();
    }

    doOnMouseDown(cb) {
        this._element.addEventListener("mousedown", (event) => {
            const pos = Util.extractMousePosition(event, Util.extractElementSize(this._parent));
            cb(pos);
        });
    }

    doOnResize(cb) {
        let prevSize = this.rect.size;

        let observer = new MutationObserver(() => {
            const newSize = this.rect.size;
            if (prevSize.equals(newSize)) {
                return;
            }

            cb(newSize);
            prevSize = newSize;
        });

        observer.observe(this._element, {attributes: true, attributeFilter: ['style']});
    }

    get rect() {
        const elemRect = this._element.getBoundingClientRect();
        const parentRect = this._parent.getBoundingClientRect();

        return new Rect(
            (elemRect.left - parentRect.left) / parentRect.width,
            (elemRect.top - parentRect.top) / parentRect.height,
            elemRect.width / parentRect.width,
            elemRect.height / parentRect.height
        );
    }

    set rect(rect) {
        this._rect = rect.clone();
        this._reposition();
    }

    get selected() {
        return this._element.classList.contains("shape_selected");
    }

    set selected(selected) {
        if (selected === this.selected) {
            return;
        }

        if (selected) {
            this._element.classList.add("shape_selected");
        } else {
            this._element.classList.remove("shape_selected");
        }
        this._reposition();
    }

    _reposition() {
        const parentRect = this._parent.getBoundingClientRect();
        const offset = ShapeView._getOffset(this._parent);

        this._element.style.left = `${offset.left + parentRect.width * this._rect.left}px`;
        this._element.style.top = `${offset.top + parentRect.height * this._rect.top}px`;

        const border = +getComputedStyle(this._element).borderTopWidth.slice(0, -2);

        this._element.style.width = `${parentRect.width * this._rect.width - 2 * border}px`;
        this._element.style.height = `${parentRect.height * this._rect.height - 2 * border}px`;
    }

    static _getOffset(elem) {
        const rect = elem.getBoundingClientRect();
        return {
            left: rect.left + window.scrollX,
            top: rect.top + window.scrollY,
        }
    }
}