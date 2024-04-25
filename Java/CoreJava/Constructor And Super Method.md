![Source](https://youtu.be/UC_aqNUFyVI?list=PLsyeobzWxl7pe_IiTfNyr55kwJPWbgxB5)

**Note :** Every class in java extends from the ==Object== class by default and has a lot of methods.

A **Constructor** is a special method that calls itself when an object of that class is created.

Example
```java
class Human {

	private String name;
	
	private int age;
	
	public Human(String name, int age) {
	
		this.name = name;
		
		this.age = age;
	
	}
	
	public String getData() {
	
		return this.name + " && "+ this.age;
	
	}

}

public class lesson {

	public static void main(String[] args) {
	
		Human obj = new Human("Human1",10);
		
		System.out.println(obj.getData());
	
	}

}
```


The **super()** method is called by default in every class. It is used to access the constructor of the parent class.

Also the **this()** method will call the constructor of the same class. 

Example
```java
class Human {

	private String name;
	
	private int age;
	
	  
	
	public Human() {
	
		this.name = "";
		
		this.age = 0;
	
	}
	
	public String getData() {
	
		return this.name + " && "+ this.age;
	
	}

}

class Job extends Human{

	public Job() {
	
		super();
		
		System.out.println("In Job");
	
	}

}

public class lesson {

	public static void main(String[] args) {
	
		Human obj = new Human();
		
		System.out.println(obj.getData());
		
	}

}
```