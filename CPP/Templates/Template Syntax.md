![Source](https://youtu.be/kKJeekDKU30?list=PLu0W_9lII9agpFUAlPFe_VNSlXW5uE0YL)

**Note :**
1. Templates in cpp are used to follow the D.R.Y principle or the don't repeat yourself principle.
2. Templates are used to do ==generic== programming.
3. And , templates are also known as parameterised classes.

```cpp
#include <iostream>

using std::cout, std::cin, std::string, std::endl;

template <class T>
class Vector {
	public:
		T *arr; // {i,j,k}
		int size;
		
		Vector(int size) {
			this->size = size;
			arr = new T[size];
		}
		
		T dot(Vector v1) {
			T dot_product = 0;
			
			for(int i = 0; i < size; i++) {
				dot_product += this->arr[i] * v1.arr[i];
			}
					
			return dot_product;
		}
};

int main() {
	Vector <double>v(3);
	v.arr[0] = 10;
	v.arr[1] = 20;
	v.arr[2] = 30;
	
	Vector <double>v1(3);
	v1.arr[0] = 1.6;
	v1.arr[1] = 2;
	v1.arr[2] = 3;
	
	cout << v.dot(v1) << endl;
	
	return 0;
}
```
