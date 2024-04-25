![Source](https://youtu.be/d7ng_aV4qdI?list=PLu0W_9lII9agwh1XjRt242xIpHhPT2llg)

```python
import time

if __name__ == "__main__":

	h = time.strftime("%H")
	
	m = time.strftime("%M")
	
	s = time.strftime("%S")
	
	timestamp = f"{int(h)-12}:{m}:{s}"
	
	print(timestamp)
```
