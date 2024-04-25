![Source](https://youtu.be/scWc3F8LbOo?list=PLu0W_9lII9agwh1XjRt242xIpHhPT2llg)


```python
if __name__ == "__main__":
	l = ["apple", "orange", "pear"]
	
	# 1.Appending items to the list
	l.append("peach")
	print(l) # ['apple', 'orange', 'pear', 'strawberry']
	
	# 2.Sorting a list
	l.sort() # sorts in ascending order
	print(l) # ['apple', 'orange', 'peach', 'pear']
	
	l.sort(reverse=True) # sorts in descending order
	print(l) # ['pear', 'peach', 'orange', 'apple']

	# 3.Reverse the list	
	l.reverse()
	print(l) # ['apple', 'orange', 'peach', 'pear']
	
	# 4.Index of teh first occurence of an item in the list
	i = l.index("pear")
	print(i) # 3
	
	# 5.Counting the number of occurences of an item in the list
	c = l.count("orange")
	print(c) # 1
	
	# 6.Copying the items of the list
	l1 = l.copy() # this allows memory reassignment
	print(l1)
	
	# 7.Inserting values into a list
	l.insert(1,"tomato"
	print(l) # ['apple', 'tomato', 'orange', 'peach', 'pear']
	
	# 8.Extending a list
	l1 = ["more fruits"]
	l.extend(l1)
	print(l) # ['apple', 'tomato', 'orange', 'peach', 'pear', 'more fruits']
	
	# 9.Concatinating lists
	l1 = [1,2,3,4]
	l2 = ["a","b","c"]
	l3 = l1+l2
	print(l3) # [1, 2, 3, 4, 'a', 'b', 'c']

	# 10. Poping values from a list
	k = l.pop()
	print(k) # more fruits
	print(l) # ['apple', 'tomato', 'orange', 'peach', 'pear']
```

