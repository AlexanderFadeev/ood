import Model from "./Model.mjs";
import View from "./View.mjs";
import Presenter from "./Presenter.mjs";


export default class App {
    constructor() {
        this.model = new Model();
        this.view = new View();
        this.presenter = new Presenter(this.model, this.view);
    }

    start() {
        this.presenter.init();
    }
}