import Util from "../common/Util.mjs";
import Rect from "../common/Rect.mjs";

export default class ShapeView {
    constructor(className, parent) {
        this._parent = parent;
        this._element = document.createElement("div");

        this._element.classList.add("shape");
        this._element.classList.add(className);

        parent.appendChild(this._element);
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
        const parentRect = this._parent.getBoundingClientRect();
        const offset = ShapeView._getOffset(this._parent);

        this._element.style.left = `${offset.left + parentRect.width * rect.left}px`;
        this._element.style.top = `${offset.top + parentRect.height * rect.top}px`;

        this._element.style.width = `${parentRect.width * rect.width - 2}px`;
        this._element.style.height = `${parentRect.height * rect.height - 2}px`;
    }

    static _getOffset(elem) {
        const rect = elem.getBoundingClientRect();
        return {
            left: rect.left + window.scrollX,
            top: rect.top + window.scrollY,
        }
    }
}