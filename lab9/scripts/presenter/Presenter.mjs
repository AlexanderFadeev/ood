import * as View from "../view/View.mjs";
import Shape from "../model/Shape.mjs";
import ShapePresenter from "./ShapePresenter.mjs";
import Rect from "../common/Rect.mjs";

const defaultShapeRect = new Rect(1 / 4, 1 / 4, 1 / 2, 1 / 2);

export default class Presenter {
    constructor(model, view) {
        this._model = model;
        this._view = view;

        this._shapePresenters = new Map();

        this._view.doOnWindowResize(this._onWindowResize.bind(this));
        this._view.doOnTabClick(this._view.setActiveTab.bind(this._view));

        this._view.enableButton(View.buttonOpen, false); //TODO
        this._view.doOnButtonClick(View.buttonSave, this._save.bind(this));
        this._view.enableButton(View.buttonSaveAs, false); //TODO

        this._view.doOnButtonClick(View.buttonUndo, this._model.undo.bind(this._model));
        this._view.doOnButtonClick(View.buttonRedo, this._model.redo.bind(this._model));

        this._view.doOnButtonClick(View.buttonAddRectangle, () => {
            this._model.addShape(Shape.Type.rectangle, defaultShapeRect.clone());
        });
        this._view.doOnButtonClick(View.buttonAddEllipse, () => {
            this._model.addShape(Shape.Type.ellipse, defaultShapeRect.clone());
        });

        this._model.onShapeAdded.connect(this._onShapeAdded.bind(this));
        this._model.onShapeRemoved.connect(this._onShapeRemoved.bind(this));
        this._model.onHistoryUpdate.connect(this._updateHistoryButtons.bind(this));

        this._updateHistoryButtons();
        this._view.removeLoader();
    }

    _onWindowResize() {
        this._view.onWindowResize();

        for (let shapePresenter of this._shapePresenters.values()) {
            shapePresenter.onWindowResize();
        }
    }

    _onShapeAdded(id) {
        const shape = this._model.getShape(id);
        let shapePresenter = new ShapePresenter(shape, this._view);
        this._shapePresenters.set(id, shapePresenter);
    }

    _onShapeRemoved(id) {
        this._shapePresenters.get(id).remove();
        this._shapePresenters.delete(id);
    }

    _updateHistoryButtons() {
        this._view.enableButton(View.buttonUndo, this._model.canUndo());
        this._view.enableButton(View.buttonRedo, this._model.canRedo());
    }

    _save() {
        this._view.saveFile(this._model.serialize(), 'application/json', "shapes_data.json")
    }
}