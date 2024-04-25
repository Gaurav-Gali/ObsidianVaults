![Source](https://youtu.be/Fyc86kVIePE?list=PLsyeobzWxl7pe_IiTfNyr55kwJPWbgxB5)

**Boxing** : Storing a primitive value into a wrapper object
**Note** : when boxing is done automatically (implicitly) then its called **auto-boxing**

**Un-boxing** : Retrieving a boxed value
**Auto Un-boxing** : Un-boxing automatically (implicitly)

Also for every primitive class there exist a class for it
- int : Integer
- char : Character
- double : Double

How to use classes for primitive data types and boxing
```java
public class lesson {
	public static void main(String[] args) {
		// primitives
		int age = 18;
		// class based declaration (old and depricated) ![don't use]
		Integer age = new Integer(18);
		// class based declaration (recomended)
		Integer age = 18;
	}
}
```

Methods associated with these classes
```java
public class lesson {
	public static void main(String[] args) {
		Integer age = 18;
		// since age is now a object so the value cannot be directly printed
		int ageVal = age.intValue();

		String a = "123";
		int aInt = Integer.parseInt(a);
	}
}
```