![Source](https://youtu.be/2hR-uWjBAgw?t=458)

**Things to be done**
1. Sign in to firebase console.
2. Create a project and start the setup for the web apps.

3. Install firebase sdk using npm
```bash
npm install firebase
```

4. Create a config folder in the source "src" directory and add a **firebase.js** file in it.
```javascript
import { initializeApp } from "firebase/app";
import { getAnalytics } from "firebase/analytics";

const firebaseConfig = {
  apiKey: "AIzaSyDB-4g0E6FtWcDgEktL-Ny_hn-KZimiTm0",
  authDomain: "fir-react-dfba8.firebaseapp.com",
  projectId: "fir-react-dfba8",
  storageBucket: "fir-react-dfba8.appspot.com",
  messagingSenderId: "22452552145",
  appId: "1:22452552145:web:db839d8c116a6598a520f0",
  measurementId: "G-2C0QM72QYS"
};

// Initialize Firebase
const app = initializeApp(firebaseConfig);
const analytics = getAnalytics(app);
```

5. Install the firebase CLI tools globally in your machine (**will be used later for deployement**)
```bash
npm install -g firebase-tools
```
