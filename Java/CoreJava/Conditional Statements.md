![Source](https://youtu.be/74Q7POjS7mQ?list=PLsyeobzWxl7pe_IiTfNyr55kwJPWbgxB5)

How to write a conditional statement
```java
class hello {

	public static void main(String[] args) {
	
		int a = 10;
		
		int b = 20;
		
		if (a >= 18 && b <= 18) {
		
			System.out.println("Both Are above 18");
		
		} else if (a >= 18 && b <= 18) {
		
			System.out.println("Only a is above 18");
		
		} else if (a <= 18 && b >= 18) {
		
			System.out.println("Only b is above 18");
		
		} else {
		
			System.out.println("Both are under 18");
	
		}

	}

}
```

**Ternary Operator**
```java
class hello {

	public static void main(String[] args) {
	
		int a = 10;
		
		int b = 20;
		
		int greater = (a>b) ? a : b;
		
		System.out.println(greater);
	
	}

}
```

**Switch Statement Standard**
```java
class hello {

	public static void main(String[] args) {
	
		int num = 10;
			
		switch(num) {
		
			case 10 :
			
				System.out.println("Num is 10");
			
				break;
			
			case 20:
			
				System.out.println("Num is 20");
			
				break;
			
			default:
			
				System.out.println("Num is not 10 or 20");
		
		}
	
	}

}
```

**Switch Statement without ==break== keyword**
```java
class hello {

	public static void main(String[] args) {
	
		int num = 10;
		
		switch(num) {
		
			case 10 -> System.out.println("Num is 10");
			
			case 20 -> System.out.println("Num is 20");
			
			default -> System.out.println("Num is not 10 or 20");
		
		}
	
	  
	}

}
```

**Switch Statement assigned to a variable**
```java
class hello {

	public static void main(String[] args) {
	
		int num = 10;
		
		String output = switch(num) {
		
			case 10 -> "Num is 10";
			
			case 20 -> "Num is 20";
			
			default -> "Num is not 10 or 20";
		
		};
			
		System.out.println(output);
	
	}

}
```

**Switch Statement using the ==yield== keyword**
```java
class hello {

	public static void main(String[] args) {
	
		int num = 10;
		
		String output = switch(num) {
		
			case 10 : yield "Num is 10";
			
			case 20 : yield "Num is 20";
			
			default : yield "Num is not 10 or 20";
		
		};
			
		System.out.println(output);
	
	}

}
```