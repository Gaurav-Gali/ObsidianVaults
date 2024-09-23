![Source](https://youtu.be/5r_ERSm7NKE?list=PLsyeobzWxl7pe_IiTfNyr55kwJPWbgxB5)

Types of Errors
1. Compile Time Errors
	- Error cause due to semantics or incorrect syntax.
	- Error will be pointed out by the compiler
2. Runtime Errors
	- Error caused during the execution of the program
	- Ex : division of *0/0*
3. Logical Error
	- Error caused when the desired output is not obtained
	- that is , the error is caused due to a logical flaw caused by the programmer

<hr/>

Handling the exceptions
```java
public class lesson {

	public static void main(String[] args) {

		int i = 40;
		
		int j = 10;
		
		try {
			int res = (int) (i/j);
			System.out.println(res);
		}
		catch(Exception e) {
			System.out.println("Err...");
		}
		
	}
}
```

Try Catch with finally block
```java
public class lesson {

	public static void main(String[] args) {
	
		int i = 0;
		
		int j = 0;
		
		try {
		
			j = 18/i;
			
			System.out.println("No Err");
		
		} catch(Exception e) {
		
			System.out.println("Err : " + e);
		
		} finally {
		
			System.out.println("Done!");
		
		}
		
		
	}
}
```

**Note** : the *finally* block gets executed irrespective of try or catch block , upon either of those's execution the finally block executes.

Common exception cases
1. ArithmeticException
2. ArrayIndexOutOfBoundsException
```java
public class lesson {
	
	public static void main(String[] args) {
	
		int arr[] = new int[5];
		
		try {
		
			System.out.println(10/20);
			
			System.out.println(arr[5]);
		
		} catch(ArithmeticException e) {
		
			System.out.println("Math Err !!");
		
		} catch(ArrayIndexOutOfBoundsException e) {
		
			System.out.println("Arr Err !!");
		
		} catch(Exception e) {
		
			System.out.println("Random Err !!");
		
		}
	
	}

}
```

Exception Hierarchy
1.  Object
2. Throwable
	- Error
		- These are such that it cannot be resolved by the developer
		- Ex : IO/Error , Out of memory , Thread death etc.
	- Exception
		- These are such that it can be handled
		- Different types of exceptions
			- Runtime (a.k.a) **unchecked exceptions**
				- ArithmeticException
				- ArrayIndexOutOfBoundsException
				- NullPointerException
			- SQLException
			- IO Exception

<hr/>

Throw And Throws Keyword

```java
public class lesson {

	public static void main(String[] args) {
	
		int i = 0;
		
		int j = 20;
		
		try {
		
			j /= i;
		
			if (i == 0) {
			
				throw new ArithmeticException("zero division Err !!");
			
			}
		
		} catch (ArithmeticException e) {
		
			System.out.println("Err : " + e);
		
		}
	
	}

}
```

<hr/>

Custom Exceptions
```java
class RandomException extends Exception {

	public RandomException(String message) {
	
		super(message); // sending the message to the constructor of the parent class which is the Exception class.
	
	}

}

public class lesson {

	public static void main(String[] args) {
	
		try {
		
			int i = 2;
			
			if (i == 2) {
			
				throw new RandomException("Random Exception");
			
			}
			
			System.out.println(20/i);
		
		} catch(RandomException e) {
		
			System.out.println("Err : " + e);
		
		}
	
	}

}
```

Implementation of throws
```java
class A {

	public void divide(int x , int y) throws ArithmeticException {
	
		System.out.println(x/y);
	
	}

}

public class lesson {

	public static void main(String[] args) {
	
		A obj = new A();
		
		try {
		
			obj.divide(0, 0);
		
		} catch(ArithmeticException e) {
		
			System.out.println("Div by Zero : " + e);
		
		}
		
	}

}
```

