![Source](https://youtu.be/d3jXofmQm44?list=PLGjplNEQ1it_oTvuLRNqXfz_v_0pq6unW)

==**Note :** async await >> promise chains >> callback hell==

<table border="1px">
<tbody>
<tr>
<td>&nbsp;Synchronous</td>
<td>Asynchronous&nbsp;</td>
</tr>
<tr>
<td>
<p>&nbsp;Each instruction waits for the previous&nbsp; instruction to execute.</p>
</td>
<td>
<p>Asynchronous code doesn't wait wait for the previous instruction to get executed , rather gets executed immedeately</p>
</td>
</tr>
</tbody>
</table>

The **setTimeout()** function is used to run the required block of code after the given time limit which is mentioned in milliseconds.

```js
// The setTimeout() function
setTimeout(() => {
	console.log("After 2 sec");
}, 2000);
```

<hr>

### Callbacks
It is a function which is passes as an argument to another function.
```js
// Callback function
function calc(a, b, fn) {
	console.log(fn(a, b));
}

calc(1, 2, (a, b) => {
	return a + b;
});
```

###### Callback hell
It is when callback are nested below one and other forming a pyramid structure.
It is also called **Pyramid of Doom**.
```js
// Callback hell
function getData(ID, getNewData) {
	setTimeout(() => {
		console.log("Data", ID);
		if (getNewData) {
			getNewData();
		}
	}, 2000);

}

getData(1, () => {
	getData(2, () => {
		getData(3);
	});
});

```

<hr>

### Promises
In JS a promise is an object that represents an eventual completion or failure of a task.
It is a solution to callback hell.

A promise object has 2 handles or call backs , resolve and reject , these are provided by JS.

A promise has 3 states :
1. Pending
2. Fulfilled / Resolved
3. Rejected

```js
// Promises
function getData(ID, getNewData) {
	return new Promise((resolve, reject) => {
	
	setTimeout(() => {
		console.log("Fetched", ID);
		getNewData() && getNewData;
		resolve("Success");
	}, 2000);
	});
}

let promise = getData(1, () => {
	getData(2);
});
```
###### Using / Handling Promises
1. I the promise is resolved then we can use the ==**.then()**== method to handle it.
	1. We can pass a function into this method and it will only be executed when the promise is resolved.
	2. The call back takes an argument called res or response that will return the message passes through the resolve method.
2. I the promise is rejected then we can use the ==**.catch()**== method to handle it.
	1. We can pass a function into this method and it will only be executed when the promise is rejected.
	2. The call back takes an argument called err or error that will return the message passes through the reject method.

```js
promise.then((res) => {
	console.log(res);
});

promise.catch((err) => {
	console.log("Server error");
});
```

###### Promise Chaining
It is the process of handling multiple promises one after another.
```js
// Promise Chains

let p1 = getData(1);

p1.then((res) => {
	console.log(res);
	return getData(2);
}).then((res) => {
	console.log(res);
});
```

<hr>

### Async and Await
This concept is used to make asynchronous programming simple.

Here an async function always returns a promise and await pauses the execution of it's surrounding async functions until the promise is resolved.
```js
// Async Await
function weatherAPI() {
	return new Promise((resolve, reject) => {
		setTimeout(() => {
			console.log("Fetched weather");
			resolve(200);
		
		}, 2000);
		
	});
	
}

async function getWeather() {
	await weatherAPI(); //1st call
	await weatherAPI(); //2nd call
}

getWeather();
```

#### IIFE
It stands for Immediately Invoked Function Expression

It is used when you don't want to separately call the function , instead it will be called immediately after the function is declared.
```js
// Implementaion with normal function declaration
(
	fucntion foo() {
		...
	}
)();

// Implementaion with arrow function
(
	() => {
		...
	}
)();

// Implementaion with async arrow function
(
	async () => {
		...
	}
)();
```
