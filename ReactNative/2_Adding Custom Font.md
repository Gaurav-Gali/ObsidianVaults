![Source](https://youtu.be/ZcZu1NYx-WE?t=1161)

1. Add the following in the **layout.tsx** file to configure the installed fonts.
```javascript
import { Stack } from "expo-router";
import { useFonts } from "expo-font";

export default function RootLayout() {
	//Defining fonts
	useFonts({
		// Playfair Display
		"playfair": require("../assets/fonts/playfair_display/PlayfairDisplay-Regular.ttf"),
		
		"playfair-medium": require("../assets/fonts/playfair_display/PlayfairDisplay-Medium.ttf"),
		
		"playfair-bold": require("../assets/fonts/playfair_display/PlayfairDisplay-Bold.ttf"),
	
		// Roboto
		"roboto": require("../assets/fonts/roboto/Roboto-Regular.ttf"),
		"roboto-medium": require("../assets/fonts/roboto/Roboto-Medium.ttf"),
		"roboto-bold": require("../assets/fonts/roboto/Roboto-Bold.ttf"),
	
	});
	
	return (
		<Stack>
			<Stack.Screen name="index" />
		</Stack>
	);

}
```