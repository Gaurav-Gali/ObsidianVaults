![Source](https://youtu.be/1xuDEPftKV0?list=PLsyeobzWxl7pe_IiTfNyr55kwJPWbgxB5)

The main reason why we use **annotations** is to debug and resolve errors in compile time rather than finding bugs in the run time which can be ==hard to fix==

Example use of @Override annotation
```java
class A {

	public void show() {
	
		System.out.println("In A");
	
	}

}

class B extends A {

	@Override
	public void show() {
	
		System.out.println("In B");
	
	}

}

public class lesson {

	public static void main(String[] args) {
	
		A obj = new B();
		
		obj.show();
	
	}

}
```

**Note** : there are more annotations , but in core java it won't be used mostly rather it would be most often used in various frameworks like spring-boot etc.
