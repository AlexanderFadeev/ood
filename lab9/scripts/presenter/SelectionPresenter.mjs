import * as View from "../view/View.mjs";

export default class SelectionPresenter {
    constructor(view) {
        this._view = view;

        this._selectedShapeID = null;
        this._view.enableButton(View.buttonDelete, false)
    }

    onShapeClicked(id) {
        this.selectedShapeID = id;
    }

    get selectedShapeID() {
        return this._selectedShapeID;
    }

    set selectedShapeID(id) {
        if (this._selectedShapeID !== null) {
            this._view.getShape(this._selectedShapeID).selected = false;
        }
        if (id !== null) {
            this._view.getShape(id).selected = true;
        }

        this._selectedShapeID = id;
        this._view.enableButton(View.buttonDelete, id !== null)
    }
}