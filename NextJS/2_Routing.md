![Source](https://youtu.be/Zq5fmkH0T78?t=1674)

### Routing
**Note** : ==NextJS follows file based routing, where each folder that u create will be considered as a new route and each route/folder will contain it's own layout.tsx and page.tsx files.==

1. Creating routes (say : localhost:3000/dashboard/analytics)
	1. creates a dashboard directory
		1. creates a page.tsx and layout.tsx for this route if required.
		2. creates an analytics directory
			1. creates a page.tsx and layout.tsx for this route if required.
2. Creating dynamic routes (say : localhost:3000/users/id)
	1. creates a users directory
		1. creates a page.tsx and layout.tsx for this route if required.
		2. creates a \[id] directory. Note that it is enclosed with square brackets because it is a dynamic route.
	2. to access the dynamic data
```ts
import React from "react";

const page = async ({ params }: { params: Promise<{ id: string }> }) => {
	const id = (await params).id;
	return (
		<div>
			<h1>Page ID: {id}</h1>
		</div>
	);
};

export default page;
```
==**Note** : In earlier versions of NextJS, params was a synchronous prop but in the newer version it is a asynchronous promise.==


### Route Groups
**Note** : ==You would want to organise the code as the project grows , so it is a good idea to store some pages in a separate folder.
Since NextJS uses file based routing, the new folder will act as a new route.
To prevent this we use route groups.
Simply wrap the folder name with parenthesis and move all the files including the respective layout.tsx files into it.==

#### Accessing Routes / Linking
There are a few ways of performing linking in NextJS
1. Link tag
	1. The link tag has attributes similar to the anchor tag but is far more powerful.
	2. The main advantage of using this tag is that, it prefetches the content of the following routes as the respective links appear on screen.
	3. This results in faster linking and somewhat minimal loading time.
	4. Usage is as follows
```tsx
import Link from 'next/link'

<Link href="dhasboard/">Dashboard</Link>
```
2. useRouter hook
	1. This hook can be used to perform routing.
	2. Although it dose not automatically prefetches the content of the following routes like the Link tag, it can be done manually when needed.
	3. This hook allows for better control over the routing with the following functions.
```tsx
'use client'
import { useRouter } from 'next/navigation'

const router = useRouter();

const App = () => {
	// 1. The push function
	// This function allows to navigate to the new page while adding the page to the browser's history stack.
	router.push("dashboard/", {scroll: true});

	// 2. The replace function
		// This function allows to navigate to the new page whitout adding the page to the browser's history stack.
	router.replace("dashboard/", {scroll: true});

	// 3. The refresh function
	// This function is used to refresh the current route and relaod the contents.
	router.refresh();

	// 4. The prefetch function
	// This fucntion is used to prefetch the contents of the next route inorder to minimise loading.
	router.prefetch("dashboard/");

	// 5. The forward function
	// This function is used to navigate to the next route.
	router.forward();

	// 6. The back fucntion
	// This function is used to navigate to the previous route.
	router.back()
}
```


#### Parallel Routing
- This is a new feature introduced in NextJS15 and allows us to optimise and make the application more performant.
- Here components of a page can be extracted as it's own page and can be given it's own loading and error pages.
- This provides an easier and concise approach than code-splitting.

#### Implementation
1. Say we have 2 components **analytics** and **profile** , we then create 2 directories called **@analytics** and **@profile** and within it we create the respective page.tsx, loading.tsx and error.tsx files.
2. When required to use the pages , it can be imported to the layout.tsx just like components.
``
- @analytics/page.tsx/
```tsx
import React from "react";

const Analytics = () => {
	return (
		<h1>Analytics page</h1>
	);
};

export default Analytics;
```

- @profile/page.tsx/
```tsx
import React from "react";

const Profile = () => {
	return (
		<h1>Profile page</h1>
	);
};

export default Profile;
```

- layout.tsx/
```tsx
import React from "react";

const RootLayout = ({children,analytics,profile}:{children:React.ReactNode, analytics:React.ReactNode, profile:React.ReactNode}) => {
	return (
		<div>
			<div>{children}</div>
			<div>{analytics}</div>
			<div>{profile}</div>
		<div/>
	);
};

export default RootLayout;
```


**Note** : If you want to add addition follow routes, then the following has to be taken into consideration
1. Say you create a new route dashboard/settings, in that case you would also have to create the settings route within all the parallel routes as well.
2. The best practise is to use the default.tsx file in the parallel route directory. This file would be rendered when the route has nothing to do. So usually it just returns a null.
3. Alternatively, you can also create a \[...catchAll] route that would render the contents in it's page.tsx when the route won't have access to the required route.


<hr>

#### Intercepting Routes
This way of routing is used when you would want to visit a particular route but when you visit it from another specific route , you want to display a different result.
- Say, you have routes ==profile== and ==users== when visiting the profile page separately you would see the contents of the page.tsx within the profile directory.
- But when navigating to the profile page from the users page, you might want to see a different result.
- This can be accomplished using intercepting routes.

Say the folder structure is as follows
1. app
	1. profile
		1. page.tsx
	2. users
		1. (..)profile
			1. page.tsx
		2. page.tsx

**Note** : use (..) this to move back one directory and use (...) to move directly to the root.
Here when the profile page is visited the contents in the page.tsx of the (..)profile directory will be rendered.

#### Modals
A very common use case of intercepting routes is when using modals. The problem this addresses is when creating a modal, the url wont be sharable.
==Hence with the help of parallel routes with intercepting routes we can created sharable modals.==
[Implementaion](https://nextjs.org/docs/app/building-your-application/routing/intercepting-routes)
