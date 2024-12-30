// 1. Generic Functions
function showArray<T>(arr: T[]) {
    console.log(arr);
}

showArray<number>([1, 2, 3]);
showArray<string>(["a", "b", "c"]);

// 2. Generic Types
type ApiResponse<T> = {
    status: string;
    data: T;
};

const response: ApiResponse<{ name: string; age: number }> = {
    status: "success",
    data: { name: "John Doe", age: 30 },
};

// Alternatively
type UserResponse = ApiResponse<{ name: string; age: number }>;
const userResponse: UserResponse = {
    status: "success",
    data: { name: "Jane Smith", age: 25 },
};

// 4. Passing default values for generic types
interface IUser<T = { name: string; age: number }> {
    name: string;
    age: number;
}

