![Source](https://youtu.be/RO1ZYW9NAzg?list=PLu0W_9lII9agpFUAlPFe_VNSlXW5uE0YL)

**Note** :
1. Since reusability is a important feature of OOPS , we use inheritance to leverage that.
2. We can inherit from classes and add more features to them.
3. Inheritance basically is when we create another class and extend some features from it's parent class.
4. Here the existing class is called the ==Base Class== and the inherited class is called the ==Derived== class.
5. A derived class has more than one base class.
6. Visibility mode : It reflects how the inheritance happens (public / private).
	7. If ==public== : then public data of the base class will be public data of the derived class. 
	8. If ==private== : then public data of the base class will be private data of the derived class.
	9. By default the visibility is private.
	10. Private data of the base class is never inherited.
	11. If you want to access that data then use getters and setters.
	12. Or you can use the ==protected== access modifier.
		1. This access modifier keeps the variables in them private but allows it to be accessed in it'd derived class.
7. There are different types of Inheritance : 
	1. Single level Inheritance / Normal Inheritance
	2. [[Multiple Inheritance]]
	3. Hierarchical Inheritance
	4. [[Multilevel Inheritance]]
	5. Hybrid Inheritance / Multi Path Inheritance

**==Important==**
<table border=2>
	<tr>
		<td>Access Modifier Type</td>
		<td>Public Derivation</td>
		<td>Private Derivation</td>
		<td>Protected Derivation</td>
	</tr>
	<tr>
		<td>Private Members</td>
		<td>Not Inherited</td>
		<td>Not Inherited</td>
		<td>Not Inherited</td>
	</tr>
	<tr>
		<td>Protected Members</td>
		<td>Protected</td>
		<td>Private</td>
		<td>Protected</td>
	</tr>
	<tr>
		<td>Public Members</td>
		<td>Public</td>
		<td>Private</td>
		<td>Protected</td>
	</tr>
</table>

```cpp
#include <iostream>

using std::cout, std::cin, std::string, std::endl;

class Employee {
	protected:
		int id;
		double salary;
	public:
		Employee(){}
		
		Employee(int id, double salary) {
			this->id = id;
			this->salary = salary;
		}
		
		void display(void) {
			cout << id << " : " << salary << endl;
		}

};

class Programmer : public Employee {
	string language;
	public:
		Programmer(int Id , double Salary, string language) {
			id = Id;
			salary = Salary;
			this->language = language;
		}
		
		void display(void) {
			cout << id << " : " << salary << " : " << language << endl;
		}

};

int main() {
	Employee e1(1, 25.5);
	e1.display();
	
	Programmer p1(2, 30, "cpp");
	p1.display();
	
	return 0;
}
```
