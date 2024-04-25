![Source](https://youtu.be/EEJUPXFKe8Q?list=PLu0W_9lII9agpFUAlPFe_VNSlXW5uE0YL)

1. ==Constructors== are used to set values to the class variables while initialising it in order to prevent the addition use of [getters and setters](https://www.w3schools.com/cpp/cpp_encapsulation.asp) which is a concept under [[Encapsulation]].
2. It is automatically invoked as soon as an object of the class has been initialised.
3. They can't return any value.
4. **Note** : parameters can be passed in the constructors like an usual [[Function]].
5. **Note** : The constructors which do ==NOT== accept any parameters are called **default** constructors.
6. ==Importantly== a constructor should always be public in order to be accessed anywhere.
7. We can't refer to their memory addresses.

```cpp
#include <iostream>

using std::cout, std::cin, std::endl;

class Complex {
	int real, imag;
	double real1,imag1;
	
	public:
		// Defining the constructor
		Complex(int x, int y);
		Complex(double x, double y);
		
		// Normal Class Methods
		void display(void);
		void displayDouble(void);

};

// Creating the constructor
// Here 'y' is a default parameter
Complex::Complex(int x, int y=20) {
	real = x;
	imag = y;
}

void Complex::display(void) {
	cout << real << " + " << imag << "i" << endl;
}

// Constructor Overloading
Complex::Complex(double x, double y=20.34) {
	real1 = x;
	imag1 = y;
}

void Complex::displayDouble(void) {
	cout << real1 << " + " << imag1 << "i" << endl;
}

int main() {
	// Explicit Calling
	Complex c1 = Complex(2, 3);
	
	// Implicit Calling
	Complex c2(10.3);
	
	c1.display();
	c2.displayDouble();
	
	return 0;
}
```
