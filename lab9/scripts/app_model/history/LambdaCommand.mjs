export default class LambdaCommand {
    constructor(execute, unexecute) {
        this._execute = execute;
        this._unexecute = unexecute;
    }

    execute() {
        this._execute();
    }

    unexecute() {
        this._unexecute();
    }
}