![Source](https://youtu.be/PTBZ674EsvI?list=PLu0W_9lII9agwh1XjRt242xIpHhPT2llg)

**Standard Decorator**
```python
def greeter(fx):
	def mfx():
		print("Welcome")
		fx()
		print("Thank You")
	
	return mfx

@greeter
def message():
	print("I am John")

if __name__ == "__main__":
	message()
```

**Decorator With Arguments**
```python
def greeter(fx):
	def mfx(*args, **kwargs):
		print("Welcome")
		fx(*args, **kwargs)	
		print("Thank You")
		
	return mfx

@greeter
def message(msg):
	print(msg)

if __name__ == "__main__":
	message("I am Joseph")
```

