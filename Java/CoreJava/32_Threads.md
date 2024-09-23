![Source](https://youtu.be/UfMM924sBvg?list=PLsyeobzWxl7pe_IiTfNyr55kwJPWbgxB5)

A thread is a part of a cpu where it would take up tasks.
Generally while multitasking , multiple threads would be utilised.

How to implement threads in java
```java
class A extends Thread {

	public void run() {
	
		for (int i = 0 ; i < 100 ; i++) {
		
			System.out.println("In A");
		
		}
	
	}

}

class B extends Thread {

	public void run() {
	
		for (int i = 0 ; i < 100 ; i++) {
		
			System.out.println("In B");
		
		}
	
	}

}


public class lesson {

	public static void main(String[] args) {
	
		A obj1 = new A();
		
		B obj2 = new B();
		
		  
		
		obj1.start();
		
		obj2.start();
	
	}

}
```

==Points to be noted==
1. A class in order to be executed in another thread , it must be extending the thread class.
2. Only one method is allowed in that class.
3. The method's name ==must== be *run*.
4. An object of that [[5_Class And Object Theory]] can be created as usual.
5. While calling the *run* method use the method *start()* to start execution in a new thread.

What is a **scheduler** ?
- A scheduler allows a thread to execute with a set time interval
- basically it splits the time of execution of different tasks among the available threads.

Thread Priority
- Range of priority
	- Minimum : 1
	- Maximum : 10
	- Default : 5
- Note : while setting the thread priority it doesn't set the priority rather it suggests that idea to the scheduler.
```java
class A extends Thread {

	public void run() {
	
		for (int i = 0 ; i < 100 ; i++) {
		
			System.out.println("In A");
		
		}
	
	}

}

class B extends Thread {

	public void run() {
	
		for (int i = 0 ; i < 100 ; i++) {
		
			System.out.println("In B");
		
		}
	
	}

}


public class lesson {

	public static void main(String[] args) {
	
		A obj1 = new A();
		
		B obj2 = new B();
		
		obj1.setPriority(Thread.MAX_PRIORITY)
		
		obj1.start();
		
		obj2.start();
	
	}

}
```
**Note** : direct integer values from 1 - 10 can be set in the setPriority method.


Sleep in Threads
```java
class A extends Thread {
	public void run() {
		for (int i = 0 ; i < 100 ; i++) {
			System.out.println("In A");
			try {
				Thread.sleep(10);
			} catch (InteruptedException e) {
				e.printStackTrace();
			}
		}
	}
}

class B extends Thread {
	public void run() {
		for (int i = 0 ; i < 100 ; i++) {
			System.out.println("In B");
			try {
				Thread.sleep(10);
			} catch (InteruptedException e) {
				e.printStackTrace();
			}
		}
	}
}

public class lesson {
	public static void main(String[] args) {
		A obj1 = new A();
		B obj2 = new B();

		obj1.setPriority(Thread.MAX_PRIORITY);
		
		obj1.start();
		obj2.start();
	}
}
```

**Note** : The *sleep* method under Thread class might return an *InteruptedException* so always wrap the implementation of the sleep method within a try-catch block.
Also the sleep method accepts time in **milli-seconds**.

**Race Conditions**
- This commonly occurs when 2 threads work with the same mutable data.
- Basically it is caused when 2 threads are trying to mutate the same data in the exact same time.
- Prevention
	- Making the mutable data *thread safe*
	- in order to implement this use the *synchronized* keyword with the method that is trying to mutate the data
	- The *join()* method can also be used which will wait until the execution of another thread is completed.


**Thread States**
1. New
	1. This is when the instance of the thread is created which is done with the *start()* method
2. Runnable
	1. This is when the thread is ready to execute but is put on wait by the scheduler
3. Running
	1. This is when the thread is actually executing its tasks with the *run()* method
4. Waiting
	1. This is when the thread is not ready for execution , usually done using *sleep()* method.
	2. The *wait()* method can also be used , and the *notify()* method can be used to bring the thread back to it's runnable state.
5. Dead
	1. This is when the thread is done with it's execution , can be done with the stop() method.