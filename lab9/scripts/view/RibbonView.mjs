let i = 0;
export const tabFile = i++;
export const tabEdit = i++;

const tabs = [
    "tabFile",
    "tabEdit",
];

i = 0;
export const menuFile = i++;
export const menuEdit = i++;

const menus = [
    "menuFile",
    "menuEdit",
];

i = 0;
export const buttonNew = i++;
export const buttonOpen = i++;
export const buttonSave = i++;
export const buttonSaveAs = i++;
export const buttonUndo = i++;
export const buttonRedo = i++;
export const buttonDelete = i++;
export const buttonAddRectangle = i++;
export const buttonAddEllipse = i++;

const buttons = [
    "buttonNew",
    "buttonOpen",
    "buttonSave",
    "buttonSaveAs",
    "buttonUndo",
    "buttonRedo",
    "buttonDelete",
    "buttonAddRectangle",
    "buttonAddEllipse",
];

export default class RibbonView {
    constructor() {
        this._activeTabID = tabEdit;
    }

    doOnButtonClick(id, handler) {
        this._button(id).addEventListener(`click`, () => {
            if (this._button(id).classList.contains("button_inactive")) {
                return;
            }

            handler();
        });
    }

    enableButton(id, enabled) {
        const active = "button_active";
        const inactive = "button_inactive";

        this._button(id).classList.replace(
            enabled ? inactive : active,
            enabled ? active : inactive
        )
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

    _tab(id) {
        return document.getElementById(tabs[id]);
    }

    _button(id) {
        return document.getElementById(buttons[id]);
    }

    _menu(id) {
        return document.getElementById(menus[id]);
    }
}