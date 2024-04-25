![Source](https://youtu.be/ANpUQgyRPKk?list=PLu0W_9lII9agpFUAlPFe_VNSlXW5uE0YL)

**Note :**
1. Using the arrow operator ==->== we can directly de reference the value of  class [[Pointers]].

```cpp
#include <iostream>

using std::cout, std::cin, std::string, std::endl;

class Complex {
	int real , imag ;
	public:
		void setData(int real , int imag) {
			this->real = real;
			this->imag = imag;
		}
		
		void getData(void) {
			cout << real << " + " << imag <<"i"<< endl;
		}
};

int main() {
	Complex* c = new Complex;
	(*c).setData(2, 3);
	(*c).getData();
	
	// Using arrow operator
	Complex* c2 = new Complex;
	c2->setData(4, 5);
	c2->getData();
	
	return 0;
}
```
