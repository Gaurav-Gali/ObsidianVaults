![Source](https://youtu.be/HOrutCnp2zo?list=PLu0W_9lII9agwh1XjRt242xIpHhPT2llg)

```python
if __name__ == "__main__":
	s1 = {1,2,3,4,5}
	s2 = {4,5,6,7,8}
	
	# 1.Union of sets
	s = s1.union(s2)
	print(s) # {1, 2, 3, 4, 5, 6, 7, 8}
	
	# 2.Updating the sets
	s1.update(s2)
	print(s1) # {1, 2, 3, 4, 5, 6, 7, 8}
	
	# 3.Intersection of sets
	s = s1.intersection(s2)
	print(s) # {4, 5, 6, 7, 8}
	
	s1.intersection_update(s2)
	print(s1) # {4, 5, 6, 7, 8}
	
	# 4.Symmetric difference between sets
	s = s1.symmetric_difference(s2)
	print(s)
	
	s1.symmetric_difference_update(s2)
	print(s1)
	
	# 5.Difference between sets
	s = s1.difference(s2)
	print(s)

	s1.difference_update(s2)	
	print(s1)
	
	# 6.Is disjoint in sets
	print(s1.isdisjoint(s2)) # True
	
	# 7.Is supersets
	print(s1.issuperset(s2)) # False
	
	# 8.Is subset
	print(s1.issubset(s2)) # True
	
	# 9.Adding a single item to the set
	s1.add(20)
	print(s1)
	
	# 10.Adding multiple elements to the set
	s1.update(s2)
	print(s1)
	
	# 11.Removing elements from the set
	s1.remove(20) # causes error is the element is not there
	print(s1)
	
	# 12.Discarding elements from the set
	s1.discard(20) # does not cause error is the element is not there
	print(s1)

	# 13.pop a random element from the set	
	print(s1.pop())
	
	# 14.clearing all elements
	s1.clear()
	
	# 15.deleting the set
	del s1
```

