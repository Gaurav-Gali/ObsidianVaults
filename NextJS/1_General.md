![Source](https://youtu.be/Zq5fmkH0T78?t=663)

[Official Documentation](https://nextjs.org/docs)

```bash
// Creates a NextJS application
npx create-next-app@latest

// Runs the NextJS application in development mode
npm run dev
```

**Note** :
To ensure that all the other dependencies use the latest version of react, add the following to the package.json file and ==don't delete the node modules directory, it might break things==
```json
// Add this above the dependencies object
"packageManager": "nрm@11.0.0",
"overrides": {
	"react": "$react",
	"react-dom": "$react-dom"
},
```
#### File and Folder Structure
1. tsconfig.json
	1. It defines what should be type checked, ignored and the rules to be followed.
2. tailwind.config.ts
	1.  It defines the tailwindcss configurations and additional configs can be setup.
3. postcss.config.mjs
	1. It governs all plugins that go with CSS, in this case it uses tailwind css as a plugin.
4. package-lock.json
	1. It locks the projects versions and all other versions of the dependencies and their sub-dependencies.
5. package.json
	1. It is a simpler version of the package-lock.json file.
	2. It stores all the current dependencies.
	3. It has scripts that will be ran to when the app has to be served.
	4. Contents in the script
		1. dev : runs the project in development mode with HMR(Hot Module Replacement) support.
		2. build : builds a optimised production ready version of the project.
		3. start : runs the project in production mode.
		4. lint : runs the eslint for all files.
6. public folder
	1. This folder contains all the static data for the project, like images and other static media.
7. app folder
	1. The most important folder where all the work would be done.
	2. page.tsx
		1. This page serves the homepage of the application
	3. layout.tsx
		1. All changes made to this file will affect the entire project.
		2. Hence all the fonts and metadata are stored here.


<hr>

### Layout file
**Note** : ==The layout.tsx file acts as a parent for all the pages in that directory. Hence you would want to add the Navbar and Footer in the main layout.tsx file so that they would appear on all pages automatically.==


<hr>

### Theory

**Note** : ==By default NextJS will convert all components into server side component unless specified something else.==
So if we try to console.log() in a server component , then the the result will be logged in the running server(check the terminal).
But nextJS will also display it in the browser under the tag of **Server** to specify that it has been rendered from the server.

#### Server Components
1. These components are rendered on the server and the compiled HTML is sent to the client.
2. This reduces the amount of JS that would be sent to the client.
3. This increases performance and speed.
4. Since the components are rendered from the server, the components would be able to access the vital server side resources like Databases, files etc.

#### Client Components
1. These components are rendered on the client side(browser).
2. In order to make a component a client component, use the =='use client'== directive at the top of the component.
3. Client components are used when the components requires interactivity.


#### Server Side Pre-rendering
1. Note that a server component is only rendered on the server whereas the client component is rendered both on the server and the client.
2. The HTML is rendered and a static shell is sent to client and the client hydrates it to fully render it. This process is called server side pre-rendering.
3. The content of the HTML that don't require interactivity will be rendered on the server side and the contents which requires interactivity will be left as placeholders which will be filled later when the client tries to hydrate it.

#### Static Site Generation (SSG)
1. SSG pre-renders the static sites of the website(the pages which dose not have any dynamic data) into HTML during the build process.
2. This HTML is directly served to the browser, hence improving SEO and improves performance.

#### Incremental Static Regeneration (ISR)
1. It is a powerful technique in NextJS15 where it bridges static and dynamic content.
2. It allows to pre-render the static HTML and only modify the dynamic parts of the page when required.
3. This proves to be efficient because the entire page won't have to be re-built.

<hr>
### Error pages
NextJS allows us to create a special file in each directory called **error.tsx** that can be used to display the custom error page when any error occurs.
**Note** : Error pages must be client components.
In order to address global errors, use the global-error.tsx file to display the error messages.

### Loading pages
It is really simple to display loaders , just create a loading.tsx file which contains the content that will be rendered when the page is loading.

<hr>

