![Source](https://youtu.be/wIPHER2UBB4?list=PL-51WBLyFTg2vW-_6XBoUpE7vpmoR3ztO)

### One to Many
 ![[onetomany.png]]
**Note** : Here a customer can have many orders but an order can have only one customer.
```python
customer = models.ForeignKey(CustomerModel, null=True, on_delete=models.SET_NULL)
product = models.ForeignKey(ProductModel, null=True, on_delete=models.SET_NULL)
```

### Many to Many
![[manytomany.png]]
**Note** : Here a product can have many tags and a tag can have many products.
Django by default creates a third table to keep tally of the relationship.
```python
tags = models.ManyToManyField(TagModel)
```
