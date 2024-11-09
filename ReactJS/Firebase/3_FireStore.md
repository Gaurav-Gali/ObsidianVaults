![Source](https://youtu.be/2hR-uWjBAgw?t=2157)


1. Import from "firebase/firestore"
```javascript
import {getFirestore} from "firebase/firestore";

const app = initializeApp(firebaseConfig);
export const db = getFirestore(app);
```

2. Reading data from DB
```javascript
// Firebase Imports
import { db } from "../config/firebase";
import { collection, getDocs } from "firebase/firestore";

const [movies, setMovies] = useState([]);

const moviesCollectionRef = collection(db, "movies");

try {
	const data = await getDocs(moviesCollectionRef);
	const filteredData = data.docs.map((doc) => {
		return { ...doc.data(), id: doc.id };
	});
	setMovies(filteredData);
} catch (error) {
	console.error("Error: ", error);
}
};
```
**Note** : Make sure to change the rules in the firebase console in order to make changes.

<hr>

### Posting data to the database
1. Import addDoc from "firebase/firestore"
```javascript
import { collection, getDocs,addDoc } from "firebase/firestore";

const moviesCollectionRef = collection(db, "Movies");

const onMovieSubmit = async () => {
	const data = {
		"title" : "ABC",
		"year" : 2005,
		"hasWonAwards" : true
	};
	await addDoc(moviesCollectionRef , data);
};
```

<hr>

### Deleting data from the database
1. Import the following
```javascript
import { collection, getDocs,addDoc , deleteDoc } from "firebase/firestore";
```

2. Implementing deletion
```javascript
const deleteMovie = async (movieId) => {
	const movieDoc = doc(db , "Movies" , moviesId);
	await deleteDoc(movieDoc);
};
```

<hr>

### Updating data from the database
1. Importing the follwoing
```javascript
import { collection, getDocs,addDoc , updateDoc } from "firebase/firestore";
```

2. Implementing updation
```javascript
const updateMovie = async (movieId) => {
	const movieDoc = doc(db , "Movies" , moviesId);
	await updateDoc(movieDoc , {title : updatedTitle});
};
```
