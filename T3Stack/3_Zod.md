![Source](https://youtu.be/L6BE-U3oy80)

### Basic Types
```ts
import {z} from 'zod';  
  
export const CreateTodoValidation = z.object({  
  name: z.string().min(3),  
});  
  
export const UpdateTodoValidation = z.object({  
  id: z.string().uuid(),  
  name: z.string().min(3).optional(),  
  done: z.boolean().optional(),  
  createdAt: z.date()
});
```

### Parsing and Validation
```ts
// 1. Parse
userSchema.parse({ name: "John", age: 30 });
// ✅ Returns: { name: "John", age: 30 }

userSchema.parse({ name: "John" });
// ❌ Throws: ZodError (because `age` is missing)


// 2. SafeParse
const result = userSchema.safeParse({ name: "John", age: 30 });

// -> Returns
{
  success: true,
  data: { name: "John", age: 30 }
}

// -> Error
{
  success: false,
  error: ZodError { issues: [...] }
}
```

**NOTE** : Parse id not ideal for API's. If invalid the server crashed. Use safeParse for API's.
1. Using Parse
```ts
export async function POST(req: Request) {
  const body = await req.json();
  const data = userSchema.parse(body); // throws if invalid
  // if invalid, server crashes
  return Response.json(data);
}
```

2. Using Safe Parse
```ts
export async function POST(req: Request) {
  const body = await req.json();
  const result = userSchema.safeParse(body);

  if (!result.success) {
    return Response.json({ error: result.error.errors }, { status: 400 });
  }

  return Response.json(result.data);
}
```
