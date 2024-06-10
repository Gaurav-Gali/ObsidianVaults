![Source](https://youtu.be/N-O4w6PynGY?list=PLGjplNEQ1it_oTvuLRNqXfz_v_0pq6unW)

### Prototypes in JS
Every object in JS by default inherits some properties from it's respective prototype objects which is stores within the object itself.

==The type of the prototype is actually either a reference to it's object or null.==

If a prototype object and the object have the same methods then the object's method will override the prototype's method.

We can set prototype using \_proto__
###### 2 ways of declaring methods in an object
```js
// Declaring methods in an object
const Employee = {
	// Methods can be declared without providing the key as well
	calculateTax() {
		console.log("10% tax deducted");
	},
	
	// Methods can be declared by providing the key
	calculateTax2: function () {
		console.log("20% tax deducted");
	},
};
```

If we want to set the above mentioned **Employee** object as a prototype for another object we can do that by the following :
```js
const Emp1 = {
	salary: 10000,
};

// Setting Employee as the prototype for Emp1
Emp1.__proto__ = Employee;

Emp1.calculateTax();
```

<hr>
### Classes In JS
A class is a program code template for creating new objects, these objects would have some **state** variables and methods inside it. 

```js
// Creating classes
class Car {
	start() {
		console.log("Start car");
	}

	setModel(model) {
		this.model = model;
	}
}

// Creating objects of a class
let toyota = new Car();
toyota.setModel("Camry");

console.log(toyota.model);
```

###### Constructors
It is a special methods of a class that is automatically invoked when an object of the class is created.
use the ==constructor()== methods to declare constructors.

==If a constructor is not declared , then it is automatically created and invoked by the **new** keyword==

```js
// Creating classes with constructor
class Bike {
	constructor(model, milage) {
		this.model = model;
		this.milage = milage;
	}
	
	getData() {
		console.log(this.model , this.milage);
	}
}

let RE = new Bike("Meteor", 35);
RE.getData();
```

<hr>

### Inheritance
It the process of passing down properties and methods from the parent class to it's child class.

==To perform inheritance the **extends** keyword is used==

If the parent and the child class have a same method then the method from the child class will be executed , this is called **Method Overriding**
```js
// Performing Inheritance
class Human {
	sleep() {
		console.log("Sleeping...");
	}
}

class Engineer extends Human {
	work() {
		console.log("Solves problems...");
	}
}

let Eng1 = new Engineer();
Eng1.sleep();
Eng1.work();
```


###### Super keyword
This keyword is used to access the constructor of the parent class.

While creating constructors in the child class , it is important to call the constructor of the parent class using the super() keyword.
```js
// Implementing the super keyword
class Parent {
	constructor(name) {
		this.name = name;
	}
	eat() {
		console.log("Eating...");
	}
}

class Child extends Parent {
	constructor(name, age) {
		super(name);
		this.age = age;
	}

	work() {
		super.eat();
		console.log("Starting to work...");
	}
}

let child1 = new Child("ChildName", 20);
console.log(child1.name, child1.age);
child1.work();
```
