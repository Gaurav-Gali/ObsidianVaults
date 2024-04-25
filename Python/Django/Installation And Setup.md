![Source](https://youtu.be/PtQiiknWUcI)

1. Installing Django
```bash
conda install django
```

2. Starting a project
```bash
django-admin startproject
```

3. Making the migrations
```bash
python manage.py makemigrations
```

4. Migrating
```bash
python manage.py migrate
```

5. Starting the server
```bash
python manage.py runserver 8000
```

6. Starting the app
```bash
python manage.py startapp appname
```
**Note** : 
1. When creating apps , don't forget to add it to the installed apps in the settings.py file.
2. To do so , in the settings.py file in the installed apps , add the following
```python
INSTALLED_APPS = [
	"appname.apps.AppnameConfig" # This provides the actual path to the app
]
# Here the "appname" is a placeholder for the actual app's name
```
3. For setting up static and media files
```python
STATIC_URL = 'static/'
MEDIA_URL = 'images/'

STATICFILES_DIRS = [
	BASE_DIR/"static"
]
```
