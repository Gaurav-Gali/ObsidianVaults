![Source](https://youtu.be/AG_7wWFBquQ?list=PLsyeobzWxl7pe_IiTfNyr55kwJPWbgxB5)

**Note** : ==Interface== is not a class ! rather all the methods by *default* are considered **public** and **abstract**

**Points to remember** upon inheriting
- ==class -- class== : *extends*
- ==class -- interface== : *implements*
- ==interface -- interface==: *extends*

Also when using [[Class And Object Theory]] you create objects whereas in interfaces we create [[Data Types]]

With in an interface you can also declare variables but those variables are
1. Final [[Final Keyword]]
2. Static : *the value of something static will remain the same throughout all it's instances*

How to implement an Interface
```java
interface A {
	int age = 10; // since variable is final and static the value is set.
	String name = "Name";
}

public class lesson {
	public static void main(String[] args) {
		
	}
}
```
**Note** : instance of a *interface* cannot be cerated , hence a class must inherit the interface to start creating objects.

```java
interface A {
	void show();
	void config();
}

class B implements A {
	public void show() {
		System.out.println("In Show");
	}
	public void config() {
		System.out.println("In Config");
	}
}

public class lesson {
	public static void main(String[] args) {
		B obj = new B();
		obj.show();
		obj.config();
	}
}
```
**Note** : instead of the *extends* keyword use the *implements* keyword to inherit interfaces

Also more than one interface can be implemented with a class , just that all the methods defined inside a interface must be defined in that class too.
```java
interface A {
	void show();
}
interface B {
	void config();
}

class New implements A,B {
	public void show() {
		System.out.println("In Show");
	}
	public void config() {
		System.out.println("In Config");
	}
}

public class lesson {
	public static void main(String[] args) {
		
	}
}
```

If you want to inherit an interface to another interface simply use the **extends** keyword.