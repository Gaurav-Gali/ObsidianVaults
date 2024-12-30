// 1. A variable of one type cannot be assigned to a variable of another type.
let num: number = 42;
let str: string = "Hello";
// num = str; // This will cause an error because num is declared as a number and str is declared as a string. They are not compatible types.

// 2. Assigning multiples types to a variable is possible.
let mixed: number | string = 42; // This is also called as union types.
mixed = "Hello";

// 3. Using the typeof operator to check the type of a variable.
console.log(typeof num); // Output: number
console.log(typeof str); // Output: string
console.log(typeof mixed); // Output: number or string

// 4. Type for arrays.
let arr: number[] = [1, 2, 3];
arr.push(4); // This is valid because the array is declared as a number type.

let arr1: (string | number)[] = ["Hello", 42]; // This is also valid because the array is declared as a union of string and number types.

// 5. Type for objects.
let obj: { name: string; age: number; phone?: string } = { name: "John", age: 30 };
obj.name = "Jane";

// 6. Any type can be used to represent any value.
let anyVar: any = 42;
anyVar = "Hello";
anyVar = true;

