export default class Signal {
    constructor() {
        this._slots = [];
    }

    connect(slot) {
        this._slots.push(slot);
    }

    emit(...args) {
        this._slots.forEach((slot) => {
            slot(...args);
        });
    }
}
