![Source](https://youtu.be/ZcZu1NYx-WE?t=393)

1. Creating an expo project
```bash
npx create-expo-app@latest
```

2. Upon creating the project , move into the created project and run the following command to start a fresh project
```bash
npm run reset-project
```

3. Running the expo project
```bash
npm run start
```

4. Disabling headers from being seen
```javascript
<Stack screenOptions={{ headerShown: false }}>
	<Stack.Screen name="index" />
</Stack>
```

5. Adding routing
```javascript
import { useRouter } from "expo-router";

<TouchableOpacity
	onPress={() => router.push("auth/signup")}
	className="p-3 h-12 mt-3 flex items-center justify-center bg-gray-200 w-full rounded-lg border border-gray-600"
>
	<Text className="font-[roboto] text-gray-600">
		Sign up
	</Text>
</TouchableOpacity>

// router.back();
// router.replace('auth/signup');
```
**Note** : *The routing structure is similar to nextJS i.e it follows folder based routing.*

6. Disabling header using header options
```javascript
import { useRouter,useNavigation } from "expo-router";

const navigation = useNavigation();

useEffect(() => {
	navigation.setOptions({
		headerShown: false,
	});
}, []);
```

7. Setting up text inputs and password fields
```javascript
<TextInput
	secureTextEntry={true}
	className="font-[roboto] p-3 h-[50px] rounded-lg border border-gray-600"
	placeholder="Enter your password"
/>
```

8. ==TouchableOpacity vs Pressable==
	1. TouchableOpacity
		- **Purpose**: Used to handle taps with a built-in fading effect, making the component appear semi-transparent when pressed.
		- **Opacity Effect**: Reduces the opacity of the component while it's being pressed, giving visual feedback.
		- **Customization**: Limited; primarily used for simple onPress actions and provides a basic "tap" interaction without extensive customizability for press states.
		- **Common Use Case**: Ideal for simple button or touch interactions where the opacity feedback is sufficient.
	2. Pressable
		- **Purpose**: A more versatile component introduced in React Native 0.63, allowing for multiple states such as `onPressIn`, `onPressOut`, `onLongPress`, and more.
		- **Flexible Feedback**: Enables developers to define custom styles and behaviors based on various states (like `pressed`, `hovered`, `focused`).
		- **Customization**: High; you can use the `style` prop to apply different styles depending on the press state, making it more suitable for complex interactions.
		- **Common Use Case**: Useful for creating custom buttons and interactions that need detailed press feedback and styling changes.

==**Note** : Use the shortcut *cmd+shift+k* to bring up the keyboard in IOS.==

**Note** : You can delete the *app-example* folder , since it just gives an example of the file structure in the project.

<hr>

### General Folder Structure
1. **app** folder
	1. layout.tsx
		1. Contains all the stack information.
		2. The application is going to finally render through this component.
	2. index.tsx
		1. It is the base component of the project.
