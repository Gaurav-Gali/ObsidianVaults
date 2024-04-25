![Source](https://youtu.be/aKnc1A5NOKo?si=LjZyjw4bIgfubD5-)

**Note**  : 
1. Creating  [[Arrays]] of objects using dynamic memory allocation is not recommended.
2. This might lead in issues which might be difficult to tackle.

```cpp
#include <iostream>

using std::cout , std::cin , std::endl;

class Employee {
	int id;
	static int base_salary;
	int salary;
	
	public:
		void setData(int ID){
			id = ID;
			salary = base_salary * ID;
		}
	
	void getAll(void) {
		cout << "ID : " << id << ", salary : " << salary << endl;
	}

};

// Setting Base Salary
int Employee::base_salary=100;

int main() {

	Employee e1, e2, e3;
	Employee arr[3] = {e1, e2, e3};
	
	// Setting Employee ID
	for (int i = 0; i < 3; i++) {
		arr[i].setData(i + 1);
	}
	
	// Getting Employee's Data	
	for(int i = 0; i < 3; i++) {
		arr[i].getAll();
		cout << "\n";
	}
	
	return 0;
}
```

### Passing Objects as Arguments
```cpp
#include <iostream>

using std::cout , std::cin , std::endl;

class Complex {
	int real;
	int imag;
	
	public:
		void setNum(int r , int i) {
			real = r;
			imag = i;
		}
		
		void setNumBySum(Complex c1 , Complex c2) {
			real = c1.real + c2.real;
			imag = c1.imag + c2.imag;
		}
		
		void display(void) {
			cout << real << " + i" << imag << endl;
		}

};

int main() {
	Complex c1, c2, c3;
		
	c1.setNum(1, 4);
	c2.setNum(2, 2);
	c3.setNumBySum(c1, c2);
	
	c1.display();
	c2.display();
	c3.display();
	
	return 0;
}
```
