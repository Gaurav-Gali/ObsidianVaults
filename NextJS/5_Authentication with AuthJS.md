![Source](https://youtu.be/Zq5fmkH0T78?t=4194)

**Note** : ==AuthJS was previously called NextAuth==

[AuthJS Official Documentation](https://authjs.dev/getting-started/installation?framework=Next.js)

```bash
// Installation
npm install next-auth@beta

// Setting up a secret
npx auth secret

```

==Follow the documentation for the setup==

#### Setting up GitHub OAuth
1. [Create an OAuth app in from your GitHub account](https://docs.github.com/en/apps/oauth-apps/building-oauth-apps/creating-an-oauth-app)
2. Add the respective env variables , make sure to generate the respective variables from the auth page in your GitHub account.
```.env
AUTH_GITHUB_ID=""
AUTH_GITHUB_SECRET=""
```
3. Add the following configuration in the **auth.ts** file
```auth.ts
import NextAuth from "next-auth"
import GitHub from "next-auth/providers/github"
 
export const { handlers, auth, signIn, signOut } = NextAuth({
  providers: [GitHub],
})
```


#### Using the auth object in the code
```tsx
import Link from "next/link";
import { auth, signIn, signOut } from "@/auth";

const App = async() => {
	const session = await auth();
	return (
		<>
		{session && session?.user ? (
			<div>
				<Link href="/startup/create">Create</Link>
				<form
					action={async () => {
					"use server";
					await signOut({ redirectTo: "/" });
					}}
				>
					<button>Signout</button>
				</form>
				
				<Link href={`/users/${session?.id}`}>
					{session?.user?.name}
				</Link>
			</div>
		) : (
			<div>
				<form
					action={async () => {
					"use server";
					await signIn('github');
					}}
				>
					<button>Signin</button>
				</form>
			</div>
		)}
		</>
	);
}
```

**Note** : ==Since client components like the button element with the onCLick() method cannot be used as server components , use server actions instead==
```tsx
<form
	action={async () => {
	"use server";
		await signOut();
	}}
>
	<button>Signout</button>
</form>
```
