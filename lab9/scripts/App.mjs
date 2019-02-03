import Model from "./model/Model.mjs";
import View from "./view/View.mjs";
import Presenter from "./presenter/Presenter.mjs";


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