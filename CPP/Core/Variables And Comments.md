![Source](https://www.youtube.com/watch?v=jigb6W35zHc&list=PLu0W_9lII9agpFUAlPFe_VNSlXW5uE0YL&index=3&pp=iAQB)

```cpp
#include <iostream>

using namespace std;

// This is a global variable
int g_num = 20;

int main() {
	// This is a local variable
	int num = 10;
	bool condition = true;
	
	cout << "Local : " <<num << endl;
	cout << "Global : " << g_num << endl;
	
	// Prints 1 if true and 0 if false
	cout << "Condition : " << condition << endl;
	
	return 0;
}
```
