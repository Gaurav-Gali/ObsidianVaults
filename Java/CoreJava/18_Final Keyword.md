![Source](https://youtu.be/OBYeDzgsOxc?list=PLsyeobzWxl7pe_IiTfNyr55kwJPWbgxB5)

The **final** keyword is used to create constants i.e once it is used then that can be modified.

And it can be used with the following
- variable
```java
public class lesson {
	public static void main(String[] args) {
		final int PI = 3.14;
		PI = 2.14; // this will lead to an error since PI is now a constant
	}
}
```
- method
```java
class Human {
	public final void show() {
		System.out.println("In Human")
	}
}

class Occupation extends Human {
	// this show method will raise error because final keyword in the show() method of the parent (Human) class will restrict method Over riding.
	public void show() {
		System.out.println("In Occupation")
	}
}

public class lesson {
	public static void main(String[] args) {
		Human obj1 = new Occupation();
	}
}
```
- class
```java
final class Human {
	public void show() {
		System.out.println("In Human")
	}
}

// this class will show error since the final keyword in parent (Human) restricts inheritance
class Occupation extends Human {
	public void show() {
		System.out.println("In Occupation")
	}
}

public class lesson {
	public static void main(String[] args) {
		Human obj1 = new Occupation();
	}
}
```

In conclusion the **final** keyword , 
- with variables will turn them into constants.
- with methods will restrict [[13_Method Overriding]]
- with classes will restrict Inheritance. *refer :* [[5_Class And Object Theory]]