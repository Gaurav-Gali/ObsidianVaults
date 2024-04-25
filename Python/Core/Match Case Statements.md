![Source](https://youtu.be/bthQCK1QAmQ?list=PLu0W_9lII9agwh1XjRt242xIpHhPT2llg)

**Note** : This feature is available from python 3.10 onwards

```python
if __name__ == "__main__":
	stat = input("Are you an adult : ")
	
	match stat:
		case "yes":
			print("You can vote")
		case "no":
			print("You can't vote")
		case _ if stat == "idk":
			print("Thinking about it")
		case _ :
			print("Error")
```

