![Source](https://youtu.be/a7baAGCBA9U?list=PLu0W_9lII9agwh1XjRt242xIpHhPT2llg)

**Declaring a class**
```python
class Person:
	name = "John"
	occupation = "Painter"
	
	def info(self):
		print(f"{self.name} is a {self.occupation}")

if __name__ == "__main__":
	p1 = Person()
	p1.name = "Joseph"
	p1.occupation = "Writer"
	p1.info()
```

**Constructor**
```python
class Person:
	name = "John"
	occupation = "Painter"
	def __init__(self , name , occupation):
		self.name = name
		self.occupation = occupation
		
	def info(self):	
		print(f"{self.name} is a {self.occupation}")

  

if __name__ == "__main__":
	p1 = Person("Akash" , "HR")
	p1.info()
```

Getters And Setters
```python
class Person:
	_name = ""
	
	def __init__(self, name: str):
	self._name = name

	@property	
	def get_name(self):
		return self._name
	
	@get_name.setter
	def set_name(self, name: str):
		self._name = name

if __name__ == "__main__":
	p1 = Person("John")
	print(p1.get_name) # John
	
	p1.set_name = "Michael"
	print(p1.get_name) # Michael
```
