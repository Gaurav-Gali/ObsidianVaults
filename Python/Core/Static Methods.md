![Source](https://youtu.be/GcSVYNSsJxo?list=PLu0W_9lII9agwh1XjRt242xIpHhPT2llg)

```python
class Person:
	def __init__(self):
		self.__name = "John"
		
		@staticmethod
		def greet(name:str):
			return "Hello, " + name

if __name__ == "__main__":
	p1 = Person()
	print(p1.greet("Albert"))
```

**Note** : The static method does not require the self keyword since it is not associated directly with the class.