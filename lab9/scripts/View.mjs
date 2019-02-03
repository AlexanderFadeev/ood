import Shape from "./Shape.mjs";
import Util from "./Util.mjs";

export const tabFile = 0;
export const tabHome = 1;
export const tabInsert = 2;
export const tabView = 3;
export const tabFormat = 4;

const tabs = [
    "tabFile",
    "tabHome",
    "tabInsert",
    "tabView",
    "tabFormat",
];

export const buttonAddRectangle = 0;
export const buttonAddEllipse = 1;

const buttons = [
    "buttonAddRectangle",
    "buttonAddEllipse",
];

export class View {
    constructor() {
        this._activeTabID = null;
        this._shapes = {};
        this._shapeUpdateHandlers = [];

        this._setActiveTab(tabHome);
        this._onTabClick(this._setActiveTab.bind(this));

        document.addEventListener("mousemove", this._onMouseMove.bind(this));
        document.addEventListener("mouseup", this._onMouseUp.bind(this));
        window.addEventListener("resize", this._onWindowResize.bind(this));

        this._onWindowResize();
    }

    removeLoader() {
        document.getElementById("loader").remove()
    }

    doOnButtonClick(id, handler) {
        this._button(id).addEventListener(`click`, handler);
    }

    _onTabClick(handler) {
        for (let id = 0; id < tabs.length; id++) {
            const idCopy = id;
            this._tab(id).addEventListener('click', () => {
                handler(idCopy);
            });
        }
    }

    _setActiveTab(newActiveTabID) {
        if (newActiveTabID === this._activeTabID) {
            return;
        }

        this._tab(newActiveTabID).classList.replace("tab_inactive", "tab_active");
        if (this._activeTabID !== null) {
            this._tab(this._activeTabID).classList.replace("tab_active", "tab_inactive");
        }

        this._activeTabID = newActiveTabID;
    }

    _tab(id) {
        return document.getElementById(tabs[id]);
    }

    _button(id) {
        return document.getElementById(buttons[id]);
    }

    addRectangle(id) {
        this._addShape(id, "rectangle")
    }

    addEllipse(id) {
        this._addShape(id, "ellipse")
    }

    _addShape(id, className) {
        const parent = document.getElementById("editorSpace");
        this._shapes[id] = new Shape(id, className, parent, this._onShapeUpdate.bind(this));
    }

    doOnShapeUpdate(cb) {
        this._shapeUpdateHandlers.push(cb);
        console.log(this);
    }

    _onShapeUpdate(id, dimensions) {
        for (let handler of this._shapeUpdateHandlers) {
            const d = dimensions;
            handler(id, d.left, d.top, d.width, d.height);
        }
    }

    _onMouseMove(event) {
        const pos = Util.extractMousePosition(event);
        this._forEachShape((shape) => {
            shape.handleMouseMove(pos);
        });
    }

    _onMouseUp(event) {
        const pos = Util.extractMousePosition(event);

        this._forEachShape((shape) => {
            shape.handleMouseUp(pos);
        });
    }


    _onWindowResize() {
        let editorSpace = document.getElementById("editorSpace");
        const w = editorSpace.clientWidth;
        const h = w * 9 / 16;
        editorSpace.style.height = `${h}px`;

        this._forEachShape((shape) => {
            shape.onWindowResize();
        })
    }

    _forEachShape(cb) {
        for (let id in this._shapes) {
            if (!this._shapes.hasOwnProperty(id)) {
                continue;
            }

            let shape = this._shapes[id];
            cb(shape);
        }
    }
}