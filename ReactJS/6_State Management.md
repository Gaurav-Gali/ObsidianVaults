[Source](https://www.youtube.com/watch?v=eILUmCJhl64&t=20325s)

1. State represents data that changes over time.
2. State is local and private to that component.
3. State changes cause the component to re-render.
4. For functional components use the useState() hook.
5. React functions that start with "use" are called ==hooks==.
6. Hooks should only be used inside components.
7. Parent components can pass the hooks down to children via props.
8. Lifting states up : share states between components by moving it to their closest common ancestor.
![[states.png]]

**Note**
1. ==The Spread operator is used to apply all the value of an array in another array==
2. This can be used to add more elements to the array.
```js
let lis = [1,2,3];
let newLis = [...lis , 4,5,6];
```

3. While setting states , if the value to be set depends on the previous value , then this method won't be recommended , since react works asynchronously , it might lead to errors.
```js
const [lis , setLis] = useState([1,2,3]);

const updateLis = () => {
	let newLis = [..lis , 4,5,6];
	setLis(newLis);
};
```

4. Instead the set function will provide you with the previous value , which can be accessed using a call back function.
```js
const [lis , setLis] = useState([1,2,3]);

const updateLis = () => {
	setLis((curVal) => {
		let newLis = [...curVal , 4,5,6];
		return newLis;
	});
	// or

	setLis((curVal) => [...curVal , 4,5,6]);
};
```

### Use Ref Hook

**The problem with useState** :
When a state is being updated , even though the value is not being used in the ui , the component will re-render.
To avoid this we use useRef , which does not re-render when it's state is changed.

```js

```