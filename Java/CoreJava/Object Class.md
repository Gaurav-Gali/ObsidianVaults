![Source](https://youtu.be/dCt9sfZV8Sg?list=PLsyeobzWxl7pe_IiTfNyr55kwJPWbgxB5)

Every class in java extends the *object class*
```java
class Computer {

	int price;
	
	String model;

}

public class lesson {

	public static void main(String[] args) {
	
		Computer obj1 = new Computer();
		
		System.out.println(obj1);
	
	}

}
```
Output : Computer@7ad041f3 , here Computer is the class name and 7ad041f3 is the hexadecimal value of the *hash code*.

```java
class Computer {

	int price;
	
	String model;

}

public class lesson {

	public static void main(String[] args) {
	
		Computer obj1 = new Computer();
		
		System.out.println(obj1);
		System.out.println(obj1.toString());
	
	}

}
```

**Note** : both the print methods will print the same output which is *Computer@7ad041f3*
This is because by default the object calls the ==toString()== method.

What is **hash code** ?
- Hash code is the encryption technique
- It converts the data in a class to a single string or entity

Different methods for the object class
- toString()
	- returns the string of the *[classname] + "@" + [hashcode]*
- equals()
	- returns a boolean after checking if the two objects are equal.