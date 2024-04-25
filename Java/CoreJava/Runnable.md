![Source](https://youtu.be/Z4KSgLpY0d8?list=PLsyeobzWxl7pe_IiTfNyr55kwJPWbgxB5)

A **runnable** is an alternative way of creating a [[Threads]].
This is because in java you cannot extend multiple classes , hence in order for a thread class to be able to extend other classes you can use runnable because it is an interface.

**Note** : The thread class itself implements from the Runnable interface.

Implementation
```java
class A implements Runnable {

	public void run() {
	
		for(int i = 9 ; i < 5 ; i++) {
		
			System.out.println("In A");
			
			try {
			
				Thread.sleep(10);
			
			} catch (InterruptedException e) {
			
				e.printStackTrace();
			
			}
			
		}
	
	}

}

  

class B implements Runnable {

	public void run() {
	
		for(int i = 9 ; i < 5 ; i++) {
		
			System.out.println("In B");
			
			try {
			
				Thread.sleep(10);
			
			} catch (InterruptedException e) {
			
				e.printStackTrace();
			
			}
		
		  
		
		}
	
	}

}

  

public class lesson {

	public static void main(String[] args) {
	
		Runnable obj1 = new A();
		
		Runnable obj2 = new B();
		
		  
		
		Thread t1 = new Thread(obj1);
		
		Thread t2 = new Thread(obj2);
		
		t1.start();
		
		t2.start();
	
	}

}
```