![Source](https://www.youtube.com/watch?v=a7dfSBrTZtE&list=PLu0W_9lII9agpFUAlPFe_VNSlXW5uE0YL&index=10&pp=iAQB)

### For loop
```cpp
#include <iostream>

using namespace std;

int main() {
	for(int i=0 ; i<10; i++) {
		cout << i << endl;
	}
	return 0;
}

```

### While loop
```cpp
#include <iostream>

using namespace std;

int main() {
	int count = 0;
	while(count <= 10) {
		cout << count << endl;
		count++;
	}
	return 0;
}
```

### Do While loop
```cpp
#include <iostream>

using namespace std;

int main() {
	int count = 20;
	do {
		cout << count << endl;
	} while(count <= 10);
	return 0;
}
```

### Break and Continue
1. ==Break== statement is used to terminate the iteration.Hence this exits the loop
2. ==Continue== statement is used to skip the current iteration.Hence it will move to the next iteration without breaking the loop.