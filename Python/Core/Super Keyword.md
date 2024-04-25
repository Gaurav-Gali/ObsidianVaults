![Source](https://youtu.be/P648reefNh0?list=PLu0W_9lII9agwh1XjRt242xIpHhPT2llg)

The **super()** method is used to call a method from one of it's parent classes

```python
class Parent:
	def parent_method(self):
		print("parent method was called")
	
class Child(Parent):
	def call_parent_method(self):
		print("Inside Child method")
		super().parent_method()

if __name__ == "__main__":
	c = Child()
	c.call_parent_method()
```

Useful example for the super() method
```python
class Employee:
	def __init__(self,name , age):
		self.name = name
		self.age = age

class Programmer(Employee):
	def __init__(self,name,age,language):
		super().__init__(name , age)
		self.language = language

if __name__ == "__main__":
	p = Programmer("John" , 20 , "Python")
```
Here instead of typing self.name and self.age again we could make use of the __init__ method of the parent class to get those values.