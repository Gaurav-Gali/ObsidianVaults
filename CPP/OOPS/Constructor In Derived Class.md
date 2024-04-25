![Source](https://youtu.be/gvOO4H7j_qI?list=PLu0W_9lII9agpFUAlPFe_VNSlXW5uE0YL)

**Note :**
1. When we don't create [[Constructors]] in the derived class , then the constructor of the Base class will be called automatically.
2. When there are arguments in the base class's constructor , then the derived class needs to pass arguments to the Base class constructor.
3. If both the Base and derived classes have a constructor then the base class constructor will be executed first.
4. When doing [[Multiple Inheritance]] with a normal base class and a [[Virtual Base Class]] the constructor of the [[Virtual Base Class]] will be executed first.


```cpp
#include <iostream>

using std::cout, std::cin, std::string, std::endl;

class Base1 {
	int data1;
	public:
		Base1(int data) {
			data1 = data;
		}
		
		void show1(void) {
			cout << "Data 1 : " << data1 << endl;
		}

};

class Base2 {
	int data2;
	public:
		Base2(int data) {
			data2 = data;
		}
		
		void show2(void) {
			cout << "Data 2 : " << data2 << endl;
		}

};

class Derived : public Base1 , public Base2 {
	int data3 , data4;
	public:
		Derived(int a , int b , int c , int d) : Base1(a) , Base2(b) {
			data3 = c;
			data4 = d;
		}
		
		void show(void) {
			cout << "Data 3 : " << data3 << endl;
			cout << "Data 4 : " << data4 << endl;
		}

};

int main() {
	Derived d1 = Derived(1, 2, 3, 4);
	
	d1.show();
	d1.show1();
	d1.show2();
	
	return 0;
}
```
