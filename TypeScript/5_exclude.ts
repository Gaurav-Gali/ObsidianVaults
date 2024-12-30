type shapes = "circle" | "square" | "cube";
type TwoDShapes = Exclude<shapes, "cube">;

const shape1: shapes = "cube";
const shape2: TwoDShapes = "square"; // Error: 'cube' is not assignable to type 'TwoDShapes'
