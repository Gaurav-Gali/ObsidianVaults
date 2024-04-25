![Source](https://youtu.be/BLb6-ZgxqHg?list=PLu0W_9lII9agpFUAlPFe_VNSlXW5uE0YL)

**Note** : 
1. When a class is derived from another class which has been derived from another base class is known as Multilevel [[Inheritance]]
2. If inheritance happens like this ==A --> B --> C== then A is the base class for B and B is the base class for C.
3. B here is also called the intermediate base class.

```cpp
#include <iostream>

using std::cout, std::cin, std::string, std::endl;

class Student {
	protected:
		int roll;
	
	public:
		void set_roll(int roll) {
			this->roll = roll;
		}
		
		void get_roll(void) {
			cout << roll << endl;
		}

};

typedef struct Marks {
	int maths;
	int physics;
} marks;

class Exam : public Student {
	protected:
		marks m1;
	
	public:
		void set_marks(int math , int phy) {
			m1.maths = math;
			m1.physics = phy;
		}
		
		void get_marks(void) {
			cout << m1.maths << endl;
			cout << m1.physics << endl;
		}

};

class Result : public Exam {
	public:
		void display(void) {
			get_roll();
			get_marks();
			cout << "Avg : " << (m1.maths+m1.physics)/2 << endl;
		}
};

int main() {
	Result r1;
	
	r1.set_roll(1404);
	r1.set_marks(100, 98);
	r1.display();
	
	return 0;
}
```
