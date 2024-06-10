![Source](https://youtu.be/UmRtFFSDSFo?list=PLGjplNEQ1it_oTvuLRNqXfz_v_0pq6unW)

### Loops

**Note** :
1. The brake statement is used for terminating the loop when required.
2. The continue statement is used for skipping the current iteration.

3. For loop
```js
for (let i = 0; i < 10; i++) {
	if (i === 5) {
		continue;
	}
	console.log(i);
}
```
2. While loop
```js
let num = 10;
while (num !== 0) {
	console.log(num);
	num--;
}
```
3. Do-while loop
```js
num = 10;

do {
	console.log(num);
	num--;
} while (num !== 0);
```
4. for-of loop , it is used for iterators like strings and arrays
```js
let s = "Hello";
for (let i of s) {
	console.log(i);
}
```
5. for-in loop , it is used for objects
```js
let student = {
	Name: "Student",
	Age : 19,
};

for (let key in student) {
	console.log(`value is ${student[key]}`);
}
```

### Strings
**Note** :
1. string interpolation is the concept where placeholder are used within the template literals to form a single string.
2. The string methods don't affect the original string , instead they return the required output.

```js
// Creating a string
let str = "Hello world";

// Creating template literals
let Name = "Person";
let t = `Hello ${Name}`;
console.log(t);

// length is a internal string variable that returns the size of the string
let l = str.length;
console.log(`length of ${str} is ${l}`);

// Accessing the string's indices
console.log(str[0]);

// String Methods
console.log(str.toUpperCase());
console.log(str.toLowerCase());
console.log(str.trim());

console.log(str.slice(0, 3));
console.log(str.concat(t));
console.log(str.replace("Hello", "Hey"));
console.log(str.charAt(0));
```
