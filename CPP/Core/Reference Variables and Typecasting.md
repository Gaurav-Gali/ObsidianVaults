![Source](https://www.youtube.com/watch?v=a7Wim2t053E&list=PLu0W_9lII9agpFUAlPFe_VNSlXW5uE0YL&index=7&pp=iAQB)

**Note** : This ==::== is the *scope resolution operator*.
And the local variables will override the global variables.
So use ==::== before the variable name to specify that it's referring to the global variable and not the local one.

### Typecasting
1. This is the process of converting a value from one datatype to another.
2. Ex : if you want to convert an int to float
```cpp
#include <iostream>
using namespace std;

int main() {
	int num = 20;
	float num2 = (float) num;
	// alternatively
	float num2 = float(num);
	return 0;
}
```
3. Note that a decimal number by default will be a type of double. Hence to specify that it's a floating point number , use ==f== after the decimal number. Ex : 22.44f