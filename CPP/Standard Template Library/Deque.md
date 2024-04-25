![Source](https://www.youtube.com/watch?v=WgMPrLX-zsA&t=1064s)

**Note** :
1. In a deque , the push and pop operations can be done from both the ends of the data structure.

```cpp
#include <iostream>
#include <deque>

using namespace std;

int main() {
	deque<int> dec;
	
	// Pushing elements in the front and back
	dec.push_back(10);
	dec.push_back(20);
	dec.push_back(30);
	dec.push_front(-10);
	
	for(int i : dec) {
		cout << i << " ";
	}
	
	cout << endl;
	
	// Popping elements in the front and back
	dec.pop_back();
	dec.pop_front();
	
	for(int i : dec) {
		cout << i << " ";
	}
	
	cout << endl;
	
	// Checking if a deque is empty
	cout << dec.empty() << endl;
	
	// Erasing elements of a deque
	dec.erase(dec.begin(), dec.begin() + 1);
	for(int i : dec) {
		cout << i << " ";
	}
	
	return 0;
}
```
