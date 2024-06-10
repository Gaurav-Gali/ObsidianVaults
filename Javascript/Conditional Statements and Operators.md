![Source](https://youtu.be/Zg4-uSjxosE?list=PLGjplNEQ1it_oTvuLRNqXfz_v_0pq6unW)

### Operators
1. Arithmetic Operators
	1. + , - , \* , / , % \** , these are binary operators.
	2. ++ , -- , these are unary operators.
2. Assignment Operators
	1. += , -= , \*= , /= , %= , **=
3. Comparison Operators
	1. == , === , != , !== , < , > , <= , >=
	2. Here == and != will only check the value and not the data type.
	3. In === and !== , it will check both value and data type.
4. Logical Operators
	1. && , || , !

### Conditional Statements
```javascript
// Basic conditional statement
if (1 == "1") {
	console.log("Condition is true");
} else {
	console.log("Condition is false");
}

// If-else-if ladder
if (1 === "2") {
	console.log("Same Type");
} else if (1 == "2") {
	console.log("Different type");
} else {
	console.log("Different value");
}

// Ternary operator
1 == "1" ? console.log("Same Value") : console.log("Different Value");

// Switch case
let fruit = "Apple";

switch (fruit) {
	case "Apple":
	case "Orange":
	case "Pineapple":
		console.log("Available");
		break;
	default:
		console.log("Not Available");
}
```
