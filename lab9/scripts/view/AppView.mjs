import DocumentView from "./DocumentView.mjs";
import RibbonView from "./RibbonView.mjs";

export default class AppView {
    constructor() {
        this.document = new DocumentView();
        this.ribbon = new RibbonView();
    }

    doOnWindowResize(cb) {
        window.addEventListener("resize", cb);
    }

    doOnKeyPressed(key, handler) {
        document.addEventListener("keydown", (event) => {
            if (event.key !== key) {
                return;
            }

            handler();
        })
    }

    removeLoader() {
        document.getElementById("loader").remove()
    }

    openFile() {
        let input = document.createElement('input');
        input.type = 'file';

        let reader = new FileReader();

        input.addEventListener('change', (event) => {
            const file = event.target.files[0];
            reader.readAsText(file, 'utf-8');
        });

        let promise = new Promise((resolve, reject) => {
            reader.onload = () => {
                resolve(reader.result);
            };
            reader.onerror = (event) => {
                reader.abort();
                reject(event.error);
            };
        });

        input.click();

        return promise;
    }

    saveFile(data, type, name) {
        let blob = new Blob([data], {type: type});

        let a = document.createElement('a');
        a.href = window.URL.createObjectURL(blob);
        a.download = name;

        a.click();
    }
}