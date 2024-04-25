![Source](https://youtu.be/EX6Tt-ZW0so?list=PL-51WBLyFTg2vW-_6XBoUpE7vpmoR3ztO&t=274)
**Note** :
1. Create a file forms.py in the app directory
2. Add the {% csrf_token %} tag in the html form , since it prevents ==Cross-Site-Resource-Forgery==

In forms.py
```python
from django.forms import ModelForm
from . models import (
	OrderModel
)

class OrderForm(ModelForm):
	class Meta:
		model = OrderModel
		fields = "__all__"
```

In views.py
```python
def CreateOrderView(request):
	form = OrderForm()
	context = {
		'form' : form
	}
	if request.method == "POST":
		form = OrderForm(request.POST)
		if form.is_valid():
			form.save()
			return redirect("/")
	
	return render(request, "accounts/order_form.html" , context=context)

def UpdateOrderView(request,pk):
	order = OrderModel.objects.get(id=pk)
	form = OrderForm(instance=order)
	
	context = {
		'form' : form
	}
	
	if request.method == "POST":
		form = OrderForm(request.POST , instance=order)
		if form.is_valid():
			form.save()
			return redirect("/")
	return render(request, "accounts/order_form.html" , context=context)

def DeleteOrderView(request,pk):
	order = OrderModel.objects.get(id=pk)
	context = {
		'item' : order
	}
	if request.method == "POST":
		order.delete()
		return redirect("/")
	
	return render(request, "accounts/delete.html", context=context)
```
