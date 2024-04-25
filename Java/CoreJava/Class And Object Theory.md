![Source](https://youtu.be/3yOLNV9BF8A?list=PLsyeobzWxl7pe_IiTfNyr55kwJPWbgxB5)

**Note -** The ==public== keyword is used to make the particular class of a function visible/accessible from elsewhere

How to design a class
```java
public class hello {

	public static void main(String[] args) {
	
		// Arithmetic class instance
		
		Arithmetic arith = new Arithmetic();
		
				
		// Arithmetic class objects
		
		int sum1 = arith.add(10, 20);
		
		int sum2 = arith.add(20, 30);
		
	
		// Output
		
		System.out.println(sum1);
		
		System.out.println(sum2);
	
	
	}

}

class Arithmetic {

	public int add(int x, int y) {
	
		return x+y;
	
	}

}
```

**Variable Arguments**
```java
import java.util.ArrayList;

import java.util.List;

  

class maths {

	List<Integer> numArr = new ArrayList<Integer>();
	
	public maths(int ...nums) {
	
		for (int i : nums) {
		
			numArr.add(i);
		
		}
	
	}
	
	public int sum() {
	
		return numArr.stream().reduce(0, (a,b) -> a + b);
	
	}
	
	public void show() {
	
		numArr.forEach(n -> System.out.println(n));
	
	}

}

public class lesson {

	public static void main(String[] args) {
	
		maths m = new maths(1,2,3,4,5);
		
		int s = m.sum();
		
		System.out.println(s);
	
	}

}
```
**In Short** : use "..." before the argument's name in order to accept variable arguments
**Note** : the variable arguments are received in the form of an array.