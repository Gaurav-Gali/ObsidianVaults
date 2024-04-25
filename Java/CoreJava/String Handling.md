![Source](https://youtu.be/Ae-r8hsbPUo?t=6049)

Utility methods for string handling

```java
public class prog {

	public static void main(String[] args) {
	
		String s1 = "This is string 1";
		String s2 = "This is string 2";
		
		// 1.Length of a string
		System.out.println(s1.length());

		// 2.string concatination
		System.out.println(s1.concat(s2));

		// 3.string formatting
			// %s is as a "String" placeholder
			// %d is used as a "digit" placeholder
			// %f is used as a "float" placeholder
			// %b is used as a "boolean" placeholder
		System.out.println(String.format("%s is a fruit" , "apple"));

		// 4.finding the character in the given index
		System.out.println(s1.charAt(0));
		
		// 5.finding relationship between strings
		System.out.println(s1.equals(s2));

		// 6. finding the index of the character
			// if the character is not in the string then it returns -1
			// if the character occurs more than once then the first occurance's                index is returned String
		System.out.println(s1.indexOf("i"));

		// 7. string manipulation
		System.out.println(s1.replace("i" , "s"));

		// 8.splitting the string
		
		for(String s : s1.split(" ")) {
		
			System.out.println(s);
		
		}

		// 9.substring the string
			// the last limit won't be counted for the range
		System.out.println(s1.substring(1,3));

	}

}
```
