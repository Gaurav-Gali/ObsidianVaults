![Source](https://www.youtube.com/watch?v=jCfR7CFlzts&list=PLu0W_9lII9agpFUAlPFe_VNSlXW5uE0YL&index=14&pp=iAQB)

==Enums== are used to make the code base more readable

```cpp
#include <iostream>

using namespace std;

int main() {
	enum Meal{
		breakfast,
		lunch,
		dinner
	};
	
	Meal m1 = breakfast;
	Meal m2 = lunch;
	Meal m3 = dinner;
	
	cout << m1 << " " << m2 << " " << m3 << endl;
	  
	return 0;
}
```
