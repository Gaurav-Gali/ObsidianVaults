![Source](https://youtu.be/bwHr9G5VIls?list=PLsyeobzWxl7pe_IiTfNyr55kwJPWbgxB5)

**Note** : 
- There exist default class within the **System** class called *in* which has a method *read* to read inputs and returns an integer value.
	- The problem with this is that it returns the ascii value of the input which you have to then convert to the desired result.
	- Also it only scans one byte at a time , so in order to read larger values you have to use loops which gets quite complex for simple input functionality.

<hr/>

The old way of doing this (not recommended 😨)
```java
public class lesson {
	public static void main(String[] args) {
		System.out.println("Enter a numer : ");

		InputStreamReader in = new InputStreamReader(System.in);
		BufferReader bf = new BufferReader();

		int num = Integer.parseInt(bf.readLine());
		System.out.println(num)

		bf.close()
	}
}
```

The modern way of doing this (very much recommended ☺️)
```java
import java.util.Scanner;

public class lesson {

	public static void main(String[] args) {
	
		int inp;
		
		System.out.print("Enter a number : ");
		
		  
		
		Scanner sc = new Scanner(System.in);
		
		inp = sc.nextInt();
		
		System.out.println("Your number is : "+inp);
	
	}

}
```
