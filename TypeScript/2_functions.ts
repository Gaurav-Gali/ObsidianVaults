// 1. Setting types for function parameters and return values.
const func1 = (param: number): string => {
    return `The value is ${param}`;
};

// 2. Void function that does not return anything.
const func2 = (): void => {
    console.log("This function does not return anything.");
};

// 3. Function with default parameters.
const func3 = (param1: number, param2: string = "default"): number => {
    return param1 + param2.length;
};

// 4. Optional parameter in a function.
const func4 = (param1: number, param2?: string): void => {
    if (param2) {
        console.log(`The value is ${param1} and the optional value is ${param2}`);
    } else {
        console.log(`The value is ${param1}`);
    }
};

