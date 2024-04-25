![Source](https://youtu.be/HK6gnkQIgqI?list=PLu0W_9lII9agpFUAlPFe_VNSlXW5uE0YL)

**Note** : 
1. The compiler won't allow any  [[Function]] to access a class's private data which is known is [[Encapsulation]].
2. In order for a function to access a class's private data , the class must make that particular function as ==friend==.
3. To make a function as a class's friend we use the **friend** keyword.
4. ==**Importantly**== a friend function is always out of the scope of the class.
5. These functions can be invoked without any class.

```cpp
#include <iostream>

using std::cout, std::cin, std::endl;

class Complex {
	int real, imag;
	
	public:
		// Declaring a friend function
		friend Complex addNum(Complex n1, Complex n2);
		void setNum(int n1, int n2) {
			real = n1;
			imag = n2;
		}
		
		void display(void) {
			cout << real << " + i" << imag << endl;
		}
};

Complex addNum(Complex n1,Complex n2) {
	Complex n3;	
	n3.setNum((n1.real + n2.real) , (n1.imag + n2.imag));
	return n3;
}

int main() {
	Complex c1, c2, c3;
	
	c1.setNum(2, 3);
	c2.setNum(2,7);
	c3 = addNum(c1, c2);
	
	c1.display();
	c2.display();
	c3.display();
	
	return 0;
}
```

### Friend Classes
**Note** : 
1. Forward declaration is when we declare just the function or class header before we even write it entirely.
2. If there a lot of functions which needs to be made friend with a class then a lot of friend declaration is required.
3. To solve this we can make an entire class as a friend to another class.

```cpp
#include <iostream>

using std::cout, std::cin, std::endl;

// Forward Declarations
class Complex;

class Calculator {
	public:
		void sum(int n1 , int n2) {
			cout << n1 + n2<<endl;
		}
		
		void sumComplex(Complex c1, Complex c2);
};

class Complex {
	int real, imag;	
	// For making one function as a friend
	friend void Calculator::sumComplex(Complex c1, Complex c2);
	
	// For making an entire class as a friend
	friend class Calculator;

	public:
		void setData(int n1, int n2){
			real = n1;
			imag = n2;
		}
	
	void display(void) {
		cout << real << " + i" << imag << endl;
	}

};

void Calculator::sumComplex(Complex c1, Complex c2) {
	cout << c1.real+c2.real << " + i" << c1.imag+c2.imag << endl;
}

int main() {
	Complex c1;
	
	c1.setData(1, 3);
	c1.display();
	
	return 0;
}
```
 