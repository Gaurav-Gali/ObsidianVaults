![Source](https://youtu.be/WgMPrLX-zsA?t=2722)

**Note**:
1. Data is stored as key-value pair.
2. One key will point to only only value.

```cpp
#include <iostream>
#include <map>

using namespace std;

int main() {
	map<int, string> names;
	
	// Adding the key value pairs
	names[0] = "Name1";
	names[1] = "Name2";
	names[2] = "Name3";
	names[3] = "Name4";
	
	for(auto i : names) {
		// Printing the keys
		cout << i.first << " ";
		
		// Printing the values
		cout << i.second << endl;
	}
	
	// Seeing if the key is present in the map
	cout << names.count(2) << endl;
	
	// Erasing a key value pair from the map
	names.erase(2);
	for(auto i : names) {
		// Printing the keys
		cout << i.first << " ";
	
		// Printing the values
		cout << i.second << endl;
	}
	
	// Finding values in the map
	map<int, string>::iterator itr = names.find(1);
	
	for (auto i = itr; i != names.end(); ++i) {
		cout << (*i).first << " ";
	}
	
	return 0;
}
```
