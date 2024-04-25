![Source](https://youtu.be/WgMPrLX-zsA?t=1615)

**Note** :
1. It follows LIFO principle , "Last In First Out".

```cpp
#include <iostream>
#include <stack>

using namespace std;

int main() {
	stack<string> names;
	names.push("Name1");
	names.push("Name2");
	names.push("Name3");
	
	// Getting the top element
	cout << names.top() << endl;
	
	// Popping elements from the stack
	names.pop();
	cout << names.top() << endl;
	
	// Getting the size of the stack
	cout << names.size() << endl;
	
	// Checking if the stack is empty
	cout << names.empty() << endl;
	
	return 0;
}
```
