![Source](https://youtu.be/QcLI2zGVYFo?list=PLu0W_9lII9agpFUAlPFe_VNSlXW5uE0YL)

==Static== data members value are set constant for all the objects of the class.
**Note :** 
1. In C++ a static member is defined by the ==static== keyword while initialising it
2. It should also be mentioned outside the class because the memory for it will be allocated with the memory allocated for the class and won't be allocated every time the object is created.
3. And a static variable will be by default get initialised with 0.
4. In OOPS a static variable is also called a **class variable**.
5. If you want to set the initial value of a static variable then do it outside the function.
6. When creating static functions they don't have to be mentioned outside the class.
7. And the static functions don't have access to the non-static data.

```cpp
#include <iostream>

using std::cout , std::endl;

class Employee {
	static int total_employees;
	int id;
	
	public:
		void getID(void) {
			id = ++total_employees;
			cout << id << endl;
		}

};

int Employee::total_employees; // set to 0 by default

int main() {
	Employee e1,e2,e3;
	
	e2.getID();
	e1.getID();
	e3.getID();
	
	return 0;
}
```
