![Source](https://youtu.be/ak3BxYzSqsQ?list=PLsyeobzWxl7pe_IiTfNyr55kwJPWbgxB5)

The **Stream Api** can be used to provide additional methods for working with sequence data types.
**Note** : If the stream object is used once then it can't be reused.

Implementation
1. using [[For Each in Lists]]
```java
import java.util.Arrays;

import java.util.List;

import java.util.stream.Stream;

public class lesson {

	public static void main(String[] args) {
	
		List<Integer> nums = Arrays.asList(1,2,3,4,5);
		
		Stream<Integer> s1 = nums.stream();
		
		s1.forEach(n -> System.out.println(n));
	
	}

}
```

2. using **filter()** method
```java
import java.util.Arrays;

import java.util.List;

import java.util.stream.Stream;

public class lesson {
	
	public static void main(String[] args) {
	
		List<Integer> nums = Arrays.asList(1,2,3,4,5);
		
		Stream<Integer> s1 = nums.stream();
		
		Stream<Integer> s2 = s1.filter(n -> n%2==0);
				
		s2.forEach(n -> System.out.println(n));
	
	}

}
```

3. using **map()** method
```java
import java.util.Arrays;

import java.util.List;

import java.util.stream.Stream;

public class lesson {

	public static void main(String[] args) {
	
		List<Integer> nums = Arrays.asList(1,2,3,4,5);
		
		Stream<Integer> s1 = nums.stream();
		
		Stream<Integer> s2 = s1.map(n -> n*2);
		
		  
		
		s2.forEach(n -> System.out.println(n));
	
	}

}
```
**Note** : the map method can be used to modify the values within the stream based on the condition , whereas in forEach() method the values cannot be modified.

4. using **reduce()** method
```java
import java.util.Arrays;

import java.util.List;

import java.util.stream.Stream;

public class lesson {

	public static void main(String[] args) {
	
		List<Integer> nums = Arrays.asList(1,2,3);
		
		Stream<Integer> s1 = nums.stream();
		
		int s2 = s1.reduce(0 , (a,b) -> a+b);
		
		System.out.println(s2);
	
	}

}
```
**Note** : The *reduce()* takes 2 arguments
	1. *0* is recommended
	2. The operation to be performed

2. using **sorted()** method

```java
import java.util.Arrays;

import java.util.List;

import java.util.stream.Stream;

public class lesson {

	public static void main(String[] args) {
	
		List<Integer> nums = Arrays.asList(1,2,3,4,5,6,7,8,9,0);
		
		Stream<Integer> s1 = nums.stream()
		
		.filter(n -> n%2==0)
		
		.sorted();
		
		s1.forEach(n -> System.out.println(n));
		
	}

}
```

**Note** : If you want to make use of threads then use *parallelStream()* instead of *stream()*.
Also when using *parallelStream()* you can't use *sorted()* method.