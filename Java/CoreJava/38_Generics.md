![Source](https://youtu.be/XMvznsY02Mk)
The concept of **Generics** allows you to specify data types as per requirement.

Implementing Generics
```java
import java.util.ArrayList;

import java.util.List;


class Maths<T extends Number> {

	List<T> valuesArr = new ArrayList<T>();
	
	public Maths(T ...values) {
	
		for(T i : values) {
		
		valuesArr.add(i);
		
		}
		
	}
	
	public void show() {
	
		valuesArr.forEach(n -> System.out.println(n));
	
	}
	
}

  

public class lesson {

	public static void main(String[] args) {
	
		Maths<Integer> m = new Maths<Integer>(1,2,3,4,5);
		
		  
		
		m.show();
	
	}

}
```
**Note** : according to convention use a ==capital== letter while defining the type in the angular brackets.