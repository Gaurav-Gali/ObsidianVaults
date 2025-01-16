![Source](https://youtu.be/b3ue9WL5fk8?list=PLC3y8-rFHvwjOKd6gdf4QtV1uYNiQnruI)

**Note** : Route handlers can be created with the route.ts file within it's route directory to server as api endpoints on that particular route.
It's crucial that the route.ts and page.tsx don't be in the same level. 
Hence, it's best practise to create an api directory and put the route.ts file within that to avoid unexcepted behaviours.
If the page.tsx and route.ts are kept on the same level, then the route.ts takes priority.

Convention : Create an api folder in the app directory and then create the routes within it.
- app
	- api
		- comments
			- route.ts
- data.ts

```ts
// data.ts
type Comment = {
	id: number;
	content: string;
	author: string;
	timestamp: Date;
};

export const comments: Comment[] = [
	{
		id: 1,	
		content: "This is a great comment!",
		author: "John Doe",
		timestamp: new Date(),
	},
	
	{
		id: 2,	
		content: "I agree with this comment.",
		author: "Jane Smith",
		timestamp: new Date(),
	},
	
	{
		id: 3,	
		content: "This is a very good comment!",
		author: "Alice Johnson",
		timestamp: new Date(),
	},
	
	{
		id: 4,	
		content: "I think this comment needs more work.",
		author: "Bob Brown",
		timestamp: new Date(),
	}

];
```

```ts
// comments/route.ts
import { comments } from "@/app/data";
import { NextRequest, NextResponse } from "next/server";

export const GET = async (request: NextRequest) => {
	return NextResponse.json({ data: comments });
};

// Post request

export const POST = async(request: Request) => {
	try {
		const data = await request.json();
		return NextResponse.json(
			{ message: "Data received successfully", data },
			{ status: 200 }
		);
	} catch (error) {
		return NextResponse.json(
			{ error: "Failed to process request" },
			{ status: 500 }
		);
	}
}
```



