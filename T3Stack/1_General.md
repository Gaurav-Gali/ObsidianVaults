![Source](https://youtu.be/BwB-sPy6iwQ)

### Setup
```bash
pnpm create t3-app@latest
```

### Neon Config
```env
DATABASE_URL="postgresql://postgres:password@localhost:5432/t3test"
```
==add the neon url here to connect with db==

### Drizzle Setup
```bash
pnpm add drizzle-orm pg @neondatabase/serverless zod
pnpm add -D drizzle-kit tsx
```

use the following commands to make migrations and migrate
```bash
pnpm drizzle-kit generate
pnpm drizzle-kit migrate
```

### ENV.JS
*This is a very important file, as it allows to create type safe env variables*
```js
import { createEnv } from "@t3-oss/env-nextjs";  
import { z } from "zod";  
  
export const env = createEnv({  
  
  server: {  
    DATABASE_URL: z.string().url(),  
    NODE_ENV: z  
      .enum(["development", "test", "production"])  
      .default("development"),  
  },  
  
  client: {  
    // NEXT_PUBLIC_CLIENTVAR: z.string(),  
  },  
  
  runtimeEnv: {  
    DATABASE_URL: process.env.DATABASE_URL,  
    NODE_ENV: process.env.NODE_ENV,  
    // NEXT_PUBLIC_CLIENTVAR: process.env.NEXT_PUBLIC_CLIENTVAR,  
  },  
  
  skipValidation: !!process.env.SKIP_ENV_VALIDATION,  
  emptyStringAsUndefined: true,  
});
```

1. All the server only env variables go within the ==server== construct, and these variables can only be used within server components, If used in client components, then it will throw an error.
2. All the client or the public env variables go within the ==client== construct, and these variables can be accessed by the browser.
3. The actual path of the either kind of variables needs to be passed in the ==runtimeEnv== construct, this would fetch the actual variable values during runtime.
4. Here ==zod== is used for type definitions for these env variables, so that smart autocomplete can be available when we would have to access these variables.