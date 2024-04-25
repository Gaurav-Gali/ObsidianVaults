![Source](https://youtu.be/0YQ_yhX46uk?list=PLu0W_9lII9agpFUAlPFe_VNSlXW5uE0YL)

**Note :**
1. When creating  [[Pointers in OOPS]]  , and if we create them for a base class and then store a reference of the derived class in it , it would access the function of the base class only.

```cpp
#include <iostream>

using std::cout, std::cin, std::string, std::endl;

class Base {
	public:
		int base_var;
		void display(void) {
			cout << "Base : " << base_var << endl;
		}

};

class Derived : public Base {
	public:
		int derived_var;
		void display(void) {
			cout << "Derived : " << derived_var << endl;
		}
};

int main() {
	// Creating a pointer to the base class
	Base *base_pointer;
	
	// Creating a pointer to the derived class
	Derived *derived_pointer;
	
	// Creating objects of the base and derived classes
	Base base_obj;
	Derived derived_obj;
	
	// Referencing derived class object to the base class pointer
	base_pointer = &derived_obj;
	base_pointer->base_var = 10;
	
	// base_pointer->derived_var = 10; is not allowed
	base_pointer->display();
	
	// Referencing base class object to the derived class pointer
	derived_pointer = &derived_obj;
	derived_pointer->derived_var = 20;
	derived_pointer->base_var = 30; // Here we can change the base variable
	derived_pointer->display();
	
	return 0;
}
```
