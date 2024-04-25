![Source](https://www.youtube.com/watch?v=AY96XFqb934&list=PLu0W_9lII9agpFUAlPFe_VNSlXW5uE0YL&index=9&pp=iAQB)

### Basic If else
```cpp
#include <iostream>

using namespace std;

int main() {
	int age = 18;

	if(age >= 18) {
		cout << "Can Vote !" << endl;
	} else {
		cout << "Can't Vote" << endl;
	}

	return 0;
}
```

### If else if ladder
```cpp
#include <iostream>

using namespace std;

int main() {
	int age = 18;

	if(age >= 18) {
		cout << "Can Vote !" << endl;
	}else if(age > 0 && age < 18) {
		cout << "Can't Vote" << endl;
	}else {
		cout << "Not a valid age !!" << endl;
	}

	return 0;
}
```

### Switch case
```cpp
#include <iostream>

using namespace std;

int main() {
	int day = 1;
	switch(day) {
		case 1:
			cout << "Mon" << endl;
			break;
		case 2:
			cout << "Tue" << endl;
			break;
		case 3:
			cout << "Wed" << endl;
			break;
		case 4:
			cout << "Thur" << endl;
			break;
		case 5:
			cout << "Fri" << endl;
			break;
		case 6:
		case 7:
			cout << "Weekend" << endl;
			break;
		default :
			cout << "Not a valid day !!" << endl;
	}
	return 0;
}
```