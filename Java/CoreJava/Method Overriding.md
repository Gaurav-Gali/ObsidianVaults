![Source](https://youtu.be/CLzgS08equQ?list=PLsyeobzWxl7pe_IiTfNyr55kwJPWbgxB5)

While **inheriting** from another class, the child class can have a same method so that when the child class is accessed it would give preference to the method of that class.

How it works
```java
class Human {
	private String name;
	public String name(String name) {
		this.name = name;
		return name;
	}
}

class More extends Human {	
	public void name(String name) {
		System.out.println(name);
	}
}

class lesson {
	public static void main(String[] args) {
		Human obj1 = new Human(); // class instance for Human
		More obj2 = new More(); // class instance for More

		System.out.println(obj1.name("Name")) // calls name in Human class
		obj2.name("Name") // calls name in the More class.
	}
}
```