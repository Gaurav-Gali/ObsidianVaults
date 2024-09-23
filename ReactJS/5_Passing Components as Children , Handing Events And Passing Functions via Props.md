![Source](https://youtu.be/eILUmCJhl64?t=17450)

### Passing Components as Children
==Children is a special prop for passing elements into components==
Accessed with **props.children**

New elements can be passed as props in the parent side
```js
import Display from "./Display";

const Show = (props) => {
	return (
		<Display>
			<h1>Header</h1>
			<span>Content</span>	
		</Display>
	);
} 
```

Props which are passed in the parent component can be accessed in the child by **props.children**
```js
const Display = (props) => {
	return (
		<div>
			<strong>Title</strong>
			{props.children}
		</div>
	);
}
```

<hr>

### Handling Events
1. React events use camel case i.e onClick , onChange etc.
2. It uses **Synthetic** events and not direct browser events.
3. Event handlers can be functions or arrow functions.
4. Use **onChange** for controlled form events.
5. Avoid inline arrow functions in JSX for performance.
![[eventhandling.png]]

**onClick()**
```js
// 1. By defining inline arrow functions
<button
	onClick={() => console.log("delete")}
	className="btn btn-danger"
>
	Delete
</button>

// 2. By passing a external function to handle the event
<button
	onClick={() => handleDelete()}
	className="btn btn-danger"
>
	Delete
</button>

// 3. By passing the event object of the event into the external function.
<button
	onClick={(event) => handleDelete(event)}
	className="btn btn-danger"
>
	Delete
</button>

// 4. By passing the reference of the external function
<button
	onClick={handleDelete}
	className="btn btn-danger"
>
	Delete
</button>
```

**OnChange()**
```js
const handleInput = (e) => {
	let value = e.target.value;
	console.log(value);
}

<input
	type="text"
	className="form-control"
	placeholder="Enter Task"
	onChange={(e) => handleInput(e)}
/>
```

<hr>

### Passing Functions via Props
1. ==Enables upward communication from child to parent==
2. Here the parent passes an event handler for the child to execute.

```js
// 1. Parent component
<List
	listData={listElements}
	handleDelete={() => console.log("Deleted")}
/>

// 2. Child component
<button
	onClick={props.handleDelete}
	className="btn btn-danger"
>
	Delete
</button>
```

