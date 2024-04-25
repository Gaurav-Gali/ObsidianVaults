![Source](https://youtu.be/8SQL9-cQmsw?list=PLu0W_9lII9agpFUAlPFe_VNSlXW5uE0YL)

**Note :**
1. Simply add multiple class in the template separate with a comma to create a multi parameter template.

```cpp
#include <iostream>

using std::cout, std::cin, std::string, std::endl;

template <class T1, class T2> // multiple parameters
class MyClass {
	T1 data1;
	T2 data2;
	public:
		MyClass(T1 data1, T2 data2) {
			this->data1 = data1;
			this->data2 = data2;
		}
		
		void display(void) {
			cout << "Data1 : " << data1 << endl;
			cout << "Data2 : " << data2 << endl;
		}
};

int main() {
	MyClass<int, double> c1(10, 20.4);
	c1.display();
	return 0;
}
```

2. In order to give default parameters:
```cpp
template <class T1 , class T2=int>
class A{};

int main() {
	A<> a1;
}
```
