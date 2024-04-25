![Source](https://youtu.be/FPu9Uld7W-E)

**Note** :
1. Remember ==Time Complexity== is **not** equal to time taken to run the program
2. Definition : 
	1. The rate at which the time taken increases with respect to the input size.

### The Big O Notation
(a.k.a) The upper bond context
**Rules :**
1. Always calculate time complexity with the worst case scenario or the ==Big O Notation==.
2. Avoid Constants.
3. Avoid lower values.

Example
1. Linear Time Complexity
```cpp
int n = 5;
for(int i = 0 ; i < n ; i++) {
	cout << i << endl;
}
// Here the loop is executed n times
// Thus there are n number of tasks being executed
// Time complexity will be O(n)
```

2. Quadratic Time Complexity
```cpp
int n = 5;
for(int i = 0 ; i < n ; i++) {
	for(int j = 0; j < n; j++) {
		cout << j << endl;
	}
}
// Here the outer loop is executed n times and the inner loop is executed n times
// Thus there are nxn number of tasks being executed
// Time complexity will be O(n^2)
```
