![Source](https://youtu.be/DZiJGALRfJQ?list=PLc20sA5NNOvozFyzAoKqc-_qLluA-IeYd)

### arange() , zeros() , ones() , full() and reshape()
```python
np.arrange(1,6,2,dtype=float)
np.zeros((2,3),dtype=int) # creates 2 rows and 3 columns with zeros
# note: for zeros() the default dtype is float
np.ones((2,3),dtype=int) # creates 2 rows and 3 columns with ones
# note: for ones() the default dtype is float
np.full((2,3),5,dtype=int) # creates 2 rows and 3 columns with specified value
# note: for full() the default dtype is the type of the specified vaue
arr.reshape(2,3) # reshapes the given array without removing any elements
# note: the dimension in reshape should be factors of the size of the array.
```

### repeat() , concatenate(),vstack(),r_ keyword , hstack()
```python
arr = np.repeat([4],5,axis=0)
# note: if a 2d matrix is given and the axis is 0 then the output will be a vertical array else horizontal
np.concatenate([arr1,arr2] , axis = 0); # this allows for duplicates.
np.vstack([arr1,arr2]) # concatinates on axis 0
np.hstack([arr1,arr2]) # concatinates on axis 1
np.r_[arr1,arr2] # concatinates on axis 0
```

### power(), diff(), cross(), sort()
```python
arr = np.power(arr1 , 2) # all the elemnts will be squared
arr = np.diff(arr1) # it will subtract the next element fromt eh previous element and store the result in the previous element's index
arr = np.diff((arr1,arr2),axis=0) #subracts the elements from arr2 and arr1.
arr = np.cross(arr1,arr2) # return the cross product of the 2 arrays
arr = np.sort([1,2,3,4,5]) # returns the sorted array
arr = np.sort(([1,2,3],[4,5,6]),axis=0)
```

### flatten(),sum(),nansum(),cumprod(),cumsum()
```python
arr = np.array([[1,2,3,4],[5,6,7,8]])
arr.flatten("C") # creates a 1d array from a 2d matrix
arr.flatten("F")
arr.sum(axis=0) # sum of all values
arr.sum(axis=1) # array of sum of vertical values
arr.nansum(axis=1) # array of sum of vertical values excludes the nan values
np.cumsum(arr) # performs cummulative sum of an array
np.cumprod(arr) # performs cummulative product of an array
```