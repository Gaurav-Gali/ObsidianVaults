![Source](https://youtu.be/Tfc5LfHSGmE?t=182)

### Installation
```bash
npm i drizzle-orm @neondatabase/serverless dotenv
npm i -D drizzle-kit
```

### Connection with NeonDB
1. Get the env variable from neon
2. Make sure to choose AWS as it provides better integration with other tools

### Setting up
1. @/db/drizzle.ts
```ts
import { drizzle } from "drizzle-orm/neon-http";
import { neon } from "@neondatabase/serverless";
import * as schema from "@/db/schema";

const database_url = process.env.DATABASE_URL!;

const sql = neon(database_url);
export const db = drizzle(sql, { schema });
```

2. @/db/schema.ts
```ts
import { boolean, uuid, pgTable, text, timestamp } from "drizzle-orm/pg-core";  
  
export const todo = pgTable("todo", {  
  id: uuid("id").primaryKey().defaultRandom(),  
  name: text("name").notNull(),  
  done: boolean("done").default(false).notNull(),  
  createdAt: timestamp("created_at").notNull().defaultNow(),  
});
```

3. @/drizzle.config.ts
```ts
import { config } from "dotenv";
import { defineConfig } from "drizzle-kit";

config({path: ".env"})

export default defineConfig({
	out: "./drizzle_migrations",
	schema: "./db/schema.ts", // or "./db/schemas/*.ts" if creating seperate files for each schema
	dialect: "postgresql",
	dbCredentials: {
		url: process.env.DATABASE_URL!,
	},
});
```
==**Note** : This files runs under the NodeJS environment, so you would have to manually load the .env file==

4. Creating migrations and pushing it to NeonDB
```bash
// Creates the migrations
npx drizzle-kit generate

// Applies the changes
npx drizzle-kit migrate
```

5. @/actions/todo-actions.ts
```ts
"use server";

import { db } from "@/db/drizzle";
import { todo } from "@/db/schema";
import { eq } from "drizzle-orm";

export const getTodos = async () => {
	const data = await db.select().from(todo);
	return data;
};

export const addTodo = async (name: string) => {
	await db.insert(todo).values({
		name: name,
	});
};

export const deleteTodo = async (id: string) => {
	await db.delete(todo).where(eq(todo.id, id));
};

export const updateTodo = async (id: string, name: string) => {
	await db
	.update(todo)
	.set({
		name: name,
	})
	.where(eq(todo.id, id));
};
```
==**Note** : always make the requests when in server mode, this does not work in client mode.==

### Full Crud Example
```ts
import { db } from "@/server/db";  
import { tables } from "@/server/db/schema";  
import { desc, eq } from "drizzle-orm";  
import type { TodoType } from "@/types/todoType";  
  
export const getAllTodos = async (): Promise<TodoType[]> => {  
  return db.select().from(tables.todo).orderBy(desc(tables.todo.createdAt));  
};  
  
export const createTodo = async (name: string): Promise<TodoType | undefined> => {  
  const result = await db  
    .insert(tables.todo)  
    .values({  
      name,  
      done: false,  
    })  
    .returning();  
  
  return result[0];  
};  
  
export const updateTodo = async (  
  id: string,  
  data: { name?: string; done?: boolean }  
): Promise<TodoType | undefined> => {  
  const result = await db  
    .update(tables.todo)  
    .set({  
      ...data,  
      ...(data.name !== undefined && { name: data.name }),  
      ...(data.done !== undefined && { done: data.done }),  
    })  
    .where(eq(tables.todo.id, id))  
    .returning();  
  
  return result[0];  
};  
  
export const deleteTodo = async (id: string): Promise<void> => {  
  await db.delete(tables.todo).where(eq(tables.todo.id, id));  
};
```

### Creating proper route handlers
```ts
import { NextResponse } from "next/server";  
import { getAllTodos, createTodo, updateTodo, deleteTodo } from "@/server/actions/todoActions";  
  
export async function GET(): Promise<NextResponse> {  
  const todos = await getAllTodos();  
  return NextResponse.json(todos);  
}  
  
export async function POST(request: Request): Promise<NextResponse> {  
  try {  
    const body = await request.json() as { name?: string };  
    const { name } = body;  
  
    if (!name || typeof name !== "string" || name.trim() === "") {  
      return NextResponse.json(  
        { error: "Todo name is required" },  
        { status: 400 }  
      );  
    }  
  
    const newTodo = await createTodo(name.trim());  
  
    if (!newTodo) {  
      return NextResponse.json(  
        { error: "Failed to create todo" },  
        { status: 500 }  
      );  
    }  
  
    return NextResponse.json(newTodo, { status: 201 });  
  } catch (error) {  
    return NextResponse.json(  
      { error: "Failed to create todo" },  
      { status: 500 }  
    );  
  }  
}  
  
export async function PATCH(request: Request): Promise<NextResponse> {  
  try {  
    const body = await request.json() as { id?: string; name?: string; done?: boolean };  
    const { id, ...data } = body;  
  
    if (!id) {  
      return NextResponse.json(  
        { error: "Todo ID is required" },  
        { status: 400 }  
      );  
    }  
  
    const updatedTodo = await updateTodo(id, data);  
  
    if (!updatedTodo) {  
      return NextResponse.json(  
        { error: "Todo not found" },  
        { status: 404 }  
      );  
    }  
  
    return NextResponse.json(updatedTodo);  
  } catch (error) {  
    return NextResponse.json(  
      { error: "Failed to update todo" },  
      { status: 500 }  
    );  
  }  
}  
  
export async function DELETE(request: Request): Promise<NextResponse> {  
  try {  
    const body = await request.json() as { id?: string };  
    const { id } = body;  
  
    if (!id) {  
      return NextResponse.json(  
        { error: "Todo ID is required" },  
        { status: 400 }  
      );  
    }  
  
    await deleteTodo(id);  
  
    return NextResponse.json({ success: true, message: "Todo deleted" });  
  } catch (error) {  
    return NextResponse.json(  
      { error: "Failed to delete todo" },  
      { status: 500 }  
    );  
  }  
}
```

