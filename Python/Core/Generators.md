![Source](https://youtu.be/ixd-u3pmsUc?list=PLu0W_9lII9agwh1XjRt242xIpHhPT2llg)

Generators are special functions in python that allow you to create an **iterable** sequence of values.

Printing values with the ***next()*** function
```python
def num_gen(n: int):
	for i in range(1, n+1):
		yield i

if __name__ == "__main__":
	gen = num_gen(10)
	
	print(next(gen)) # 1
	print(next(gen)) # 2
	print(next(gen)) # 3
```

Printing values with a loop
```python
def num_gen(n: int):
	for i in range(1, n+1):
		yield i

if __name__ == "__main__":
	gen = num_gen(10)
	
	for i in gen:
		print(i)
```

