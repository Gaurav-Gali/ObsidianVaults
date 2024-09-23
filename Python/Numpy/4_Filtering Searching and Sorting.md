![Source](https://youtu.be/NHcGFYJMhHk?list=PLc20sA5NNOvozFyzAoKqc-_qLluA-IeYd)

### Searching
```python
np.where(arr==4),
np.where(arr>=4),
np.where(arr!=4),
np.where(arr<=4)
```

### Filtering
```python
a = np.array([1,2,6,4,7,3,5])

fa1 = [i for i in range(len(a)) if i % 2 == 0] # filter array
fa2 = a > 3

print(a[fa1])
print(a[fa2])
```

### Insertion and Deletion
```python
 np.insert(arr,(2,4,6),25)
 np.delete(arr,0)
```