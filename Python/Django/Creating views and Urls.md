![Source](https://youtu.be/PtQiiknWUcI?t=2129)
**Note**:
1. Create a separate urls.py file in each app for better routing.
2. The path function can take 3 arguments
	1. The url path
	2. A function from the views file
	3. A name property for using django template tags.
### views.py
```python
# BaseApp views.py
from django.shortcuts import render
from django.http import HttpResponse

# Views
def Home(request):
	return HttpResponse("Home Page")

def Room(request):
	return HttpResponse("Room Page")
```

### Root urls.py
```python
# Root urls.py
from django.contrib import admin
from django.urls import path , include

urlpatterns = [
	# Admin Urls
	path('admin/', admin.site.urls),
	
	# BaseApp Urls
	path('' , include('base.urls')),
]
```

### App urls.py
```python
# BaseApp urls.py
from django.urls import path
from .views import (
	Home,
	Room
)

urlpatterns = [
	path('' , Home , name='home'),
	path('room/' , Room , name='room'),
]
```

### Creating dynamic urls

1. urls.py file
```python
path('room/<str:pk>/' , Room , name='room')
```
**Note** :
1. str : a regular string , "room/1"
2. int : a regular integer , "room/1"
3. slug : "room/room-id-1"

2. views.py file
```python
def Room(request, pk):
	cur_room = None
	
	for room in rooms:
		if room["id"] == int(pk):
			cur_room = room
		
	context = {
		'room': cur_room
	}
		
	return render(request, 'base/room.html' , context)
```
