![Source](https://youtu.be/G3e-cpL7ofc?t=4670)

**Note**:
1. Specificity : When multiple selectors use the same property then the order of priority is of that which has more specific definition , for example a class selector would have more priority than a paragraph selector.

### Text Elements
1. \<strong>\</strong> : It is used to strongly or boldly display the text within.
2. \<u>\</u> : It is used to underline the text within.
3. \<span>\</span> : It is the most generic text element with no styles.
4. \<i>\</i> : It is used to display the text within in italics.
5. \<b>\</b> : it is used to display the text in bold.

### Simple youtube card render
```html
<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="UTF-8" />
		<meta name="viewport" content="width=device-width, initial-scale=1.0" />
		<title>Text</title>
		<link rel="stylesheet" href="text.css" />
	</head>
	
	<body>
		<p class="video-title">
			How to write efficient code?
		</p>
		<p class="video-stats">
			4.5M views &#183 10months ago
		</p>
		<p class="video-author">
			Programmer on youtube &#10003
		</p>
		<p class="video-desc">
			Lorem ipsum dolor sit, amet consectetur adipisicing elit. Sapiente       accusamus consectetur incidunt beatae dolorum totam dignissimos est                obcaecati, voluptates adipisci nesciunt, quas doloribus nobis?
		</p>
	</body>
</html>
```

```css
* {
	margin: 0;
	padding: 10px;
}

.video-title {
	font-family: Arial, Helvetica, sans-serif;
	font-size: 30px;
	font-weight: bold;
	text-align: left;
	width: 350px;
	margin-bottom: 2px;
	line-height: 40px;
}

.video-stats {
	font-family: Arial, Helvetica, sans-serif;
	font-size: 17px;
	color: rgb(96, 96, 96);
	margin-bottom: 10px;
}

.video-author {
	font-family: Arial, Helvetica, sans-serif;
	font-size: 17px;
	color: rgb(96, 96, 96);
	margin-bottom: 10px;
	text-decoration: underline;
}

.video-desc {
	font-family: Arial, Helvetica, sans-serif;
	font-size: 17px;
	width: 350px;
	line-height: 20px;
	color: rgb(96, 96, 96);
}
```

### Result
![[text.png]]
