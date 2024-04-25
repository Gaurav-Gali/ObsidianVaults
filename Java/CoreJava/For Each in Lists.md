![Source](https://youtu.be/Z4PIhp5ifcw?list=PLsyeobzWxl7pe_IiTfNyr55kwJPWbgxB5)

Implementation of **For Each**
```java
import java.util.Arrays;

import java.util.List;

public class lesson {

	public static void main(String[] args) {
		
		List<Integer> nums = Arrays.asList(1,2,3,4,5);

		nums.forEach(num -> System.out.println(num));
	
	}

}
```

Implementing for each with a **consumer** [[Interface]]
```java
import java.util.List;
import java.util.Arrays;
import java.util.function.Consumer;

  
public class lesson {

	public static void main(String[] args) {
		
		List<Integer> nums = Arrays.asList(1,2,3,4,5);
		
		Consumer<Integer> con = new Consumer<Integer>() {
		
			public void accept(Integer num) {
			
				System.out.println(num%2==0?num+"(even)":num+"(odd)");
			
			}
		
		};
		
		nums.forEach(con);
		
		}
}
```