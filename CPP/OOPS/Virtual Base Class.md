![Source](https://youtu.be/eYV-TohBaa0?list=PLu0W_9lII9agpFUAlPFe_VNSlXW5uE0YL)

**Note :**
1. When [[Inheritance]] is performed by 2 derived classes from one class and another class inherits from both the derived classes , an ambiguity error occurs when the last derived class refers a variable from the original base class.
2. In order to prevent this we use a virtual base class.

```cpp
#include <iostream>

using std::cout, std::cin, std::string, std::endl;

class Student {
	protected:
		int roll;
	
	public:
		void setroll(int roll) {
			this->roll = roll;
		}
};

class Acad : virtual public Student {
	public:
		int phy, maths;
		void setmarks(int m1 , int m2) {
			phy = m1;
			maths = m2;
		}
};

class Sports : virtual public Student {
	public:
		int score;
			void setscore(int score) {
			this->score = score;
		}
};

class child : public Acad , public Sports {
	public:
		void display(void) {
			cout << "Roll : " << roll << endl;
			cout << "Score : " << score << endl;
			cout << "Marks : " << phy+maths << endl;
		}

};

int main()
{
	child c1;
	
	c1.setroll(10);
	c1.setmarks(100, 98);
	c1.setscore(140);
	c1.display();
	
	return 0;
}
```
