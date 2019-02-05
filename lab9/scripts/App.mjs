import Document from "./model/Document.mjs";
import View from "./view/View.mjs";
import Presenter from "./presenter/Presenter.mjs";


export default class App {
    constructor() {
        this.model = new Document();
        this.view = new View();
        this.presenter = new Presenter(this.model, this.view);
    }
}