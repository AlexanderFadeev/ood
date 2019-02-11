import Presenter from "./presenter/Presenter.mjs";
import AppModel from "./app_model/AppModel.mjs";
import AppView from "./view/AppView.mjs"

export default class App {
    constructor() {
        new Presenter(new AppModel(), new AppView());
    }
}