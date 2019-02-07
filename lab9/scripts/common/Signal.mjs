export default class Signal {
    constructor() {
        this._slots = [];
    }

    connect(slot) {
        this._slots.push(slot);
    }

    emit(...args) {
        for (let slot of this._slots.values()) {
            slot(...args);
        }
    }
}