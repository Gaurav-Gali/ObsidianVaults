![Source](https://www.youtube.com/watch?v=WgMPrLX-zsA&t=1354s)

**Note** :
1. In a list , you can't access items directly.
2. In order to access them , you must traverse the list.

```cpp
#include <iostream>
#include <list>

using namespace std;

int main() {
	list<int> l;
	
	// Pushing elements into the list
	l.push_back(10);
	l.push_back(30);
	l.push_front(20);
	
	for(int i:l) {
		cout << i << " ";
	}
	
	cout << endl;
	
	// Popping elements from the list
	l.pop_back();
	
	for(int i:l) {
		cout << i << " ";
	}
	cout << endl;
	
	// Erasing elements from the list
	l.erase(l.begin());
	
	for(int i:l) {
		cout << i << " ";
	}
	cout << endl;
	
	return 0;
}
```
