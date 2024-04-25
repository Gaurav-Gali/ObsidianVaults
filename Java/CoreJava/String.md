![Source](https://youtu.be/cV-sOpOgof8?list=PLsyeobzWxl7pe_IiTfNyr55kwJPWbgxB5)

**Note** : ==String== is a class and have multiple methods to perform actions to it

Common Declaration (Preferred)
```java
public class lesson {
	public static void main(String[] args) {
		String name = "NAME";
	}
}
```

Correct Declaration using the *constructor*
```java
public class lesson {
	public static void main(String[] args) {
		String name = new String("NAME");
	}
}
```

Since **String** is immutable so every new instance will occupy new space in the ==heap==
In order to prevent this we use the **StringBuffer**

Declaration of StringBuffer
```java
public class lesson {

	public static void main(String[] args) {
	
		StringBuffer name = new StringBuffer("NAME");
	
	}

}
```
**Note :** by default 16 bytes will be allocated using this method.

Some Methods in StringBuffer
- capacity()
	- returns the byte size of the buffer
- append(string)
	- appends the argument into the string
- toString()
	- returns the entire buffer as a single string
- deleteCharAt(index)
	- deleted a character at the parameter index
- insert(index , string)
	- inserts a string in the mentioned index in the buffer

The ==static== block allows to declare static variables and calls it only once