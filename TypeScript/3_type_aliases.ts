// 1. Creating a type alias
type UserType = {
    name: string;
    age: number;
    phone?: string;
};

const user: UserType = {
    name: "John Doe",
    age: 30,
    phone: "123-456-7890"
};
console.log(user); // { name: 'John Doe', age: 30, phone: '123-456-7890' }

// 2. Functional type alias
type myfunc = (age: number) => string;

const func = (age) => {
    return String(age);
};
console.log(func(30)); // "30"

// 3. Interface types
interface IUser {
    name: string;
    age: number;
}

interface IClient extends IUser {
    phone?: string;
}

const client1: IClient = {
    name: "Jane Doe",
    age: 25,
    phone: "987-654-3210"
};
console.log(client1); // { name: 'Jane Doe', age: 25, phone: '987-654-3210' }
