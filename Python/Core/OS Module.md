![Source](https://youtu.be/dkVYSsL90Oo?list=PLu0W_9lII9agwh1XjRt242xIpHhPT2llg)

To create a directory
```python
import os

if __name__ == "__main__":
	os.mkdir("data")
```

To check if the path is correct
```python
import os

if __name__ == '__main__':
	if os.path.exists("data") == False:
	os.mkdir("data")
```

Renaming Files
```python
import os
import time

if __name__ == '__main__':
	if os.path.exists("data") == False:
		os.mkdir("data")
		
	for i in range(10):
		os.mkdir(f"data/day{i}")
		
	time.sleep(1)
	
	for i in range(10):
		os.rename(f"data/day{i}" , f"data/DAY{i}")
```

Listing out directories
```python
import os

if __name__ == '__main__':
	if os.path.exists("data") == False:
		os.mkdir("data")
		folders = os.listdir("data")
	
	for folder in folders:
		print(folder)
```

To get current working directory
```python
import os

if __name__ == '__main__':
	print(os.getcwd())
```

 Some Important ones
 1. os.remove()
	 1. removes a file
2. os.rmdir
	1. removes a empty directory
3. os.chdir("")
	1. changes the directory
4. shutil.rmtree()
	1. removes a directory and all it's files