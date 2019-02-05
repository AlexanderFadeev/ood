import * as View from "../view/View.mjs";
import Shape from "../model/Shape.mjs";
import Rect from "../common/Rect.mjs";
import ShapePresenter from "./ShapePresenter.mjs";

export default class Presenter {
    constructor(model, view) {
        this._model = model;
        this._view = view;

        this._shapePresenters = [];

        this._lastShapeID = 0;

        this._view.doOnWindowResize(this._onWindowResize.bind(this));
        this._view.doOnTabClick(this._view.setActiveTab.bind(this._view));

        this._view.doOnButtonClick(View.buttonAddRectangle, () => {
            this._addShape(Shape.Type.rectangle);
        });

        this._view.doOnButtonClick(View.buttonAddEllipse, () => {
            this._addShape(Shape.Type.ellipse);
        });

        this._view.removeLoader()
    }

    _onWindowResize() {
        this._view.onWindowResize();

        for (let index = 0; index < this._shapePresenters.length; index++) {
            this._shapePresenters[index].onWindowResize()
        }
    }

    _addShape(type) {
        this._lastShapeID++;
        let shapeView = this._view.newShapeView(type, this._lastShapeID);
        let shapeModel = new Shape(type, new Rect(1 / 4, 1 / 4, 1 / 2, 1 / 2));
        let shapePresenter = new ShapePresenter(shapeModel, this._view, shapeView);

        this._shapePresenters.push(shapePresenter);
    }
}