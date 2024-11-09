![Source](https://youtu.be/d2kcCNkLVbA?list=PLc20sA5NNOvozFyzAoKqc-_qLluA-IeYd)

### Creating a matrix and some basic functions
```python
mat1 = np.matrix([[1,2],[4,5]])
mat2 = np.matrix([[7,8],[10,11]])
# Some common matrix operations
np.multiply(mat1,mat2)
np.dot(mat1,mat2)
np.cross(mat1,mat2)
```

### More functions
```python
mat = np.matrix([[1,2],[3,4]])

np.transpose(mat)
mat.T

mat.trace() #sum of diagonal elements
mat.diagonal() #returns the diagonal elements

np.linalg.det(mat) #determinant of the matrix
np.linalg.matrix_rank(mat) #rank of matrix
np.linalg.inv(mat) #returns the inverse of the matrix
```

