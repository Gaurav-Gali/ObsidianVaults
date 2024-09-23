![Source](https://youtu.be/VJh2u7NLLDg?list=PLsyeobzWxl7pe_IiTfNyr55kwJPWbgxB5)

The **abstract** keyword is used when you want to declare a method but don't know how to implement it.
**Note** : an *abstract* method can only be declared within a abstract class

How to implement it ?
```java
abstract class Robot {
	public abstract void abilities();
}

public class lesson {
	public static void main(String[] args) {
		
	}
}
```

==**IMPORTANT**== : You can't create an object of a abstract class

hence if you want to create an object then simply *inherit* that class with another class and create an object of that instead.