const Type = {
    rectangle: "rectangle",
    ellipse: "ellipse"
};

export default class Shape {
    constructor(type, rect) {
        this.type = type;
        this.rect = rect;
    }

    static get Type() {
        return Type;
    }
}