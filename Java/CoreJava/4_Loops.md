![Source](https://youtu.be/mzt5tmV7wxI?list=PLsyeobzWxl7pe_IiTfNyr55kwJPWbgxB5)


There are 3 types of loops in JAVA
1. while loop
2. do-while loop
3. for loop (preferred)
4. for-each loop (only works with sequence data types)

How to write a while loop
```java
	class hello {
	
	public static void main(String[] args) {
	
		int i = 1;
		
		while(i <= 5) {
		
			System.out.println("This is iteration " + i);
			
			i++;
		
		}
	
	}

}
```

How to write a do-while loop
```java
class hello {

	public static void main(String[] args) {
	
		int i = 6;
		
		do {
		
			System.out.println("Iteration " + i);
			
			i++;
		
		} while(i <= 5);
	
	}

}
```


How to write a for loop
```java
class hello {

	public static void main(String[] args) {
	
		for (int i=1; i <= 5 ; i++) {
		
			for (int j=0; j < i ; j++) {
			
				System.out.print("*");
			
			}
			
			System.out.println();
		
		}
	
	}

}
```

How to write a for each loop
```java
public class lesson

{

	public static void main(String[] args)
	
	{
	
		int nums[] = new int[3];
		
		nums[0] = 1;
		
		nums[1] = 2;
		
		nums[2] = 3;
		
		  
		
		for (int i : nums)
		
		{
		
			System.out.println(i);
		
		}
	
	}

}
```