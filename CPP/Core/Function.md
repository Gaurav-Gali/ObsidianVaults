![Source](https://www.youtube.com/watch?v=RFLFX1boGwo&list=PLu0W_9lII9agpFUAlPFe_VNSlXW5uE0YL&index=15&pp=iAQB)

```cpp
#include <iostream>

using namespace std;

// function prototypes
int sum(int a, int b);
void swap(int* a , int*b);
inline int square(int a, int times);

// main function
int main() {
	cout << sum(10,20) << endl;
	int a,b;
	
	a = 10;
	b = 20;
	
	cout << a << "\t" << b << endl;
	swap(a,b);
	cout << a << "\t" << b << endl;
	
	// calling a inline function
	cout << square(10,2) << endl;
	cout << square(10,2) << endl;
	cout << square(10,2) << endl;
	
	return 0;
}

int sum(int a, int b) {
	return a + b;
}

// call by reference
void swap(int* a , int*b) {
	int* c = a;
	a = b;
	b = c;
}

// inline function
inline int square(int a , int times=1) {
	return a*a*times;
}

```

Note : 
1. Formal parameter : They are declared in the function header
2. Actual parameter : They are used in the function call
3. Call by value : When direct values are used as the function parameter
4. Call by reference : When the memory address is used as function parameter
5. Inline functions : These function replace the code in them with the function call in compile time. 
6. Default arguments : When you pre-assign the value of a function parameter
```cpp
int add(int x , int y=20);
```
7. Constant arguments : When the value of the parameter is not meant to be modified.
```cpp
int strlen(const char s);
```
8. Function overriding : When 2 or more function are created with the same name but with different arguments is called function overriding.
```cpp
int func(int a);
int func(double a);
int func(bool a);
```
