![Source](https://youtu.be/Gs8ZPKCFlTc?list=PLsyeobzWxl7pe_IiTfNyr55kwJPWbgxB5)

A **Functional Interface** is just an [[24_Interface]] with only one method or function.
**Note** : A *functional interface* is also called as a S.A.M (Single Abstract Method) interface.

In order to make sure that the SAM has only one method you can make use of an [[26_Annotation]]
which is **@FunctionalInterface**

Example
```java
@FunctionalInterface
interface A {

	void show();

}

class B implements A {

	public void show() {
	
		System.out.println("In Show");
	
	}

}

public class lesson {
	
	public static void main(String[] args) {
	
		A obj = new B();
		
		obj.show();
	
	}

}
```

Example with Anonymous [[23_Inner class]]
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