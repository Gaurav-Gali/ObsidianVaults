![Source](https://youtu.be/fB3JHNnlRfI?list=PLu0W_9lII9agpFUAlPFe_VNSlXW5uE0YL)

**Note :**
1. Previously in [[Pointers to Derived Class]] we saw that when creating a pointer to a base class and referencing the derived class in it , when we call a function display() which is present in both the base and the derived class , it always called the function in the base class.
2. What if we want to call the function in the derived class?
3. That why we use the virtual functions. Simply add the virtual keyword before the function's name in the base class and then it will act as an virtual function.
4. So when called from the pointer , the function from the derived class will be called.
5. Basically it override's the default behaviour.
6. Rules:
	1. They cannot be [[Static Data Members]]
	2. They are accessed by object pointers
	3. Virtual functions can be friend a class
	4. If a virtual function is declared in a base class , it need not be used
	5. If a virtual function is defined in the base class , then they is no need to define it in the derived class

```cpp
#include <iostream>

using std::cout, std::cin, std::string, std::endl;

class Base {
	int a = 10;
	public:
		virtual void display(void) {
			cout << "Base a : " << a << endl;
		}
};

class Derived : public Base {
	int a = 20;
	
	public:
		void display(void) {
			cout << "Derived a : " << a << endl;
		}
};

int main() {
	// Creating a pointer of a base class
	Base *base_ptr;
	
	// Creating objects of base and derived classes
	Base base_obj;
	Derived derived_obj;
	
	// Referencing derived object to base pointer
	base_ptr = &derived_obj;
	base_ptr->display();
	
	return 0;
}
```
