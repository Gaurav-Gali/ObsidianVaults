![Source](https://youtu.be/cEOfK_L4gGA?list=PLu0W_9lII9agpFUAlPFe_VNSlXW5uE0YL)

**Note :**
1. this [[Pointers]] in cpp is used to refer to the variable of the current class.
2. this pointer can also be used to return the reference of the current class from a member function.

```cpp
#include <iostream>

using std::cout, std::cin, std::string, std::endl;

class A {
	int a;
	public:
		A& setData(int a) {
			this->a = a;
			return *this;
		}
		
		void getData(void) {
			cout << "a: " << a << endl;
		}

};

int main() {
	A a;
	a.setData(10).getData();
	return 0;
}
```
