![Source](https://youtu.be/RBAWWutf0fY?list=PLu0W_9lII9agpFUAlPFe_VNSlXW5uE0YL)

**Note :**
1. When a base class has [[Virtual Functions]] that is made just so that it has to be overridden , such classes are called abstract base class.
2. The virtual functions created must be at any cost be overridden in order for the class to be called an abstract base class.
3. **IMPORTANT : ** when you equate a function to 0 , then it will be called a "do nothing function" or "pure virtual functions".
```cpp
virtual void display(void) = 0;
```

```cpp
#include <iostream>

using std::cout, std::cin, std::string, std::endl;

class Base {
	int var;
	public:
		// Pure virtual function or do nothing function
		virtual void display(void) = 0;
};

class Derived : public Base {
	int var = 20;
	public:
		void display(void) {
			cout << "Derived var : " << var << endl;
		}
};

int main() {
	// Creating a pointer for the base class
	Base *base_ptr;
	
	// Creating objects for the derived class
	Derived d1;
	
	// Referencing
	base_ptr = &d1;
	base_ptr->display();
	
	return 0;
}
```
