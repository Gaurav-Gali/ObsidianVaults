![Source](https://youtu.be/fX89gDDDp5M?t=3931)

**Official Docs For User Management**
https://supabase.com/docs/guides/auth/managing-user-data?queryGroups=language&language=js

1. Install the supabase client library
```bash
npx expo install @supabase/supabase-js @react-native-async-storage/async-storage @rneui/themed react-native-url-polyfill
```

2. Create a supabase.js file  to setup the client configuration inside the lib folder and create the lib folder if not already there.
```js
// supabase.js
import { AppState } from "react-native";
import "react-native-url-polyfill/auto";
import AsyncStorage from "@react-native-async-storage/async-storage";
import { createClient } from "@supabase/supabase-js";

// Importing environment variables
import { SUPABASE_URL, SUPABASE_API_KEY } from "@env";

const supabaseUrl = SUPABASE_URL;
const supabaseAnonKey = SUPABASE_API_KEY;


export const supabase = createClient(supabaseUrl, supabaseAnonKey, {
	auth: {
		storage: AsyncStorage,
		autoRefreshToken: true,
		persistSession: true,
		detectSessionInUrl: false,
	},
});

// Tells Supabase Auth to continuously refresh the session automatically
// if the app is in the foreground. When this is added, you will continue
// to receive `onAuthStateChange` events with the `TOKEN_REFRESHED` or
// `SIGNED_OUT` event if the user's session is terminated. This should
// only be registered once.

AppState.addEventListener("change", (state) => {
	
	if (state === "active") {
		supabase.auth.startAutoRefresh();
	} else {
		supabase.auth.stopAutoRefresh();
	}

});
```

3. Signing Up the user
```js
const { data, error } = await supabase.auth.signUp({
	email,
	password,
	options: {
		data: {
			name
		}
	}
});

if (error) {
	Alert.alert("Couldn't Sign Up", error.message);
}
```

**Note** : Don't forget to create a RLS policy for a newly created table for the triggers to work. And for the "using" query , just put "true" in it.

4. Setting up the triggers in the SQL editor page
```sql
-- inserts a row into public.profiles
create function public.handle_new_user()
returns trigger
language plpgsql
security definer set search_path = public
as $$
begin
	insert into public.users (id, name , email)
	values (new.id, new.raw_user_meta_data ->> 'name', new.raw_user_meta_data ->> 'email');
	return new;

end;
$$;

-- trigger the function every time a user is created
create trigger createAuthUser
	after insert on auth.users
	for each row execute procedure public.handle_new_user();
```
**Note** : Change this as per your app's requirement

*Follow these steps to setup the function and triggers*
![Source](https://youtu.be/fX89gDDDp5M?t=4658)

<hr>

#### Creating a user context for accessing user data in very page
1. Create a *AuthContext.js* file in the contexts folder in the root
```js
// contexts/AuthContext.js
import { createContext, useContext, useState } from "react";

const AuthContext = createContext();

export const AuthProvider = ({ children }) => {
	const [user, setUser] = useState(null);
	
	const setAuth = (authUser) => {
		setUser(authUser);
	};
	
	const setUserData = (userData) => {
		setUser({ ...userData });
	};
	
	return (
		<AuthContext.Provider value={{ user, setAuth, setUserData }}>
			{children}
		</AuthContext.Provider>
	);

};

export const useAuth = () => {
	return useContext(AuthContext);
};
```

2. Make the context available for all the pages by updating the *_layout.tsx* file
```ts
// _layout.tsx
const _layout = () => {
	<AuthProvider>
		<RootLayout />
	</AuthProvider>
};

const RootLayout = () => {

	//Defining fonts  
	
	// Implementing authentication
	const { setAuth, setUserData } = useAuth();
		const router = useRouter();
		  
		
		useEffect(() => {
			supabase.auth.onAuthStateChange((_event, session) => {
				if (session) {
					setAuth(session?.user);
					updateUserData(session?.user);
					router.replace("/main/home");
				} else {
					setAuth(null);
					router.replace("/welcome");
				}
			});
		}, []);

	const updateUserData = async (user) => {
		const { success, data } = await getUserData(user?.id);
		if (success) {
			setUserData(data);
		}
	};
	
	return (
		<Stack screenOptions={{ headerShown: false }} />
	);
};

export default _layout;
```

3. Getting user data from database
```js
// functions/getUserData.js
import { supabase } from "../lib/supabase";

export const getUserData = async (userId) => {	
	const { data, error } = await supabase
	.from("users")
	.select()
	.eq("id", userId)
	.single();
	
	if (error) {
		return { success: false, data: error?.message };
	}
	return { success: true, data };
};
```

5. Finally set the *index.tsx* for loading
```ts
// index.tsx
import { SafeAreaView, ActivityIndicator } from "react-native";

export default function Index() {
	return (
		<SafeAreaView className="flex-1 bg-purple-50 items-center justify-center">
			<ActivityIndicator size={25} />
		</SafeAreaView>
	);
}
```

4. Implementing *Logging out*
```js
const handleLogout = async () => {
	const { error } = await supabase.auth.signOut();
	if (error) {
		Alert.alert("Couldn't Log Out", error.message);
	}
};
```

<hr>

#### Setting up environment variables

1. Adding the .env file in the root of the project
```.env
EXPO_PUBLIC_API_KEYS=XYZABCD
```
**Note** : Make sure to follow the naming convention , EXPO_PUPLIC_(Your variable name) = (Value)

2. Use the Environment variables in the required files using process.env
```js
const KEY = process.env.EXPO_PUBLIC_API_KEYS;
```

