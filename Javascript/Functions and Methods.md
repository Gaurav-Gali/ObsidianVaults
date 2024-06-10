![Source](https://youtu.be/P0XMXqDGttU?list=PLGjplNEQ1it_oTvuLRNqXfz_v_0pq6unW&t=58)

**Note**:
1. Callback Functions :
	1. They are the functions which are passed as arguments to other function.
2. Arrow functions :
	1. They are the same as callback function but the function name is not mentioned.
	2. It takes 3 arguments (value , index , array)
3. The **forEach** method/function is also called as higher ==order functions/methods==
4. Higher order functions are the functions that take a function as an argument or returns a function.
5. The **map** method acts same as the **forEach** method but it returns the values and stores it in another array after some operation. It also takes 3 arguments (value , index , array)
6. The **filter** method is also like the **map** method but it returns a condition. Which ever element of the array satisfies the condition , it will be pushed to a new array.
7. The **reduce** method is used to return a single value upon performing an operation on an array. It takes 2 arguments (accumulator , current_value).

```js
// Declaring a function
function addition(x, y) {
	return x + y;
}
console.log(addition(10, 20));

// Declaring a call back function
let arr = [1, 2, 3, 4, 5];
arr.forEach(function square(x) {
	console.log(x ** 2);
});

// Using arrow functions
arr.forEach((val) => {
	console.log(val ** 3);
});

// Map method
let arr2 = arr.map((val) => {
	return val ** 2;
});

console.log(arr2);

// Filter method
let arr3 = arr.filter((val) => {
	return val % 2 === 0;
});

console.log(arr3);

// Reduce method
const sum = arr.reduce((acc, cur) => {
	return acc + cur;
});

console.log(sum);
```
