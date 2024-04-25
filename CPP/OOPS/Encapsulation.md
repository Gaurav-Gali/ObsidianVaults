![Source](https://youtu.be/tL8vnfFFzVQ?list=PLu0W_9lII9agpFUAlPFe_VNSlXW5uE0YL)

==Encapsulation== is the process of limiting access to the data of a class like variables , methods etc.

### Some important keywords
1. public : 
	1. This allows to set the data within to be public , that is it can be accessed from anywhere.
```cpp
public:
	int d, e, f;

	int add();
	
	void setData(int a1, int b1, int c1) {
		a = a1;
		b = b1;
		c = c1;
	}
```
2. private :
	1. This allows to set the data within to be private , that is it can be accessed only from within the class.
```cpp
private:
	int a, b, c;
```

Note : If a method's header is declared within a class , it can be later modified with the help of the ==::==  **scope resolution operator**
```cpp
class Calc {
	public:
		int square(int x);
};

int Calc :: square(int x) {
	return x*x;
}
```

### Getters And Setters
Since some variables in a class are set to private , we need another way of manipulation these values. Hence we create special functions to do the same and such functions are called **getters** which retrieve the data and **setters** that store the data.
```cpp
class Employee {
	private:
		int ID , double SALARY;
	public:
		void setData(int id , doubel salary) {
			ID = id;
			SALARY = salary;
		}

		void getData() {
			cout << "ID = " << ID << endl;
			cout << "SALARY = " << SALARY << endl;
		}
}

int main() {
	Employee e1;
	//1. This will store the required data
	e1.setData(1,20000);
	//2. This will print the required data
	e1.getData();
}

```
