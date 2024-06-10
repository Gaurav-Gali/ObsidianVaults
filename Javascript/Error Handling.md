![Source](https://youtu.be/N-O4w6PynGY?list=PLGjplNEQ1it_oTvuLRNqXfz_v_0pq6unW&t=3362)

Error handling is performed using the try catch block in JS

```js
let a = 10;
let b = 20;

for (let i = a; i >= -1; i--) {
	try {
		let c = b / i;
		console.log(c);
	} catch (err) {
		throw "Can't divide by zero";
	}
}
```

The ==throw== keyword is used to send custom error messages.
