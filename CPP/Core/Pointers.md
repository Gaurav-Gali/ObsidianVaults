![Source](https://www.youtube.com/watch?v=EvYmTCx9BFs&list=PLu0W_9lII9agpFUAlPFe_VNSlXW5uE0YL&index=12&pp=iAQB)

==Pointers== store the memory address of an another variable.

```cpp
#include <iostream>

using namespace std;

int main() {
	int a = 10;
	int* b = &a;

	cout << "Value of a : " << a << endl;
	cout << "Address of a : " << b << endl;

	return 0;
}
```
