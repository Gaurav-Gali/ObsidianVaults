![Source](https://youtu.be/TCWOwavqFrw?list=PLu0W_9lII9agwh1XjRt242xIpHhPT2llg)

To use regex import the **re** module , which is a in built module.

**Match characters in regex**
1. [] : Represent a character class
2. ^ : Matches the beginning
3. $ : Matches the end
4. . : Matches any character except newline
5. ? : Matches zero or one occurence
6. | : Means OR (Matches any of the characters separated by it)
7. * : Any number of occurrence (including zero occurrence)
8. + : One or more occurrences
9. {} : Indicates number of occurrences of a preceding RE

Using the **search()** method
```python
import re

if __name__ == '__main__':
	pattern = r"met[A-Z]eorology"
	
	text = "In metKeorology, a cyclone is a large air mass that rotates around a             strong center of low atmospheric pressure, counterclockwise in the               Northern Hemisphere and clockwise in the Southern Hemisphere as                  viewed from above (opposite to an anticyclone)."
	
	match = re.search(pattern, text)
	
	print(match) # <re.Match object; span=(3, 15), match='metKeorology'>
```
**Note** : Here "r" before the string is used to specify a **raw string**

Using the **finditer()** method
```python
import re

if __name__ == '__main__':
	pattern = r"met[A-Z]eorology"
	
	text = "In metKeorology, a cyclone is a large air mass that rotates around a             strong center of low atmospheric pressure, counterclockwise in the               Northern Hemisphere and clockwise in the Southern Hemisphere as                  viewed from above (opposite to an anticyclone)."
	
	matches = re.finditer(pattern, text)
	
	for match in matches:
		indice_x, indice_y = match.span()[0], match.span()[1]
		word = text[indice_x : indice_y]
		print(word) # metKeorology
```

