![Source](https://youtu.be/Kn1RbK02YpM?list=PLsyeobzWxl7pe_IiTfNyr55kwJPWbgxB5)

What is 
1. Collection Api
	1. It's a concept about collections
2. Collection
	1. It's an [[Interface]]
3. Collections
	1. It's an [[Class And Object Theory]]

The **collection** interface in Java
1. List
	1. ArrayList
	2. LinkedList
2. Queue
	1. DeQueue
3. Set
	1. HashSet
	2. LinkedHashSet
4. Map
	1. HashMap
	2. HashTable

Implementation of ArrayList
```java
import java.util.ArrayList;

import java.util.Collection;

public class lesson {

	public static void main(String[] args) {
	
		Collection<Integer> nums = new ArrayList<Integer>();
		
		nums.add(10);
		
		nums.add(20);
		
		nums.add(30);
		
		nums.add(40);
		
		System.out.println(nums);
		
		for (int n : nums) {
		
			System.out.println(n);
		
		}
	
	}

}
```
**Note** : Use *generics* to specify the data type of the content in the ArrayList.
- Also by default the data in a Collection is considered an object.
- In ArrayList you don't have methods to get the index of elements.
- In order to avail all those functionalities use the List instead of the Collection

Implementation of List
```java
import java.util.List;

import java.util.ArrayList;


public class lesson {

	public static void main(String[] args) {
		List<Integer> nums = new ArrayList<Integer>();
		
		nums.add(10);
		
		nums.add(20);
		
		nums.add(30);
		
		System.out.println(nums);
		
		System.out.println(nums.indexOf(30));
		
		System.out.println(nums.set(1, 100));
		
		System.out.println(nums);
		
		System.out.println(nums.get(0));
		
	}

}
```

**Note** :  Arrays.asList() method can be used to create a instant List with the values as params

<hr/>

Implementation of Set
```java
import java.util.Set;

import java.util.HashSet;

public class lesson {
	
	public static void main(String[] args) {
	
		Set<Integer> nums = new HashSet<Integer>();
		
		nums.add(10);
		
		nums.add(20);
		
		nums.add(30);
			
		System.out.println(nums);
	
	}

}
```
**Note** : set does not render values in *sorted* format , some values might appear that way but usually not.

Things to remember about sets
- It holds a collection of **unique** values
- No duplicate elements
- Doesn't have any index value

If you want the values to be sorted , then instead of using a HashSet instead use a TreeSet

**Note** : Iterator is extended by the Collection interface , hence can be used by the collection objects.
It has 2 important methods
- hasNext()
	- It basically returns a boolean true if an element exists after the current element other wise false.
- next()
	- It returns the element after the current element i.e the next element in a collection.

<hr/>

Implementation of a Iterator
```java
import java.util.Iterator;

import java.util.Set;

import java.util.TreeSet;

public class lesson {

	public static void main(String[] args) {
	
		Set<Integer> nums = new TreeSet<Integer>();
		
		nums.add(10);
		
		nums.add(20);
		
		nums.add(30);
			
		Iterator<Integer> vals = nums.iterator();
	
		while(vals.hasNext()) {
		
			System.out.println(vals.next());
		
		}
	
	}

}
```

<hr/>

Maps
- store data in the form of a key value pair
- basically like dictionaries in python and object in javascript
- This does not follow any sequence

Implementation of maps
```java
import java.util.HashMap;

import java.util.Map;

public class lesson {

	public static void main(String[] args) {
	
		Map<String , Integer> students = new HashMap<>();
		
		students.put("A", 1);
		
		students.put("B", 2);
		
		students.put("C", 3);
			
		System.out.println(students.get("A"));
	
	}

}
```

Iterating through a Map
```java
import java.util.HashMap;

import java.util.Map;

public class lesson {
	
	public static void main(String[] args) {
	
		Map<String , Integer> students = new HashMap<>();
		
		students.put("A", 1);
		
		students.put("B", 2);
		
		students.put("C", 3);
		
		for (String i : students.keySet()) {
		
			System.out.println(i + " : " + students.get(i));
		
		}
	
	}

}
```

**Note** : if multiple threads are working with the same Map then using a HashTable would be better since it works synchronously.
