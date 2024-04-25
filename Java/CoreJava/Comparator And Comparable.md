![Source](https://youtu.be/ZA2oNhtNk3w?list=PLsyeobzWxl7pe_IiTfNyr55kwJPWbgxB5)

This works with the collections class with the sort method.

**Note** : as mentioned in [[Collection Api]] don't confuse Collection(Interface) with Collections(Class).

Implementing the sort method with Collections class
```java
import java.util.ArrayList;

import java.util.Collections;

import java.util.Comparator;

import java.util.List;

public class lesson {

	public static void main(String[] args) {
	
		List<Integer> nums = new ArrayList<>();
		
		nums.add(19);
		
		nums.add(33);
		
		nums.add(27);
		
		
		Collections.sort(nums);
		
		
		System.out.println(nums);
	
	}

}
```

If you want to sort using you custom method then use a comparator
Example : ==sorting based on the last digit of the number==
```java
import java.util.ArrayList;

import java.util.Collections;

import java.util.Comparator;

import java.util.List;

  

public class lesson {

	public static void main(String[] args) {
	
		List<Integer> nums = new ArrayList<>();
		
		Comparator<Integer> com = new Comparator<Integer>() {
		
			public int compare(Integer i , Integer j) {
			
				if (i%10 > j%10) {
				
					return 1;
				
				} else {
				
					return -1;
				
				}
			
			}
		
		};
		
		nums.add(19);
		
		nums.add(33);
		
		nums.add(27);
		
		Collections.sort(nums , com);
		
		System.out.println(nums);
	
	}

}
```

**Note** : by default the *Integer* class implements the ==Comparable==  interface , hence in order to use the **comperator** with any other [[Class And Object Theory]] then implement that class with the Comparable [[Interface]].