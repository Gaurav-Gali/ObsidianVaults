![Source](https://youtu.be/gFWhbjzowrM?list=PLGjplNEQ1it_oTvuLRNqXfz_v_0pq6unW&t=5)

```js
// Creating an array
let arr = [10, 20, 30, 40];
let arr2 = ["hello" , "world"]

// Accessing the array elements with index
console.log(arr[0]);

// Getting the size of the array
console.log(arr.length);

// Iterating over the array
for (let i of arr) {
	console.log(i);
}

// Array methods
arr.push(50);
console.log(arr);

arr.pop();
console.log(arr);

console.log(arr.toString());

console.log(arr.concat(arr2));

arr.unshift(0); // adds elements to the start of array
console.log(arr);

arr.shift(); // removes elements from the start of array
console.log(arr);

console.log(arr.slice(2, 4));

// Adding elements using splice
arr.splice(1, 0, 44);
console.log(arr);

// Deleting an element using splice
arr.splice(1, 1);
console.log(arr);

// Replacing an element using splice
	arr.splice(2, 1, 45)
	console.log(arr);
```
