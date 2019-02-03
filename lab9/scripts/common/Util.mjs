export default class Util {
    static extractMousePosition(event) {
        return {
            x: event.pageX,
            y: event.pageY,
        }
    }

    static throttle(cb, delay) {
        let timeout = null;
        return function (...args) {
            if (timeout !== null) {
                return;
            }

            cb(...args);
            timeout = setTimeout(() => {
                timeout = null;
            }, delay);
        };
    }
}