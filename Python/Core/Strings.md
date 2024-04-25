![Source](https://youtu.be/kMNFQYArrLg?list=PLu0W_9lII9agwh1XjRt242xIpHhPT2llg)

Defining a String in python
```python
if __name__ == "__main__":
	a = "This is a string"
```
**Note** : The ==str== keyword is used to define strings in python

<hr>

String slicing
```python
if __name__ == "__main__":
	a = "This is a string"
	
	# Slicing the first character
	print(a[0]) # >>> T

	# Slicing a chunk of characters
	print(a[0:3]) # Thi

	# Slicing using step count
	print(a[0:3:2]) # this will skip every 2 characters
```

**Notes** : 
1. If you want to start slicing from the starting of the string then the lower limit can be left blank.
2. If you want to slice till the end of the string then the upper limit can be left blank.
3. If the limits are in *negative* then the slicing happens from the end of the string.

<hr>

String Methods
```python
if __name__ == "__main__":

	a = "This is a string"
		
	# 1.Finding the length of the string
	print(len(a)) # 16
	
	# 2.Upper Case the string
	print(a.upper()) # THIS IS A STRING
	
	# 3.Lower Case the string
	print(a.lower()) # this is a string
	
	# 4.Striping the leading and trailing
	print(a.lstrip()) # removes leading
	print(a.rstrip()) # removes trailing
	print(a.strip()) # removes both leading and trailing
	
	# 5.Replace a part of the string
	print(a.replace('is', 'was')) # Thwas was a string

	# 6.Spliting the string	
	print(a.split(' ')) # ['This', 'is', 'a', 'string']

	# 7.Capitalising the string	
	print(a.capitalize()) # This is a string
	
	# 8.Centering the string
	print(a.center(50)) # <blanks>This is a string
	
	# 9.Counting the occurances
	print(a.count("i")) # 3
	
	# 10.Startsswith
	print(a.startswith("i")) # False

	# 11.Endsswith	
	print(a.endswith("i")) # False
	
	# 12.Index of a part of the string
	print(a.index("is")) # 2
	
	# 13.Finding a part of a string
	print(a.find("is")) # 2	
	print(a.find("strings")) # -1
	
	# 14.Finding if string is alpha numeric
	print(a.isalnum()) # False
	
	# 15.Finding if string is numeric
	print(a.isdigit()) # False

	# 16.Finding if string is alpha	
	print(a.isalpha()) # True

	# 17.To find if string is lowercase	
	print(a.islower()) # False
	
	# 18.To find if string is uppercase
	print(a.isupper()) # False
	
	# 19.To find if a string is printable
	print(a.isprintable()) # True
	
	# 20.To find if there are any whitespace character
	print(a.isspace()) # True
	
	# 21.To find if the string is in a title format
	print(a.istitle())
	
	# 22.To convert the string to a title format
	print(a.title()) # This Is A String
	
	# 23.Swapping the case of the string
	print(a.swapcase()) # tHIS IS A STRING
```

