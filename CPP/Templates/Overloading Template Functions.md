![Source](https://youtu.be/Y_RMNcXAM1U?list=PLu0W_9lII9agpFUAlPFe_VNSlXW5uE0YL)

**Note :**
1. When creating a function with template and performing overloading with another non-template function , then the exact match would take the highest priority.

```cpp
#include <iostream>

using std::cout, std::cin, std::string, std::endl;

void func(int a) {
	cout << "func non-template : " << a << endl;
}

template <class T>
void func(T a) {
	cout << "func template : " << a << endl;
}

int main() {
	func(10); // Non-template will be called
	func(10.5); // Template function will be called
	return 0;
}
```
