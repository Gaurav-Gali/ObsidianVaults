![Source](https://www.youtube.com/watch?v=jCfR7CFlzts&list=PLu0W_9lII9agpFUAlPFe_VNSlXW5uE0YL&index=14&pp=iAQB)

==Union== shares memory for all it's types.
```cpp
#include <iostream>

using namespace std;

typedef union money {
	int rupees;
	double dollars;
}money;

int main() {
	money m1;
	m1.rupees = 80;
	m1.dollars = 1;
	cout << m1.dollars << endl;

	return 0;
}
```