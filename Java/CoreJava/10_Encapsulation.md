![Source](https://youtu.be/YbqneqDIZh8?list=PLsyeobzWxl7pe_IiTfNyr55kwJPWbgxB5)

The ==private== keyword allows to set a class/instance variable private, that is the value is only accessible within that particular class.

**Note :** make all class/instance variables private in-order to secure it!

Example of ==encapsulation==
```java
// encapsulation
class Human {

	private String name = "";
	
	public void setName(String Name) {
	
		name = Name;
	
	}
	
	public String getName() {
	
		return name;
	
	}

}

  

public class lesson {

	public static void main(String[] args) {
	
		Human obj1 = new Human();
		
		obj1.setName("name1");
		  
		System.out.println(obj1.getName());
	
	}

}
```