![Source](https://youtu.be/i3a-G6Ebh9E?list=PLu0W_9lII9agpFUAlPFe_VNSlXW5uE0YL)

### Constants
	They are the variables that cannot be changed upon it's initialisation.
	This is to prevent any form of unwanted data manipulation or data loss.
	
```cpp
#include <iostream>

using namespace std;

int main() {
	const int num = 20;
	return 0;
}
```

### Manipulators
	They are used manipulating or formatting the output

```cpp
#include <iostream>

//For setw()
#include <iomanip>

using namespace std;

int main() {
	//1. endl
	// This allows to leave a line \n in the end of the output
	cout << "hello world" << endl;

	//2.setw()
	// This allows to leave blank space of required length
	cout << "Hello" << setw(4) << "World" << endl;

	return 0;
}
```
