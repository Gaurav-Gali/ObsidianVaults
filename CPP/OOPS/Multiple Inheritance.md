![Source](https://youtu.be/h3INeRqf2vU?list=PLu0W_9lII9agpFUAlPFe_VNSlXW5uE0YL)

**Note** : 
1. Multiple [[Inheritance]] is when a derived class has more than one base class.

```cpp
#include <iostream>

using std::cout, std::cin, std::string, std::endl;

class BaseOne {
	protected:
		int n1;
	
	public:
		void setn1(int n1) {
			this->n1 = n1;
		}

};

class BaseTwo {
	protected:
		int n2;
	
	public:
		void setn2(int n2) {
			this->n2 = n2;
		}
};

class Derived : public BaseOne , public BaseTwo {
	public:
		void display(void)
		{
			cout << n1 << " " << n2 << endl;
			cout << n1+n2 << endl;
		}
};

int main() {
	Derived d1;
	
	d1.setn1(10);
	d1.setn2(20);
	d1.display();
	
	return 0;
}
```