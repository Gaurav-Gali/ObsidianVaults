![Source](https://www.youtube.com/watch?v=ePJxpxsnkGw&list=PLu0W_9lII9agpFUAlPFe_VNSlXW5uE0YL&index=13&pp=iAQB)

==Arrays== store a collection of values of the same datatype.
Note : The array by itself is [[pointers]] pointing to the memory address of the first element.

```cpp
#include <iostream>

using namespace std;

int main() {
	//1. Directly storing all the elements
	int arr[] = {1,2,3,4,5};

	//2. Storing the values after initialising the array
	int arr2[2];
	arr2[0] = 1;
	arr2[1] = 2;

	//3. Iterating an array
	for(int i=0; i<2 ; i++) {
		cout << arr2[i] << endl;
	}

	return 0;
}
```