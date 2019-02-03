import * as View from "./View.mjs";

export default class Presenter {
    constructor(model, view) {
        this.model = model;
        this.view = view;

        this.lastShapeID = 0;
    }

    init() {
        this.view.doOnShapeUpdate((id) => {
            console.log(`Shape ${id} was moved/resized`)
        });

        this.view.doOnButtonClick(View.buttonAddRectangle, () => {
            this.lastShapeID++;
            this.view.addRectangle(this.lastShapeID);
        });

        this.view.doOnButtonClick(View.buttonAddEllipse, () => {
                this.lastShapeID++;
                this.view.addEllipse(this.lastShapeID)
            }
        );

        this.view.removeLoader()
    }
}