![Source](https://youtu.be/0d6b6fFuCkE?list=PLu0W_9lII9agwh1XjRt242xIpHhPT2llg)

**Function with positional arguments**
```python
def average(a,b):
	return (a+b)/2

if __name__ == "__main__":
	print(average(10,5))
```

**Function with keyword  arguments**
```python
def average(a,b):
	return (a+b)/2

if __name__ == "__main__":
	print(average(a=10,b=5))
```

**Function with default arguments**
```python
def average(a,b=10):
	return (a+b)/2

if __name__ == "__main__":
	print(average(10))
```

**Variable length arguments**
```python
def average(*numbers):
	avg = 0
	for i in numbers:
		avg += i/len(numbers)
			return(avg)

if __name__ == "__main__":
	print(average(10,5,5))
```
**Note** : The type of the variable length argument is a tuple

**Arbitrary keyword arguments**
```python
def student(**kwargs):
	print("First name : " , kwargs["fname"])
	
	print("Last name : " , kwargs["lname"])

if __name__ == "__main__":
	student(fname="Albert" , lname="Einstein")
```
**Note** : The type of the *kwargs* is a dictionary

