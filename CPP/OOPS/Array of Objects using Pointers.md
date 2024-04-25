![Source](https://youtu.be/OCmCyYxSi2I?list=PLu0W_9lII9agpFUAlPFe_VNSlXW5uE0YL)

**Note :**
1. We can create [[Array Of Objects]] using pointers just like how we declare normal [[Pointers in OOPS]] using the new operator.
2. when given the size of memory to be allocated as say 'n' then the compiler will allocate 'n' memory block of each with size of the class.

```cpp
#include <iostream>

using std::cout, std::cin, std::string, std::endl;

class ShopItem {
	int id, price;
	public:
		void setData(int id, int price) {
			this->id = id;
			this->price = price;
		}
		
		void getData(void) {
			cout << "ID : " << id;
			cout << ", Price : " << price << endl;
		}
};

int main() {
	ShopItem *items = new ShopItem[3];
	
	// Setting values for each item
	for (int i = 0; i < 3; i++) {
		int p, q;
		cout << i + 1 << ". Enter ID and Price : ";
		cin >> p >> q;
		(items + i)->setData(p,q);
	}
	
	cout << endl;
	
	// Getting values from each item
	for (int i = 0; i < 3; i++) {
		(items + i)->getData();
	}
	
	return 0;
}
```
