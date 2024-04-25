![Source](https://www.youtube.com/watch?v=JRKs3s15Kjc&list=PLu0W_9lII9agpFUAlPFe_VNSlXW5uE0YL&index=18&pp=iAQB)

==Recursion== is when a [[function]] is called within the same function.

```cpp
#include <iostream>

using namespace std;

int factorial(int n) {
	if(n <= 1) {
		return 1;
	}
	
	return n*factorial(n-1);
}

int main() {
	cout << factorial(3) << endl;
	return 0;
}
```
