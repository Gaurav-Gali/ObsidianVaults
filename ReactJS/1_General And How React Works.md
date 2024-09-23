![Source](https://youtu.be/eILUmCJhl64?t=270)

1. React is a JavaScript library used to build frontend UI.
2. It is developed by FaceBook on 2011 and is used to create dynamic single page applications.
3. We use components in react in order to make the code reusable.

<hr>

### How react works
1. React component
	1. The **App.jsx** is the main or the root component of the react app.
	2. It is the starting point of the entire application.
2. Virtual DOM
	1. React creates an in memory structure called the ==virtual DOM==
	2. A virtual DOM is different from the browser DOM.
	3. It's a light weight representation of where each node stands for a component and it's attributes.
3. Reconciliation Process
	1. When the component data changes , react updated the virtual DOM to mirror these changes.
	2. ==React then compares the current and the previous state of the virtual DOM.==
	3. It identifies the specific nodes that require updating.
	4. Only those nodes will be updated in the **browser's DOM**.

