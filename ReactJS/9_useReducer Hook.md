![Source](https://youtu.be/eILUmCJhl64?t=33315)

**When to use useReducer hook** ?
	==Use this hook when multiple states have to be updated in a single function.==

1. useReducer hook offers more control over states than a useState hook.
2. Components : it involves 2 main components
	1. Reducer : A pure function that takes the current state and the action and returns the next state.
	3. Action : An object describing what happened , typically having a type property.
3. Initialisation : It is invoked as 
```js
const [state , dispatch] = useReducer(reduce , initialState);
```
4. Dispatch : Actions are dispatched using the dispatch function , which invokes the reducer with the current state and the given function.
5. Use cases : Particularly useful for managing states in large components or when the next state depends on the previous state.
6. Predictable state management : Due to it's strict structure , it leads to a more maintainable and predictable state management.

![[usereducer.png]]

```js
// Reducer function
const countReducer = (state, action) => {
	switch (action.type) {
		case ACTIONS.INC:
			return { value: state.value + action.payload.value };
		
		case ACTIONS.DEC:
			return { value: state.value - action.payload.value };
		
		default:
			return state.value;
	}
};

// 1.Implementing useReducer Hook
const [count, countDispatch] = useReducer(countReducer, { value: 0 });

// 2. Using the dispatch function
<button
onClick={() =>
	countDispatch({
		type: ACTIONS.DEC,
		payload: { value: operValue },
	})
}
>
```


