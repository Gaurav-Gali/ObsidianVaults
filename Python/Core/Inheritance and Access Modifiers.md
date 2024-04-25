![Source](https://youtu.be/-KsfUaQEY9Y?list=PLu0W_9lII9agwh1XjRt242xIpHhPT2llg)

```python
class Employee:
	_name = ""
	
	def __init__(self, name: str):
		self._name = name
	
	def showDetails(self):
		print(f"{self._name} is an employee")

  
class Programmer(Employee):
	def showLang(self):
		print(f"{self._name} can also code")


if __name__ == "__main__":
	p1 = Programmer("John")
	p1.showDetails() # John is an employee
	p1.showLang() # John can also code 
```

<hr>

**Types of Access**
1. Public
	1. Can be accessed from anywhere in the code
	2. By default all class variables are public
2. Private
	1. Cannot be accessed outside it's class
	2. Except in it's children class
```python
class Person:
	def __init__(self):
	self.__name = "John"

if __name__ == "__main__":
	p1 = Person()
	print(p1.__name) # Error
```
**Note** : Private variables are prefixed with (double underscore)
1. Protected


**Name Mangling**
Even though a variable is set to be private , it can still be accessed by this process
```python
class Person:
	def __init__(self):
	self.__name = "John"

if __name__ == "__main__":
	p1 = Person()
	print(p1._Person___name) # John
```