![Source](https://youtu.be/WgMPrLX-zsA?t=2309)

**Note** :
1. A set stores only unique elements.
2. So if you push the same element 10 times , still only one occurrence of that element will be stored.
3. While popping elements from a set , it comes in the sorted order.
4. If you pop from an unordered set , then the elements don't get popped in sorted order.
5. The find function will return an iterator.

```cpp
#include <iostream>
#include <set>

using namespace std;

int main() {
	set<int> s;
	
	s.insert(4);
	s.insert(1);
	s.insert(2);
	s.insert(1);
	s.insert(3);
	
	for(int i : s) {
		cout << i << " ";
	}
	
	cout << endl;
	
	// Count function
	cout << s.count(2) << endl;
	
	// Find function
	set<int>::iterator itr = s.find(2);
	
	for (auto i = itr; i != s.end(); i++) {
		cout << *i << " ";
	}
	
	return 0;
}
```
