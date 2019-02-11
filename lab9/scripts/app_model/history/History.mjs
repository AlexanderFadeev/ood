import Signal from "../../common/Signal.mjs";

export default class History {
    constructor(size) {
        this._size = size;
        this.onUpdate = new Signal();
        this.reset();
    }

    reset() {
        this._commands = [];
        this._lastExecutedCommandIndex = -1;
        this.onUpdate.emit();
    }

    addAndExecute(cmd) {
        this.add(cmd);
        this.redo();
    }

    add(cmd) {
        this._removeUndoneCommands();
        this._commands.push(cmd);
        if (this._isOverflown) {
            this._removeOldestCommand();
        }
        this.onUpdate.emit();
    }

    canUndo() {
        return this._lastExecutedCommand !== null;
    }

    canRedo() {
        return this._lastUndoneCommand !== null;
    }

    undo() {
        this._lastExecutedCommand.unexecute();
        this._lastExecutedCommandIndex--;
        this.onUpdate.emit();
    }

    redo() {
        this._lastUndoneCommand.execute();
        this._lastExecutedCommandIndex++;
        this.onUpdate.emit();
    }

    get _lastExecutedCommand() {
        if (this._lastExecutedCommandIndex < 0) {
            return null;
        }

        return this._commands[this._lastExecutedCommandIndex];
    }

    get _lastUndoneCommand() {
        if (this._lastExecutedCommandIndex >= this._commands.length - 1) {
            return null;
        }

        return this._commands[this._lastExecutedCommandIndex + 1];
    }

    _removeUndoneCommands() {
        this._commands = this._commands.slice(0, this._lastExecutedCommandIndex + 1);
    }

    _removeOldestCommand() {
        this._commands.shift();
        this._lastExecutedCommandIndex--;
    }

    get _isOverflown() {
        return this._commands.length > this._size;
    }
}