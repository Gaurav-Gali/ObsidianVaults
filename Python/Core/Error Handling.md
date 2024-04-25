![Source](https://youtu.be/4LKo6dlku7M?list=PLu0W_9lII9agwh1XjRt242xIpHhPT2llg)

**Structure of try and except**
```python
if __name__ == "__main__":
	try:
		a = int(input("Enter a number : "))
	except ValueError:
		print("ValueError")
	except Exception as e:
		print("Error : " , e)
```

[Finally Keyword](https://youtu.be/r_iuC-IDpPM?list=PLu0W_9lII9agwh1XjRt242xIpHhPT2llg)
```python
if __name__ == "__main__":
	try:
		a = int(input("Enter a number : "))
	except ValueError:
		print("ValueError")
	except Exception as e:
		print("Error : " , e)
	finally:
		print("Always Executed")
```
**Note** : The finally block is executed irrespective of the the execution of the try and except block
**Also** : when using finally in a function , the block will be executed even it is after a return statement.

