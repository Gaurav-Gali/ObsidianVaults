![Source](https://youtu.be/YTS0ShpFsrM?list=PLu0W_9lII9agpFUAlPFe_VNSlXW5uE0YL)

```cpp
#include <iostream>

using std::cout, std::cin, std::string, std::endl;

template <class T1 , class T2>
T1 add(T1 a , T2 b) {
	return a + b;
}

template <class T>
void swap(T &a , T &b) {
	T temp = a;
	a = b;
	b = temp;
}

int main() {
	cout << add(1.45, 2.2) << endl;
	
	int a, b;
	a = 20;
	b = 30;
	cout << a << b << endl;
	
	swap(a, b);
	cout << a << b << endl;
	
	return 0;
}
```
