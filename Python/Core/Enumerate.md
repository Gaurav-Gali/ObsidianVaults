![Source](https://youtu.be/kGnYc_h1geM?list=PLu0W_9lII9agwh1XjRt242xIpHhPT2llg)

An **Enumerate** function is a built in function in python that return the index and the value of a iterable in the respective order.

Execution
```python
if __name__ == "__main__":
	l = ["apple" , "cake" , "biscuit" , "cheese"]
	for index , item in enumerate(l):
		print(str(index+1)+"." , item)
```

With the start argument
```python
if __name__ == "__main__":
	l = ["apple" , "cake" , "biscuit" , "cheese"]
	for index , item in enumerate(l , start=1):
		print(str(index)+"." , item)
```

