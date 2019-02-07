import View from "./view/View.mjs";
import Presenter from "./presenter/Presenter.mjs";
import AppModel from "./app_model/AppModel.mjs";

export default class App {
    constructor() {
        new Presenter(new AppModel(), new View());
    }
}