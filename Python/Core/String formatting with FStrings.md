![Source](https://youtu.be/ixmxgUf8yIg?list=PLu0W_9lII9agwh1XjRt242xIpHhPT2llg)

**Using the format method** *(old method)*
```python
if __name__ == "__main__":
	s = "hello I am {} and I am {} years old"
	s1 = s.format("John" , 20)
	print(s1)
```

**Using fstring**
```python
if __name__ == "__main__":
	name = "John"
	age = 20
	print(f"hello I am {name} and I am {age} years old")
```

**Using fstring with decimal point precision**
```python
if __name__ == "__main__":
	product = "Apple"
	cost = 10.874
	print(f"{product} costs {cost:.2f} rupees") # Apple costs 10.87 rupees
```

**Escape sequencing the string**
```python
if __name__ == "__main__":
	product = "Apple"
	cost = 10.874
	print(f"{{product}} costs {{cost:.2f}} rupees") # {product} costs {cost:.2f}     rupees
```

