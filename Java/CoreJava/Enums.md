![Source](https://youtu.be/k0iTgTuiEGY?list=PLsyeobzWxl7pe_IiTfNyr55kwJPWbgxB5)

**Note** : An enum behaves like a class i.e you can create a constructor , obj etc but ==cannot== extend it like normal classes.

The keyword **enum** is short for the word *enumeration*.
you can use enums for creating named constants.

How to implement it
```java
enum Status {
	Running , Failed , Pending , Success;
}
public class lesson {
	public static void main(String[] args) {
		Status stat = Status.Running;
		System.out.println(stat)
	}
}
```
**Output** : Running

Some important methods for enms.
```java
enum Status {
	Running , Failed , Pending , Success;
}

public class lesson {
	public static void main(String[] args) {
		Status[] stat = Status.values(); // returns all consts in array format
		stat[0].ordinal() // returns index of selected named constant
	}
}
```

Using enums in conditional statements and switch case.
```java
enum Status {

Running , Failed , Pending , Success;

}

public class lesson {

	public static void main(String[] args) {
	
		Status s = Status.Running;
	
	  
	
		if (s == Status.Running) {
		
			System.out.println("Running...");
			
		} else if (s == Status.Failed) {
			
			System.out.println("Failed...");
			
		} else if (s == Status.Pending) {
			
			System.out.println("Pending...");
			
		} else if (s == Status.Success) {
			
			System.out.println("Success...");
			
		} else {
			
			System.out.println("Error...");
			
		}
	
		
	}

}
```

Using enums with switch case
```java
enum Status {

Running , Failed , Pending , Success;

}

public class lesson {

	public static void main(String[] args) {
	
		Status s = Status.Running;
		
		switch(s) {
		
			case Running -> System.out.println("Running...");
			
			case Failed -> System.out.println("Failed...");
			
			case Pending -> System.out.println("Pending...");
			
			case Success -> System.out.println("Success...");
		
		}
	
	}

}
```
**Note** : within the switch case block using object *s.running* etc was not required since within the block it automatically assumes that the named constant belongs to the condition mentioned.

Implementing enums with constructors and getters and setters
```java
enum Laptop {

	Macbook(1000) , XPS(1500) , ThinkPad(850);
	
	private int price;

	// upon declaring a constructor it will automatically assign values to each      of the named constants.
	private Laptop(int price) {
	
		this.price = price;
	
	}
	
	public int getPrice() {
	
		return price;
	
	}
	
	public void setPrice(int price) {
	
		this.price = price;
	
	}
	
}

public class lesson {

	public static void main(String[] args) {
	
		Laptop l = Laptop.Macbook;
		
		l.setPrice(2000);
		
		System.out.println(l.getPrice());
	
	}

}
```