import Shape from "../model/Shape.mjs";
import ShapePresenter from "./ShapePresenter.mjs";
import Rect from "../common/Rect.mjs";
import SelectionPresenter from "./SelectionPresenter.mjs";
import * as RibbonView from "../view/RibbonView.mjs";

const defaultShapeRect = new Rect(1 / 4, 1 / 4, 1 / 2, 1 / 2);

export default class Presenter {
    constructor(model, view) {
        this._model = model;
        this._view = view;

        this._shapePresenters = new Map();
        this._selectionPresenter = new SelectionPresenter(this._view);

        this._initModel();
        this._initView();

        this._hack();
        this._view.removeLoader();
    }

    _initView() {
        this._view.doOnWindowResize(this._onWindowResize.bind(this));
        this._view.ribbon.doOnTabClick((id) => {
            this._view.ribbon.setActiveTab(id);
        });

        this._view.ribbon.doOnButtonClick(RibbonView.buttonOpen, this._open.bind(this));
        this._view.ribbon.doOnButtonClick(RibbonView.buttonSave, this._save.bind(this));
        this._view.ribbon.enableButton(RibbonView.buttonSaveAs, false); //TODO

        this._view.ribbon.doOnButtonClick(RibbonView.buttonUndo, this._model.undo.bind(this._model));
        this._view.ribbon.doOnButtonClick(RibbonView.buttonRedo, this._model.redo.bind(this._model));

        this._view.ribbon.doOnButtonClick(RibbonView.buttonAddRectangle, () => {
            this._model.addShape(Shape.Type.rectangle, defaultShapeRect.clone());
        });
        this._view.ribbon.doOnButtonClick(RibbonView.buttonAddEllipse, () => {
            this._model.addShape(Shape.Type.ellipse, defaultShapeRect.clone());
        });
        this._view.document.doOnShapeMouseDown((id) => {
            this._selectionPresenter.onShapeClicked(id);
        });
        this._view.doOnKeyPressed("Delete", this._onDelete.bind(this));
        this._view.ribbon.doOnButtonClick(RibbonView.buttonDelete, this._onDelete.bind(this));

        this._updateHistoryButtons();
    }

    _initModel() {
        this._model.onShapeAdded.connect(this._onShapeAdded.bind(this));
        this._model.onShapeRemoved.connect(this._onShapeRemoved.bind(this));
        this._model.onHistoryUpdate.connect(this._updateHistoryButtons.bind(this));
    }

    _onWindowResize() {
        this._view.document.onWindowResize();

        for (let shapePresenter of this._shapePresenters.values()) {
            shapePresenter.onWindowResize();
        }
    }

    _onShapeAdded(id) {
        let shapePresenter = new ShapePresenter(id, this._model, this._view.document);
        this._shapePresenters.set(id, shapePresenter);
    }

    _onShapeRemoved(id) {
        if (this._selectionPresenter.selectedShapeID === id) {
            this._selectionPresenter.selectedShapeID = null;
        }

        this._shapePresenters.get(id).remove();
        this._shapePresenters.delete(id);
    }

    _updateHistoryButtons() {
        this._view.ribbon.enableButton(RibbonView.buttonUndo, this._model.canUndo());
        this._view.ribbon.enableButton(RibbonView.buttonRedo, this._model.canRedo());
    }

    _open() {
        this._view.openFile().then((data) => {
            this._model.loadFile(data);
        })
    }

    _save() {
        this._view.saveFile(this._model.serialize(), 'application/json', "shapes_data.json")
    }

    _onDelete() {
        if (this._selectionPresenter.selectedShapeID === null) {
            return;
        }

        this._model.removeShape(this._selectionPresenter.selectedShapeID);
        this._selectionPresenter.selectedShapeID = null;
    }

    // First created shape cannot be normally resized
    // so we are creating and removing it
    // TODO: research
    _hack() {
        this._model.addShape("hack_shape", defaultShapeRect.clone());
        this._model.undo();
        this._model.resetHistory();
    }
}