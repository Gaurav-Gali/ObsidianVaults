![Source](https://youtu.be/Bua6LQO2vQ8?list=PLsyeobzWxl7pe_IiTfNyr55kwJPWbgxB5)

**Packages** are a way of organising and putting one type of data in a separate folder.

How to mention the respective *package* the file belongs to.
```java
package arith;

public class Calc {

	public int add(int n1 , int n2) {
	
		return n1+n2;
	
	}

}
```

How to import files from other packages
```java
import arith.Calc;

import arith.NewCalc;

public class lesson {

	public static void main(String[] args) {
		// Calc instance
		
		Calc c = new Calc();
		
		int sum = c.add(10,20);
		
		System.out.println(sum);
		
		  
		
		// NewCalc instance
		
		NewCalc c1 = new NewCalc();
		
		double sum1 = c1.add(10.5,20.0);
		
		System.out.println(sum1);
	
	}

}
```

**Note :** since *Calc* and *NewCalc* are the only packages in the **arith** package

```java
import arith.*; // this can be used to import all files in that package
```

The file and folder structure
![[Media/files.png]]