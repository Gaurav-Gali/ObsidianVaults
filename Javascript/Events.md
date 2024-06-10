![Source](https://youtu.be/_i-uLJAh79U?list=PLGjplNEQ1it_oTvuLRNqXfz_v_0pq6unW)

**Note** :
1. Changes in the state of an object is known as an event

### Using Events
```js
let elem = document.getElementsByClassName("changeColor")[0];

// Using normal events
elem.onclick = (e) => {
	console.log(e.clientX, e.clientY);
};
```

### Using Event Listeners
```js
let elem = document.getElementsByClassName("changeColor")[0];

// Using Event Listeners
elem.addEventListener("click", (e) => {
	console.log(e.type);
});

const handler2 = (e) => {
	console.log(e.type);
};

elem.addEventListener("click", handler2);

// Removing event listeners
elem.removeEventListener("click" , handler2);
```
