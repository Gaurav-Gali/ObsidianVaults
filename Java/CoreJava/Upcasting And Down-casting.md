![Source](https://youtu.be/Q8cTydJSawQ?list=PLsyeobzWxl7pe_IiTfNyr55kwJPWbgxB5)

**Note** : Both *down-casting* and *upcasting* deals with ![[Typecasting]]

Upcasting
```java
class A {
	public void show1() {
		System.out.println("In class A");
	}
}

class B extends A {
	public void show2() {
		System.out.println("In class B");
	}
}

public class lesson {
	public static void main(String[] args) {
		A obj1 = new B(); // polymorphism
		// here the type of the is explicitly changed from child to parent
		obj2 = (A) obj1;
	}
}
```
**Note** : while creating an instance when the type is set to the patent then it will implicitly type cast to the parent , hence automatically upcasting it.

Down-cating
```java
class A {
	public void show1() {
		System.out.println("In class A");
	}
}

class B extends A {
	public void show2() {
		System.out.println("In class B");
	}
}

public class lesson {
	public static void main(String[] args) {
		A obj1 = new B(); // polymorphism
		//here the type is changed from parent to child's type
		B obj2 = (B) obj1;
	}
}
```