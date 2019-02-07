class Connection {
    constructor(id, signal, slot) {
        this._id = id;
        this._signal = signal;
        this.slot = slot;
    }

    disconnect() {
        this._signal.disconnect(this._id);
    }
}

export default class Signal {
    constructor() {
        this._connections = new Map();
        this._lastConnectionID = 0;
    }

    connect(slot) {
        const id = this.newConnectionID;
        const conn = new Connection(id, this, slot);
        this._connections.set(id, conn);

        return conn;
    }

    disconnect(id) {
        this._connections.delete(id);
    }

    emit(...args) {
        for (let conn of this._connections.values()) {
            conn.slot(...args);
        }
    }

    get newConnectionID() {
        return ++this._lastConnectionID;
    }
}