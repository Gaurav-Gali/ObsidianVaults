![Source](https://youtu.be/XYwJKFB8DUk?list=PLu0W_9lII9agwh1XjRt242xIpHhPT2llg)

**Recursion** is when a function is called within itself

```python
def factorial(n):
	if n == 1:
		return 1
	else:	
		return n * factorial(n-1)

if __name__ == "__main__":
	f = factorial(3)
	print(f)
```

