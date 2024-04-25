![Source](https://youtu.be/DmgQVJXhuLQ?list=PLu0W_9lII9agwh1XjRt242xIpHhPT2llg)

1. __str__
```python
class Employee:
	def __init__(self,name , age):
		self.name = name
		self.age = age
	
	def __str__(self):
		return f"the employee is {self.name} and their age is {self.age}"

class Programmer(Employee):
	def __init__(self,name,age,language):
	super().__init__(name , age)
	self.language = language
	
	def __str__(self):
		return super().__str__()

if __name__ == "__main__":
	p = Programmer("John" , 20 , "Python")
	print(p) # the employee is John and their age is 20
```

2. __repr__
	1. This is used to describe the usability of the class
3. __len__
```python
class Employee:
	def __init__(self,name):
		self.name = name
	
	def __len__(self):
		return len(self.name)

if __name__ == "__main__":
	p = Employee("John")
	print(len(p)) # 4
```

4. __init__
	1. This is used to create the constructor for the class
5. __call__
```python
class Employee:
	def __init__(self,name):
		self.name = name
	
	def __call__(self):
		print(f"{self.name} is a employee")

if __name__ == "__main__":
	p = Employee("John")
	p() # John is a employee
```
**Note** : The __call__ method is provoked when the class's instance is being called as a function