![Source](https://youtu.be/2Y0b9nFA9s8?list=PLu0W_9lII9agpFUAlPFe_VNSlXW5uE0YL)

**Note :**
1. The ==new== keyword is used to dynamically allocate memory to values.
2. The process of getting the value stored in a pointer/memory address is called de-referencing.
3. The ==delete== keyword is used to free the allocated memory.
4. The ==delete[]== keyword is used to free the allocated memory of [[Arrays]].
```cpp
#include <iostream>

using std::cout, std::cin, std::string, std::endl;

int main() {
	// Dynamically creating values with new operator
	int *p = new int(10);
	cout << "The address of p : " << p << endl;
	cout << "The value in address of p : " << *(p) << endl;
	
	// Dynamically creating arrays
	int *arr = new int[5];
	arr[0] = 10;
	arr[1] = 20;
	arr[2] = 30;
	arr[3] = 40;
	arr[4] = 50;
	
	for (int i = 0; i< 5 ; i++) {
		cout << "arr[" << i << "] = " << arr[i] << endl;
	}
	
	// Freeing allocated memory with delete operator
	delete p;
	delete[] arr;
	
	cout << "Freed p : " << *(p) << endl;
	cout << "Freed arr : " << arr[0] << endl;
	
	return 0;
}
```
