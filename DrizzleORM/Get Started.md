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
import { boolean, integer, pgTable, text, timestamp } from "drizzle-orm/pg-core";

export const todo = pgTable("todo", {
	id: integer("id").primaryKey(),
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
	schema: "./db/schema.ts",
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

