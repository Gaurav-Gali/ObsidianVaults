![Source](https://youtu.be/Yk7EM06czXI?list=PLc20sA5NNOvozFyzAoKqc-_qLluA-IeYd)

### Iterateable
```python
for i in np.nditer(arr):
	print(i)

for ind , val in np.ndenumerate(arr):
	print(ind , val)
```

### Copy vs View
```python
#copy vs view
a2 = np.arange(10)

a3 = a2.copy()
a4 = a2.view()

a2[0] = 10

print(a3)
print(a4)
```
**Note** : The changes made in the original array will not be reflected in the copy but will be reflected in the view.

### Splitting an array
```python
np.split(a2,2,axis=0)
np.split(a2.reshape(2,5),5,axis=1)
```


### shuffle(),unique(),resize(),ravel()
```python
np.random.shuffle(arr)
np.unique(arr)
np.resize(arr,(2,3))
arr.ravel()
```