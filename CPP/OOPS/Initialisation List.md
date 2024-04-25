![Source](https://youtu.be/-Re7K7mHtv4?list=PLu0W_9lII9agpFUAlPFe_VNSlXW5uE0YL)

**Note** :
1. When [[Constructors]] get pretty complex we can use the initialisation list to easily assign values to the class variables.

```cpp
#include <iostream>

using std::cout, std::cin, std::string, std::endl;

class A 
	int num , num1;
	public:
		A(int x , int y) : num(x), num1(y) {
		cout << "num : " << num << endl;
		cout << "num1 : " << num1 << endl;
	}
};

int main() {
	A(10, 20);
	
	return 0;
}
```
