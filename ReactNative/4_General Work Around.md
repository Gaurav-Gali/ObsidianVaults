![Source](https://youtu.be/fX89gDDDp5M?t=6982)

#### Images
```js
import { View, StyleSheet } from "react-native";
import { Image } from "expo-image";
import React from "react";

const Img = (props) => {
	return (
		<View className={props.styl}>
			<Image
				style={
					props.radius
					? { ...styles.image, borderRadius: props.radius }
					: styles.image
				}
				
				source={
					props.src
					? props.src
					: "https://picsum.photos/seed/696/3000/2000"
				}
				
				contentFit="cover"
				transition={1000}
			/>
		</View>
	);
};

const styles = StyleSheet.create({
	image: {
		flex: 1,
		width: "100%",
		backgroundColor: "#0553",
	},
});

export default Img;
```

##### Usage
```js
<Img
	styl="h-[24px] w-[24px] border border-purple-200 rounded-full"
	radius={200}
/>
```
**Note** : The **styl** prop deals with the container of the image and not the image directly.

<hr>

#### Update User

```js
import { supabase } from "../lib/supabase";

  

export const updateUserData = async (userId, data) => {
	const { setUserData } = useAuth();
	const { error } = await supabase
		.from("users")
		.update(data)
		.eq("id", userId);
	if (error) {
		return { success: false };
	}
	return { success: true };
};
```

**Note** : Don't forget to update the user using the setUserData function from the AuthContext
```js
setUserData({ ...User, ...data });
```

#### Uploading to Supabase
1. Create a storage bucket
2. Add custom policies.

==**Note** : For react native the regular file uploads for blob png or mp4 won't work , so we 
need to use the ArrayBuffer for base 64==

1. install the library to convert the file to base64
```bash
npx expo install expo-file-system
npm i base64-arraybuffer
```

2. Essential service functions
```js
import * as FileSystem from "expo-file-system";
import { decode } from "base64-arraybuffer";
import { supabase } from "../lib/supabase";

export const uploadMedia = async (folderName, fileUri, isImage = true) => {
	
	try {
		let fileName = getFilePath(folderName, isImage);
		
		const fileBase64 = await FileSystem.readAsStringAsync(fileUri, {
			encoding: FileSystem.EncodingType.Base64,
		});
	
	// Converting image data to array buffer
		let imageData = decode(fileBase64);
	
	// Uploading to supabase storage bucket
		let { data, error } = await supabase.storage
			.from("uploads")
			.upload(fileName, imageData, {
				cacheControl: 3600,
				upsert: false,
				contentType: isImage ? "image/*" : "video/*",
			});
	
		if (error) {
			console.log("File upload error : ", error);
			return { success: false };
		}
		return { success: true, data: data.path };
	} catch (error) {
		console.log("File upload error : ", error);
		return { success: false };
	}
};

export const getFilePath = (folderName, isImage) => {
	const filePath = `${folderName}/${Date.now()}.${isImage ? "jpg" : "mp4"}`;
	return filePath;
};

export const getSupabaseFilePath = (filePath) => {

	return `https://dnduavdlmyrogrpsnjjm.supabase.co/storage/v1/object/public/uploads/${filePath}`;
};
```

#### Videos
1. install expo-av
```bash
npx expo install expo-av
```


#### Updating posts
```js
import { supabase } from "../lib/supabase";
import { uploadMedia } from "./uploadMedia";

export const createUpdatePost = async (data) => {
	const userId = data.userId;
	let caption = data.caption;
	let media = data.media;
	let fileResult = null;
		
	try {
		// Uploading media
		fileResult = await uploadMedia(
			media?.type === "images" ? "postImages" : "postVideos",
			media?.uri,
			media?.type === "images" ? true : false
		);
		
		if (fileResult.success) {
			media = fileResult.data;
		} else {
			return fileResult;
		}
	
		const { data, error } = await supabase
		.from("posts")
		.upsert({
			userId,
			body: caption,
			image: media,
		})
		.select()
		.single();
	
		if (error) {
			console.log(error);
			return {
				success: false,
			};
		} else {
			return { success: true, data: data };
		}
	} catch (error) {
		console.log(error);
		return {
			success: false,
		};
	}
};
```

#### Fetch Posts
```js

```