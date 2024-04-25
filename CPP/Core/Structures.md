![Source](https://www.youtube.com/watch?v=jCfR7CFlzts&list=PLu0W_9lII9agpFUAlPFe_VNSlXW5uE0YL&index=14&pp=iAQB)

==Structs== store multiple datatypes.Hence it is it's own datatype


```cpp
#include <iostream>

using namespace std;

typedef struct Employee {
	int age;
	double salary;
} emp;

int main() {
	//1. Implementing structs normally
	struct Employee e1 = {
		26,
		10_00_000
	};
	cout << e1.age << e1.salary << endl;

	//2. Impelenting structs with typedef
	emp e2 = {
		30,
		20_00_000
	};
	cout << e2.age << e2.salary << endl;
	return 0;
}
```
