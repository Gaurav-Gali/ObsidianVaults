[Source](https://youtu.be/eILUmCJhl64?t=2054)

1. Initially CRA (Create React App) was used. To start the server , use ==npm start==
2. Using **vite** is the modern way of creating react apps.
3. Vite enables HMR (Hot Module Replacement) i.e it provides real-time updates to modules.
4. **Vite** provides small bundle size and it is quicker. To start the server , use ==npm run dev==

### Installation

```bash
npm create vite@latest
```
**Note** : Set the project name , select the framework , language.

```shell
npm install
```
**Note** : This is to install all the required dependencies.

<hr>

### Project Structure
1. node_modules directory : It has all the installed node packages.
2. public directory : Contains static files that don't change.
3. src directory : Main folder for the react code.
	1. components directory : contains reusable parts of the UI.
	2. assets directory : contains images , fonts and other static files.
	3. style directory : contains css files.
	4. store directory : contains all the contexts.
4. package.json file : contains information of the react project and it's dependencies.
5. In package.json file , we could find 2 types of dependencies 
	1. dependencies : contains all the dependencies that are required for the project
	2. devDependencies : contains all the dependencies that is crucial for development but won't be deployed to production.
6. vite.config.json file : contains vite configs.

