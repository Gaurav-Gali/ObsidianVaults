![Source](https://www.youtube.com/watch?v=WgMPrLX-zsA&t=507s)

**Note** :
1. Vectors are used to create dynamic arrays in cpp.
2. When a vector is first initialised , it's default capacity would be 0.
3. When you want to add a new element and if the vector is already full , then the vector doubles it's capacity.
4. Clearing the vector will make the size as 0 but the capacity will remain the same.

```cpp
#include <iostream>
#include <vector>

using namespace std;

int main() {
	vector<int> vec;
	
	// Getting the capacity of a vector
	int size = vec.capacity();
	cout << size << endl;
	
	// Adding elements to the vector
	vec.push_back(1);
	vec.push_back(2);
	cout << vec.at(0) << endl;
	
	// Popping elements from the vector
	vec.pop_back();
	
	// To get the first and last elements
	cout << vec.front() << endl; // first
	cout << vec.back() << endl; // last
	
	// Clearing the vector
	vec.clear();
	cout << endl;
	
	// Initializing the vector with default sie and elements
	vector<int> vec1(5, 1);
	for(int i:vec1) {
		cout << i;
	}
	
	cout << endl;
	
	// Passing the elements of one vector to another vector
	vector<int> vec2(vec1);
	for(int i:vec2) {
		cout << i;
	}
	
	return 0;
}
```
