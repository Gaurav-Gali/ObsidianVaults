![Source](https://youtu.be/PtQiiknWUcI?t=2434)
**Note** :
1. Don't forget to add the templates path to the settings.py file.
2. The templates folder should be in the root directory.
3. You can also create templates specifically for apps , to do that
	1. create a templates folder within that app
	2. create a folder with the same name as that app inside the templates folder.
	3. create the required templates inside this folder.
	4. to use these templates in the views.py file use 'appname/templatename.html'

### Extending And Including Templates
1. Use the {% extends 'file.html' %} to extend from another template
2. use the {% include 'file.html' %} to include that template into the current file

### Url template tag
1. Use the url template tag to access the required urls.
2. {% url 'path' %} , here the path is the name that was mentioned in the path function in the urls file.

### Using dynamic data
1. Mention the data / variable within double curly braces =={{data}}==
2. This data needs to passed in the render function in the function in the views file.

### Different Tags
1. CSRF {% csrf_token %} , for handling forms
2. Conditionals 
```html
{% if cond %}
{% elif cond %}
{% else %}
{% endif %}
```
3. Loops
```html
{% for i in data %}
	{{forloop.counter}} - {{i}}
	{{forloop.counter0}} - {{i}}
	{{forloop.revcounter}} - {{i}}
	{{forloop.revcounter0}} - {{i}}
	{% comment %} These are conditionals that return true or false {% endcomment %}
	{{forloop.first}}
	{{forloop.last}}
	{{empty}}
{% endfor %}
```
3. Length filter , ===data | length===


### Static Files
1. Mention the path of the static files in the settings.py file
```python
STATICFILES_DIRS = [

BASE_DIR/"static"

]
```
3. use the {%load static%} template tag to use the static files.
```html
{% load static %}
<img src="{% static 'my_app/example.jpg' %}" alt="My image">
```

3. Mention the path of the media files in the settings.py file
```python
MEDIA_URL = "images/"
```