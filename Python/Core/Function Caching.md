![Source](https://youtu.be/K8gjSwc3Rlo?list=PLu0W_9lII9agwh1XjRt242xIpHhPT2llg)

Import functool in order to perform function caching.

```python
import functools
import time

@functools.lru_cache(maxsize=None)
def slow_run(n):
	time.sleep(2)
	
	return n

if __name__ == '__main__':
	print(slow_run(5))
	print(slow_run(6))
	print(slow_run(7))
	
	print()
	
	print(slow_run(5)) # runs instantly
	print(slow_run(6))
	print(slow_run(7))
```


==Important==
What is **Memoization** ?
```
In computing, memoization is used to speed up computer programs by eliminating the repetitive computation of results, and by avoiding repeated calls to functions that process the same input.
```


