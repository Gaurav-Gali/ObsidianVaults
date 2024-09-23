![Source](https://youtu.be/8C_YRYXCuwc?list=PLsyeobzWxl7pe_IiTfNyr55kwJPWbgxB5)

Dynamic Method Dispatch is used to implement run-time [[16_Polymorphism]]

How it works ?
- Basically when creating an instance of a class , if that class has a child i.e some other class extends it then while creating an instance the [[1_Data Types]] of the instance can be the parent class and the reference class can be the child class that extends the parent.

```java
class Computer {
	public void show() {
		System.out.println("In Computer Class (parent)")
	}
}

class Laptop extends Computer{
	public void show() {
		System.out.println("In Laptop Class (child)")
	}
}

public class lesson {
	public static void main(String[] args) {
		// regular way of creating a class instance
		Computer obj1 = new Computer();
		obj1.show();
		
		// Dynamic way of creatign a class instance
		obj1 = new Laptop();
		obj1.show();
	}
}
```

- Output of the first show() method : In Computer Class (parent)
- Output of the second show() method : In Laptop Class (child)

**Note :** this is a type of [[7_Method Overloading]]

In conclusion the same method show() is behaving differently with different instances (a.k.a) [[16_Polymorphism]] !
