![Source](https://youtu.be/aecXHkZ-kJY?list=PLsyeobzWxl7pe_IiTfNyr55kwJPWbgxB5)

**Note** : this is a feature from JAVA version **8** onwards.
This concept is used under [[27_Functional Interface]]  **only !**

```java
@FunctionalInterface
interface A {

void show();

}

public class lesson {

	public static void main(String[] args) {
	
		A obj = new A() {
		
			public void show() {
			
				System.out.println("In Show");
			
			}
		
		};
		
		obj.show();
	
	}

}
```

In the above code , while declaring the object of the interface A , you again have to write the following : 
```java
new A() {
	public void show() {
		System.out.println("In Show")
	}
}
```
But you already have declared all of that while making the interface.

So in order to prevent this situation **Lambda** expressions can be used.

Example 
```java
@FunctionalInterface
interface A {

	void show();

}

public class lesson {

	public static void main(String[] args) {
	
		A obj = () -> {
		
			System.out.println("In Show");
		
		};
		
		obj.show();
		
	}

}
```

With an argument passed into the function
```java
@FunctionalInterface
interface A {

	void show(int i);

}

public class lesson {

	public static void main(String[] args) {
	
		A obj = (int i) -> {
		
			System.out.println("In Show : "+i);
		
		};
				
		obj.show(10);
		
	}

}
```

**Note** : you don't have to mention the type of the argument while declaring it.
also when you have only one argument you don't even have to use *()* these brackets.

Example
```java
@FunctionalInterface
interface A {

	void show(int i);

}

public class lesson {

	public static void main(String[] args) {
	
		A obj = i -> System.out.println("In Show : "+i);
				
		obj.show(10);
		
	}

}
```