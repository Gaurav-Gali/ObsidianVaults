![Source](https://youtu.be/239ubH043lI?list=PLsyeobzWxl7pe_IiTfNyr55kwJPWbgxB5)

How to create an ==Array== 
```java
public class hello {

	public static void main(String[] args) {
	
		int arr[] = {1, 2, 3, 4, 5};
		
		System.out.println(arr[1]);
		
		arr[1] = 10;
		
		System.out.println(arr[1]);
	
	}

}
```

How to create an ==2D/Multidimensional Array==
```java
public class hello {

	public static void main(String[] args) {
	
		int arr[][] = new int[3][4];
		
		for (int i=0 ; i < arr.length ; i++) {
			
			for (int j=0 ; j < arr[i].length ; j++) {
			
				System.out.print(arr[i][j] + " ");
			
			}
			
			System.out.println();
			
			}
			
	}

}
```