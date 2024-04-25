![Source](https://youtu.be/G-Rct7Na0UQ?list=PL-51WBLyFTg2vW-_6XBoUpE7vpmoR3ztO)
**Note**:
1. Install the django-filter library
```bash
conda install django-filter
```
2. After installing add this library to the installed apps
```python
INSTALLED_APPS = [
	'django.contrib.admin',
	'django.contrib.auth',
	'django.contrib.contenttypes',
	'django.contrib.sessions',
	'django.contrib.messages',
	'django.contrib.staticfiles',
	
	# Created Apps
	'accounts.apps.AccountsConfig',
	
	# Installed
	'django_filters',
]
```
