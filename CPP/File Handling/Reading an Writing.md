![Source](https://youtu.be/LS1zjr1wog4?list=PLu0W_9lII9agpFUAlPFe_VNSlXW5uE0YL)

**Note :**
1. File handling operations can be performed using ==\#include \<fstream>== library.
2. Useful classes for working with files are :
	1. fstreambase
	2. ifstream : is from std
	3. ofstream : is from std
3. There are primarily 2 ways of opening a file in cpp:
	1. Using [[Constructors]]
	2. Using the member function ==open()==
4. Note , when the compiler encounters a space or a end-line then it would stop reading the file from there.
5. To tackle this , we have the getline(obj , str) function.
6. **IMPORTANT** always close the files after you open them using the close function.
7. ==obj.close()==
8. use the ==eof()== function to find if compiler has come to the end of the file.
9. Note that if eof() == 0 then the eof has not been reached otherwise it has been reached.

```cpp
#include <iostream>
#include <fstream>

using std::cout, std::cin, std::string, std::endl , std::ofstream , std::ifstream;

int main() {
	// Opening and writing to files
	string st1 = "Line 1 \nLine 2 \nLine 3 \nLine 4";
	ofstream out("0_test_file.txt");
	out << st1;
	out.close();
	
	// Reading from files
	string st2;
	ifstream in("0_test_file.txt");
	
	// using the getline function
	getline(in, st2);
	cout << st2 << endl;
	in.close();
	
	return 0;
}
```

```cpp
#include <iostream>
#include <fstream>

using std::cout, std::cin, std::string, std::endl , std::ofstream , std::ifstream;

int main() {
	// Using the open function to read
	ifstream in;
	in.open("0_test_file.txt");
	
	while(in.eof() == 0) {
		string s1;
		getline(in , s1);
		cout << s1 << endl;
	}
	
	in.close();
	
	return 0;
}
```
