![Source](https://youtu.be/PtQiiknWUcI?t=4055)

**Note**:
1. Models are used to define the structure of a database.
![[models.png]]
2. After creating a model always make migrations to reflect it in the database.

### Admin and Superuser
1. Creating a superuser
```bash
python manage.py createsuperuser
```

2. To go to the admin panel visit the admin/ route.
3. In order to see the created models in the admin panel add the model to the admin in the admin.py file
```python
from .models import (
	Room
)

# Register your models here.
admin.site.register(Room)
```
