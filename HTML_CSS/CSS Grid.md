![Source](https://youtu.be/t6CBKf8K_Ac)

**Note**:
1. Mostly use fractional units "fr" to size grids.
2. use the repeat function to describe the number of rows or columns.
3. use column-gap and row-gap to specify the gap between row and column respectively.
4. use "gap" if you want to set both the column and the row gap with the same value.
5. use grid-auto-rows to set the height of the rows.
6. use the minmax() function to support the grid-auto-rows property.

```html
<!DOCTYPE html>

<html lang="en">

<head>

<meta charset="UTF-8">

	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	
	<title>Document</title>
	
	<link rel="stylesheet" href="main.css">

</head>

	<body>
	
	<div class="item">
	
	Lorem, ipsum dolor sit amet consectetur adipisicing elit. Explicabo omnis cupiditate iste incidunt libero placeat commodi exercitationem nemo pariatur perferendis itaque dolores sit autem, qui repudiandae aut laborum voluptatum minus? Rem, qui minima est nobis consequuntur aliquid impedit modi obcaecati quod ipsum facere corporis repellendus porro, cum sunt, culpa hic.
	
	</div>
	
	<div class="item">Item 2</div>
	
	<div class="item">Item 3</div>
	
	<div class="item">Item 4</div>
	
	<div class="item">Item 5</div>
	
	<div class="item">Item 6</div>
	
	<div class="item">Item 7</div>
	
	<div class="item">Item 8</div>
	
	<div class="item">Item 9</div>
	
	</body>

</html>
```

```css
* {

	margin: 0;
	
	padding: 0;

}

  

.item {

	background-color: steelblue;
	
	color: white;
	
	border: 3px solid #000;
	
	padding: 50px;

}

  

body {

	margin: 10px;
	
	display: grid;
	
	grid-template-columns: repeat(3,1fr);
	
	column-gap: 10px;
	
	row-gap: 20px;
	
	grid-auto-rows: minmax(100px , auto);

}

  

.item:nth-of-type(1) {

	grid-column: 1/span 2;

}
```
