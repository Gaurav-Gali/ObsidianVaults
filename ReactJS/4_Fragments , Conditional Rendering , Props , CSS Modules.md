![Source](https://youtu.be/eILUmCJhl64?t=11016)

### Fragments
Usually when we would have to return multiple HTML tags from a function , then we would have to wrap it around another div. The problem with this is that i would add another layer to the DOM.

To prevent this we use a **React Fragment** so that that extra layer is not added.

```js
import React from 'react';

const App = () => {
	return (
		<React.Fragment>
			...
		</React.Fragment>
	);
}

export default App;
```

Shorter Syntax
```js
import React from 'react';

const App = () => {
	return (
		<>
			...
		</>
	);
}

export default App;
```

<hr>

### Map Method
This is an array method to iterate over the elements of the array.

```js
{/* Start Todos */}
	{listElements.map((item) => (
		<tr key={item.id}>
			<td>{item.id}</td>
			<td>{item.title}</td>
			<td>{item.date}</td>
			<td>
				<button className="btn btn-danger">
					Delete
				</button>
			</td>
		</tr>	
	))}
{/* End Todos */}
```

==**Note** : It is important to mention the key for the tags in the map method , this is useful for the fast rendering of the app.==

<hr>

### Conditional Rendering
```js
<tr
	clasName={item.id === 1 && "first"}
	key={item.id}
>
```

1. If-else statements
```js
if(condition === true) {
	...
} else {
	...
}
```
2. Ternary Operators
```js
{condition === true ? "for true" : "for false"}
```
4. Logical Operators
```js
{condition === true && "only for true"}
```

<hr>

### Passing Data Via Props
1. Data is passed down from the parent component to the child component.
2. The parent component passes the data to the child using attributes.
```js
import List from ".List";

<List listData={listElements} />
```
4. The child component receives the data as an object.
```js
function List({ listData }) {

	return (
	
		<>
			
			{listData.map((item) => (
			
				<tr className={item.id === 1 && "first"} key={item.id}>
				
					<td>{item.id}</td>
					<td>{item.title}</td>
					<td>{item.date}</td>
					
					<td>
						<button className="btn btn-danger">Delete</button>
					</td>
				
				</tr>
			
			))}
		
		</>
	
	);

}
```

<hr>

### CSS Modules
1. It used localised class names to avoid **global** conflicts.
2. Styles are scoped to individual components.
3. To create a css module , use this syntax : ==name.module.css==
```js
import TodoStyle from "../styles/ToDos.module.css";

className={item.id === 1 && TodoStyle.firstElement}
```

