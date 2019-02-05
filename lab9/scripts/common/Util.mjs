import Vector from "./Vector.mjs";

export default class Util {
    static extractMousePosition(event, scale) {
        return new Vector(event.pageX / scale.x, event.pageY / scale.y);
    }

    static extractElementSize(element) {
        const rect = element.getBoundingClientRect();
        return new Vector(rect.width, rect.height);
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