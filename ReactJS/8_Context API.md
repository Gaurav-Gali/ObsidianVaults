[Source](https://youtu.be/eILUmCJhl64?t=31115)

1. Prop drilling : Context API solves the issue of prop drilling.
2. Folder setup : Use a **store** folder for context files.
3. Initialisation : Use React.createContext with initial state and export it.
4. Provider : Implement with contextName.Provider in components.
5. Accessing values : Use the useContext hook.
6. Dynamic data : combine the context with state.
7. Export functions : Contexts can also export functions for actions.
8. Logic Separation : This helps the UI and business logic independent from each other.

![[contextapi.png]]
### Creating context
1. Create a context-store
```js
import {createContext} from "react";

export const TodoListContext = createContext();
```

2. Establish a context provider for all the required components to access the store
```js
import TodoListContext from "./store/todo-list-store";
import {useState} from "react";
import Todos from "./components/Todos"

function App() {
	const [list , setList] = useState([]);
	const addItem = () => {
		...
	};
	const deleteItem = () => {
		...
	};
	return (
		<TodoListContext.Provider value={
			{
				list : list,
				deleteItem : deleteItem,
				addItem : addItem 
			}
		}>
			<Todos />
		<TodoListContext.Provider/>
	);
};
```

3. Use the context in the required files using the useContext hook
```js
import TodoListContext from "./store/todo-list-store";
import {useContext} from "react";

function Todos() {
	const contextObj = useContext(TodoListContext);
	return (
		{contextObj.list.map(item => (
			<li> item </li>
		))}
	);
};
```
