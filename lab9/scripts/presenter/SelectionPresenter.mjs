import * as RibbonView from "../view/RibbonView.mjs";

export default class SelectionPresenter {
    constructor(view) {
        this._view = view;

        this._selectedShapeID = null;
        this._view.ribbon.enableButton(RibbonView.buttonDelete, false)
    }

    onShapeClicked(id) {
        this.selectedShapeID = id;
    }

    get selectedShapeID() {
        return this._selectedShapeID;
    }

    set selectedShapeID(id) {
        if (this._selectedShapeID !== null) {
            this._view.document.getShape(this._selectedShapeID).selected = false;
        }
        if (id !== null) {
            this._view.document.getShape(id).selected = true;
        }

        this._selectedShapeID = id;
        this._view.ribbon.enableButton(RibbonView.buttonDelete, id !== null)
    }
}