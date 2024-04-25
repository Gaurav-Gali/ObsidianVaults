![Source](https://youtu.be/KpwBVAYbPDA?list=PLsyeobzWxl7pe_IiTfNyr55kwJPWbgxB5)

What is ==Method Overloading==
- When a class contains multiple methods with the same name but has :
	- **different** arguments
	- same arguments with different data types

```java
class Human {
	public int add(int x , int y) {
		return x+y;
	}
	public int add(int x , int y , int z) {
		return x+y+z;
	}
}

class lesson {
	public static void main(String[] args) {
		Human obj1 = new Human(); //class instance
		System.out.println(obj1.add(1,2)); // calls the first add method
		System.out.println(obj1.add(1,2,3)); // calls the second add method 
	}
}
```