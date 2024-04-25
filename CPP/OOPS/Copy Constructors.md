![Source](https://youtu.be/jhZjyaNO4Wo?list=PLu0W_9lII9agpFUAlPFe_VNSlXW5uE0YL)

**Note** :
1. When creating [[Constructors]] with parameters , when you initialise it , it will throw an error because by default it will call the default constructor.
2. A copy constructor is a type of constructor that makes a copy of another object.
3. That is, we pass the object's reference in order to make the current object look like another.
4. When there is no copy constructor found , then the compiler provides with it's own.

```cpp
#include <iostream>

using std::cout, std::cin, std::endl;

class Num {
	int number;
	
	public:
		// Default constructor
		Num() {
			number = 0;
		}
		
		Num(int num) {
			number = num;
		}
		
		// Class methods
		void display(void) {
			cout << "Number: " << number << endl;
		}
		
		// Copy constructor
		Num(Num &obj) {
			number = obj.number;
		}

};

int main() {
	Num n1,n2,n3;
	
	n1 = Num(1);
	n2 = Num(n1);
	n3 = Num(20);
	
	n1.display();
	n2.display();
	n3.display();
	
	return 0;
}
```

**Another way of invoking copy constructors**
```cpp
Num n1(10);
// This works as n2 is initialised and invoked in the same line.
Number n2 = n1;
```
