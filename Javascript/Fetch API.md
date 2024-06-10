![Source](https://youtu.be/CyGodpqcid4?list=PLGjplNEQ1it_oTvuLRNqXfz_v_0pq6unW)

The fetch api provides an interface for receiving and sending resources.
It uses the Request and Response Objects.

==**Note** : API stands for Application Programming Interface==

The **fetch()** method takes 2 arguments
1. The API endpoint or the URL.
2. An array of options , and when this argument is not mentioned , then by default it takes a GET request.

###### Few Terms
1. AJAX
	1. It stands for Asynchronous JS And XML
	2. This format was used in the early days of JS , it is now rarely used.
2. JSON
	1. It stands for JavaScript Object Notation
	2. This is the most widely used format for sending and receiving data.
	3. This is formally known as AJAJ since JSON has replaced XML , but since it sounds odd , we refer to it as JSON.

==To use the JSON format in JS we use the **json()** method which returns a second promise which contains the parsed JSON text in JS object format. Basically Input -> JSON ; Output -> Object==

###### Handling fetch using Promises
```js
const URL = "https://cat-fact.herokuapp.com/facts";

fetch(URL).then((res) => {
	return res.json();
}).then((data) => {
	data.map((fact) => {
		console.log(fact.text);
		});
	});
```
###### Handling fetch using Async Await
```js
const URL = "https://cat-fact.herokuapp.com/facts";

const getCatFacts = async () => {
	let response = await fetch(URL);
	let data = await response.json();
	
	data.map((element) => {
		console.log(element.text);
	});
};

getCatFacts();
```

<hr>

### HTTP Verbs / Request Methods
<table border="1px">
<tr>
<th>Method</th>
<th>Description</th>
</tr>
<tr>
<td>GET</td>
<td>Requests data from a specified resource</td>
</tr>
<tr>
<td>POST</td>
<td>Submits data to be processed to a specified resource</td>
</tr>
<tr>
<td>PUT</td>
<td>Updates a current resource with new data</td>
</tr>
<tr>
<td>DELETE</td>
<td>Deletes the specified resource</td>
</tr>
<tr>
<td>HEAD</td>
<td>Similar to GET but returns only HTTP headers and no document body</td>
</tr>
<tr>
<td>OPTIONS</td>
<td>Describes the communication options for the target resource</td> 
</tr>
<tr>
<td>PATCH</td>
<td>Applies partial modifications to a resource</td>
</tr>
</table>

<hr>

### HTTP Status Codes
![[httpErrCode.JPG]]

<hr>
### Request and Response
