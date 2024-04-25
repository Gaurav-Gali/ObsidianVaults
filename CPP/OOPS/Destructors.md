![Source](https://youtu.be/rm4tGxWBkqs?list=PLu0W_9lII9agpFUAlPFe_VNSlXW5uE0YL)

**Note** : 
1. The role of a destructor is to free the memory allocated by a constructor and the objects or even the class itself after it's usage.
2. It neither takes an argument nor does it return any value.
3. As for [[Constructors]] we use the class's name itself , for a destructor we use the ==~== symbol prefixed to the class name.
4. We can create blocks in cpp just by writing the code in plain curly brackets.
```cpp
{
	cout << "In the block" << endl;
	int a = 20;
}
```
5. All the values that were in the block will be destroyed ones its out of the block.
6. ==Importantly== a destructor is invoked automatically when the objects are no longer used.

```cpp
#include <iostream>

using std::cout, std::cin, std::endl;

class Num {
int id;
static int total_objs;

	public:
		Num() {
			total_objs++;
			id = total_objs;
		}
		
		~Num() {
			cout << "Entered Destructor !!!" << endl;
			total_objs--;
		}
		
		void display(void) {
			cout << id << " total : " << total_objs << endl;
		}

};

int Num::total_objs = 0;

int main() {
	Num n1;
	n1.display();
	
	// Initialising a block
	{
		cout << "Entered a block" << endl;
		
		Num n2, n3;	
		n2.display();
		n3.display();
		
		cout << "Leaving a block" << endl;
	
	}
	
	cout << "Back to main" << endl;
	
	return 0;
}
```
