![Source](https://youtu.be/WgMPrLX-zsA?t=1820)

**Note** :
1. It follows FIFO principle , "First In First Out".

```cpp
#include <iostream>
#include <queue>

using namespace std;

int main() {
	queue<string> names;
	
	// Pushing elements to the queue
	names.push("Name1");
	names.push("Name2");
	names.push("Name3");
	
	cout << names.front() << endl;
	
	// Popping elements from the queue
	names.pop();
	cout << names.front() << endl;
	
	return 0;
}
```

<hr>
### Priority Queue
**Note** :
1. Max Heap return the maximum value element when popping elements.
2. Min Heap return the minimum value element when popping elements.

```cpp
#include <iostream>
#include <queue>

using namespace std;

int main() {
	// Max Heap
	priority_queue<int> maxh;
	
	maxh.push(0);
	maxh.push(-1);
	maxh.push(4);
	maxh.push(3);
	
	// Min Heap
	priority_queue<int, vector<int>, greater<int>> minh;
	
	minh.push(5);
	minh.push(3);
	minh.push(4);
	minh.push(2);
	
	// Printing the elements of the priority queue
	int n1 = maxh.size();
	int n2 = minh.size();
	
	for(int i = 0; i < n1; i++) {
		cout << maxh.top() << " ";
		maxh.pop();
	}
	cout << endl;
	
	for(int i = 0; i < n2; i++) {
		cout << minh.top() << " ";
		minh.pop();
	}
	
	cout << endl;
	
	return 0;
}
```
