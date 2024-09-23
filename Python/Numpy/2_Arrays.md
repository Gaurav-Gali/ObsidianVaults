![Source](https://www.youtube.com/watch?v=eIyCnvVhFXo&list=PLc20sA5NNOvozFyzAoKqc-_qLluA-IeYd&index=2&pp=iAQB)

**Note**
	Arrays refer as a collection of items that are stored at contiguous memory locations.
	It is a container that stores fixed number of elements of the same type.
	Combination of arrays can save a lot of time and reduces the size of the code.

**Advantages of NumPy arrays**
1. Uses much less memory to store data.
2. Easy to perform mathematical operations.
3. Used for creation of n-dimensional arrays.
4. Finding elements is easy.

### Creating np arrays
```python
arr = np.array([1,2,3,4,5])
```
**Note** : the type of the array will be ==<class 'numpy.ndarray'>==

### Slicing arrays
```python
# Slicing 1-d arrays
arr[:2] # returns the first 2 elements
arr[:6:3] # returns the first 6 elements with a gap of 3 elements

# Slicing 2-d arrys
arr2[:2] # returns the firt 2 arrays of the matrix
arr2[1,2:4] # returns the elements from 2 to 4 in the second array in the matrix
arr2[:3,:3] # returns the first 3 elemnts of the first and the second array in the matrix
```

### Finding the shape , size, dimension and datatype of the np array
```python
np.shape(arr) # (3,4)
np.size(arr) # 12
np.ndim(arr) # 2
arr.dtype # dtype('int64')
```

<hr>

# List vs Numpy Array

1. A list cannot directly handle mathematical operations , while array can.
2. Array consumes less memory than a list.
3. Arrays are faster than list.
4. A list can store different datatypes , while a array cannot.
