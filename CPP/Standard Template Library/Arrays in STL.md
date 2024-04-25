![Source](https://www.youtube.com/watch?v=WgMPrLX-zsA&t=297s)

**Note** :
1. Like the ordinary cpp array the STL array is also static
2. But it provides with a lot of functions that are useful

```cpp
#include <iostream>
#include <array>

using namespace std;

int main() {
	array<int, 5> arr = {1, 2, 3, 4, 5};
	
	// Getting the size of the array
	int size = arr.size();
	
	// Iterating the array
	for(int i = 0; i < size; i++) {
		cout << arr[i] << endl;
	}
	
	// Using the at() function
	cout << arr.at(0) << endl;
	
	// To check if the array is empty
	cout << arr.empty() << endl;
	
	// To get the first and last elements
	cout << arr.front() << endl; // first
	cout << arr.back() << endl; // last
	
	return 0;
}
```
