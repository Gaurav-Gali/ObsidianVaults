![Source](https://youtu.be/2hR-uWjBAgw?t=809)

### Email / Password Authentication

1. Import the required functionalities from 'firebase/auth'
```javascript
import { getAuth } from "firebase/auth";
```

2. Create an auth object to handle all authentication
```javascript
// Initialize Firebase
const app = initializeApp(firebaseConfig);
export const auth = getAuth(app);
```
**Note** : Make sure you export the auth object in order to use it in the project.

3. To create an user with the email and password provided
```javascript
import { auth } from "../config/firebase";
import { createUserWithEmailAndPassword } from "firebase/auth";

const handleSignIn = async () => {
	try {
		await createUserWithEmailAndPassword(
			auth,
			authData.email,
			authData.password
		);
	} catch (error) {
		console.error("Error: ", error);
	}
};
```
**Note** : Make sure to use async-await , since firebase usually returns promises to be handled.
**Note** : Make sure to handle errors when working with async await.

<hr>

### Handling Logout
1. Import signOut from 'firebase/auth'
```javascript
import { auth } from "../config/firebase";
import { signOut } from "firebase/auth";
```

2. Create a function to handle the logging out
```javascript
const handleSignOut = async () => {
	try {
		await signOut(auth);
	} catch (error) {
		console.error("Error: ", error);
	}
};
```

<hr>
### Google Auth Provider
**Note** : Make sure to enable Google auth in the firebase authentication console.

1. Import the google auth provider from 'firebase/auth'
```javascript
import { getAuth, GoogleAuthProvider } from "firebase/auth";
```

2.  Create a provider object to implement google auth
```javascript
const app = initializeApp(firebaseConfig);
export const auth = getAuth(app);
export const googleAuthProvider = new GoogleAuthProvider();
```

3. Implementing google auth
```javascript
import { auth , googleAuthProvider } from "../config/firebase";
import {signInWithPopup } from "firebase/auth";

const handleGoogleSignIn = async () => {
	try {
		await signInWithPopup(auth, googleAuthProvider);
	} catch (error) {
		console.error("Error: ", error);
	}
};
```
