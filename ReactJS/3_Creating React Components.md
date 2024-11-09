[Source](https://youtu.be/eILUmCJhl64?t=4089)

### File Extensions
1. .js : It is the standard javascript extension and the file contains plain/vanilla javascript code.
2. .jsx : It stands for javascript XML , it combines javascript with HTML like tags making it easy for designing UI.

<hr>
### What is JSX
###### Definition:
It determines how a UI will look whenever a component is used.

###### Not HTML
Though it looks like HTML , it is actually JavaScript XML

###### Conversion
The JSX gets converted to regular JavaScript

###### Babeljs.i/repl
This is a tool that allows you to see how JSX is transformed into regular JavaScript

<hr>
### Creating Components
==**Note** : After creating the component , don't forget to export it.==

**Note :** *The naming convention of components is Capitalised EX : SideBar , NavBar etc...*
*This is to differentiate normal HTML tags from the created components.*

```jsx
// This is an example of a functional component
function Navbar() {
	return (
		<>
			<h1>Hello world</h1>
		</>
	);
}

export default Navbar;
```
Here the App function is exported in order to access the App function elsewhere.
==This type of export is called as default export.==

<table border="1px">
<tr>
<th>Functional Components</th>
<th>Class Based Components</th>
</tr>
<tr>
<td>They are initially stateless</td>
<td>They are stateful , i.e they can manage states</td>
</tr>
<tr>
<td>Can use hooks for managing states and effects</td>
<td>They use lifecycles to access lifecycle methods</td>
</tr>
<tr>
<td>Simpler and more concise</td>
<td>They are verbose , i.e they require more boilerplate code</td>
</tr>
<tr>
<td>They are more popular</td>
<td>They are not preferred anymore</td>
</tr>
</table>

###### Exporting multiple components in the same file
```jsx
export function Sidebar() {
	return ();
}

export function Button() {
	return ();
}
```
==This type of export is called as multi named exports.==

###### Exporting multiple components in the same file with one default export.
```jsx
fucntion Sidebar() {
	return ();
}

export function Button() {
	return ();
}

export default Sidebar();
```

<hr>

**Note :** *CSS files can be directly imported into jsx files either as modules or as stylesheet*
###### To import the exported components
```jsx
import NavBar from "./Navbar";
import "index.css";

function App() => {
	return (
		<>
			<NavBar></NavBar>
		</>
	);
}

export default App;
```

###### Importing multiple exported components from the same file
```jsx
import {Sidebar , Button} from "./Comps";

fucntion App() {
	return (
		<>
			<Sidebar/>
			<Button/>
		</>
	);
}

export default App;
```

<hr>

### Creating Dynamic Components
###### Dynamic Content
This allows JSX to create dynamic and interactive UI components.

###### JavaScript Expressions
In order to use JavaScript expressions directly into JSX , use the curly brackets
```jsx
function Greet() {
	let name = "User1";
	return (
		<h1 style={{'background-color' : '#454545'}}>
			Good Morning {name} , Have a nice day !
		</h1>
	);
}

export default Greet();
```
