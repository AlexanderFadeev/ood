import ShapeView from "./ShapeView.mjs";
import Util from "../common/Util.mjs";

export const tabFile = 0;
export const tabInsert = 1;

const tabs = [
    "tabFile",
    "tabInsert",
];

export const menuFile = 0;
export const menuInsert = 1;

const menus = [
    "menuFile",
    "menuInsert",
];

export const buttonOpen = 0;
export const buttonSave = 1;
export const buttonSaveAs = 2;
export const buttonAddRectangle = 3;
export const buttonAddEllipse = 4;

const buttons = [
    "buttonOpen",
    "buttonSave",
    "buttonSaveAs",
    "buttonAddRectangle",
    "buttonAddEllipse",
];

export default class View {
    constructor() {
        this._activeTabID = tabInsert;
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

    doOnWindowResize(cb) {
        window.addEventListener("resize", cb);
    }

    onWindowResize() {
        let editorSpace = document.getElementById("editorSpace");
        const w = editorSpace.clientWidth;
        const h = w * 9 / 16;
        editorSpace.style.height = `${h}px`;
    }

    removeLoader() {
        document.getElementById("loader").remove()
    }

    doOnButtonClick(id, handler) {
        this._button(id).addEventListener(`click`, handler);
    }

    doOnTabClick(handler) {
        for (let id = 0; id < tabs.length; id++) {
            const idCopy = id;
            this._tab(id).addEventListener('click', () => {
                handler(idCopy);
            });
        }
    }

    setActiveTab(newActiveTabID) {
        if (newActiveTabID === this._activeTabID) {
            return;
        }

        this._tab(newActiveTabID).classList.replace("tab_inactive", "tab_active");
        this._menu(newActiveTabID).hidden = false;

        if (this._activeTabID !== null) {
            this._tab(this._activeTabID).classList.replace("tab_active", "tab_inactive");
            this._menu(this._activeTabID).hidden = true;
        }

        this._activeTabID = newActiveTabID;
    }

    newShapeView(type) {
        return new ShapeView(type, this._editorSpace);
    }

    _tab(id) {
        return document.getElementById(tabs[id]);
    }

    _button(id) {
        return document.getElementById(buttons[id]);
    }

    _menu(id) {
        return document.getElementById(menus[id]);
    }

    get _editorSpace() {
        return document.getElementById("editorSpace");
    }
}