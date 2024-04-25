![Source](https://youtu.be/SJzsNd7SM0g?list=PLu0W_9lII9agwh1XjRt242xIpHhPT2llg)

A doctring provides the information of the functions useage
```python
def sum(*numbers):
	'''Returns the sum of the numbers in the given arguments'''

	sum = 0
	for i in numbers:
	
		sum += i
	
	return sum

if __name__ == "__main__":
	s = sum(1,2,3)
	print(s)
	print(sum.__doc__)
```
**Note** : The doctstring should be mentioned right after the function header

What is PEP-8 ?
It stands for ==**P**ython **E**nhancement **P**roposal== 
