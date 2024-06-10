![Source](https://youtu.be/7zcXPCt8Ck0?list=PLGjplNEQ1it_oTvuLRNqXfz_v_0pq6unW)

**Note** :
1. DOM stands for Document Object Model
2. It is a way of accessing the HTML code within JavaScript
3. **Window Object** : 
	1. It is a window object and not an JavaScript object.
	2. It holds all the properties of the current window and it is also a global object.

### DOM Manipulation
1. getElementById("id");
2. getElementByClass("class");
3. getElementByTagName("tag");
4. Query Selectors
	1. document.querySelectror() , returns the first element
	2. document.quertSelectorAll() , return a NodeList
5. Properties :
	1. tagName , it return the name of the tag of the selected node.
	2. innerText , it returns the text content of the element and all it's children.
	3. innerHTML , it returns the plain text or the HTML content in the element.
	4. textContent , returns textual content even for hidden elements.
6. Attributes :
	1. getAttribute("attr");
	2. setAttribute("attr" , value);
7. Style :
	1. node.style
8. Create:
	1. document.createElement("button");
9. Insert Elements
	1. node.append() , adds at the end of the node (outside)
	2. node.prepend() , adds at the start of the node (outside)
	3. node.before() , adds before the node
	4. node.after() , adds after the node
10. Delete Elements
	1. node.remove()