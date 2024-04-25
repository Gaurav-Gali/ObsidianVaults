![Source](https://youtu.be/ZqfArYoV9Lg?list=PLu0W_9lII9agpFUAlPFe_VNSlXW5uE0YL)

**Note :**
1. When performing [[Multiple Inheritance]] , when both the Base classes have a same function name  , then when calling it through the derived class will create an ambiguity error since the compiler won't know which one to call.
2. When declared a function in the Derived class with the same name as the Base class , then the ambiguity will be automatically resolved by calling the one in the derived class.

```cpp
#include <iostream>

using std::cout, std::cin, std::string, std::endl;

class Base1 {
	public:
		void greet(void){
			cout << "Hello Base1" << endl;
		}
};

class Base2 {
	public:
		void greet(void){
			cout << "Hello Base2" << endl;
		}
		
		void New(void) {
			cout << "Hello New" << endl;
		}
};

class DerivedClass : public Base1 , public Base2 {
	public:
		void greet(void){
			Base1::greet();
		}
		
		void New(void) {
			cout << "Hello New in derived class" << endl;
		}
};

int main() {
	DerivedClass d1;
	d1.greet();
	d1.New();
	
	return 0;
}
```
