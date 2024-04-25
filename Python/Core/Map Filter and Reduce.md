![Source](https://youtu.be/OErhjT4f5Cs?list=PLu0W_9lII9agwh1XjRt242xIpHhPT2llg)

**MAP function**
```python
if __name__ == "__main__":
	cube = lambda x:x**3
	l = [1,2,3,4,5]
	lc = list(map(cube , l))
	print(lc) # [1, 8, 27, 64, 125]
```
The map functions takes every element of the sequence and applies the given function to it.

**FILTER function**
```python
if __name__ == "__main__":
	even = lambda x: x%2==0
	l = [1,2,3,4,5]
	le = list(filter(even, l))
	print(le) # [2, 4]
```

REDUCE function
```python
from functools import reduce

if __name__ == "__main__":
	l = [1,2,3,4,5,6]
	sum = reduce(lambda x,y:x+y , l)
	print(sum)
```
**Note** : The *reduce* function needs to be imported.