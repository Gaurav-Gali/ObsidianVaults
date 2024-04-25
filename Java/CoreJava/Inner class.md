![Source](https://youtu.be/UVOztdkD7WE?list=PLsyeobzWxl7pe_IiTfNyr55kwJPWbgxB5)

An **inner class** is just a [[Class And Object Theory]] within a another one.

How to implement ?
```java
class A {
	class B {
		public void show() {
			System.out.println("In B class");
		}
	}
}

public class lesson {
	public static void main(String[] args) {
		// to access the inner class object to parent is required
		A obj1 = new A();
		A.B obj2 = obj1.new B();
		obj2.show() // prints "In B class"
	}
}
```

If you don't want to access inner classes with the object of the parent class then you can make the inner class as a static class i.e it is same throughout the class .

```java
class A {
	static class B {
		public void show() {
			System.out.println("In B class");
		}
	}
}

public class lesson {
	public static void main(String[] args) {
		// to access the inner class object to parent is required
		A.B obj2 = new A.B();
		obj2.show() // prints "In B class"
	}
}
```

**Note** : only an ==inner== class can be made static , ==not== an outer class.

Anonymous inner class
```java
class A {
	public void show() {
		System.out.println("In old show");
	}
}
public class lesson {
	public static void main(String[] args) {
		A obj = new A() {
			public void show() {
				System.out.println("In New Show");
			}
		};
		Syystem.out.println(obj.show());
	}
}
```
*here the new show will be printed because the anonymous declaration has more precedence*
