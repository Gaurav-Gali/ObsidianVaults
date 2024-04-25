![Source](https://youtu.be/j2G68uQtOwM?list=PLu0W_9lII9agwh1XjRt242xIpHhPT2llg)

**Getting the keys of a dictionary**
```python
if __name__ == "__main__":
	d = {
		1 : "a",
		
		2 : "b",
		
		3 : "c",
	}
	
	print(d.keys()) # dict_keys([1, 2, 3])
```

**Getting the values of a dictionary**
```python
if __name__ == "__main__":
	d = {
		1 : "a",
		
		2 : "b",
		
		3 : "c",
	}
	
	print(d.values()) # dict_values(["a", "b", "c"])
```

**Getting both key-value pairs**
```python
if __name__ == "__main__":
	d = {
		1 : "a",
		
		2 : "b",
		
		3 : "c",
	}
	
	print(d.items()) # dict_items([(1, 'a'), (2, 'b'), (3, 'c')])
```

<hr>

**Methods used in a dictionary**
```python
if __name__ == "__main__":
	d = {
	
		1: "a",
		
		2: "b",
		
		3: "c",
	
	}
	
	# 1.Updating the dictionary
	d1 = {
		4: "d",
		5: "e",
	}
	d.update(d1)
	print(d) # {1: 'a', 2: 'b', 3: 'c', 4: 'd', 5: 'e'}
	
	# 2.Popping a dictionary
	print(d.pop(5)) # e
	
	# 3.Popitem a dictionary
	print(d.popitem()) # pops the last keyvalue pair
	
	# 4.delete an item from a dictionary
	del d[3]
	print(d) # {1: 'a', 2: 'b'}
	
	# 5. Clearing a dictionary
	d.clear()
	print(d) # {}
	
	# 6.Delete the entire dictionary	
	del d
```

