![Source](https://youtu.be/G3e-cpL7ofc?t=126)

1. HTML stand for **Hyper Text Markup Language**
2. CSS stands for **Cascading Style Sheet**
3. The anchor tag has an attribute ; target which can be used to manipulate the way the anchor tag operates the hyper link.
	1.  **\_blank:** It opens the link in a new window.
	2. **\_self:** It is the default value. It opens the linked document in the same frame.
	3. **\_parent:** It opens the linked document in the parent frameset.
	4. **\_top:** It opens the linked document in the full body of the window.
	5.  **framename:** It opens the linked document in the named frame.
4. In html, the elements which do not require a closing tag are known as **void** elements.
5. There are 3 types of implementation for **css**:
```html
<!--Inline-->
<h1 style="color:blue;">

<!--Internal (written in the head tag)-->
<head>
	<style>
		button {
			color : blue;
		}
	</style>
<head>

<!--External (written in a external file with a .css extension)-->
<!--Include the external css in the html file using the link tag within the head tag-->
<link rel="stylesheet" href="file.css" />
```
5. **CSS** rules:
	```css
		button {
			color : blue;
		}
	```
**Note** : Here the *button* is know as the css selector ; *color* is the css property ; *blue* is the css value.
### Tags
<table>
<thead>
<tr>
<th>Tags</th>
<th>Description</th>
<th>Syntax</th>
<th>Example</th>
</tr>
</thead>
<tbody>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-doctypes/" target="_blank">!DOCTYPE html</a></td>
<td>According to the HTML specification or standards, every HTML document requires a document type declaration.</td>
<td>&lt; !DOCTYPE html &gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/8c359013-f2b2-41e7-ba95-931e36070300" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-abbr-tag/" target="_blank">abbreviation</a></td>
<td>The abbreviation tag in HTML is used to define the abbreviation or short form of an element.</td>
<td>&lt;abbr title=” “&gt; … &lt;/abbr&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/3342d775-914f-491b-85a4-edce11526f20" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-acronym-tag/" target="_blank">acronym</a></td>
<td>The acronym tag in HTML is used to define the acronym that gives useful information to browsers, translation systems, and search engines.</td>
<td>&lt;acronym title=” “&gt; … &lt;/acronym&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/4ac7eaa7-949b-424b-ac20-664156d65751" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-address-tag/" target="_blank">address</a></td>
<td>The address tag in HTML indicates the contact information of a person or an organization.</td>
<td>&lt;address&gt; … &lt;/address&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/4ac7eaa7-949b-424b-ac20-664156d65751" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-a-tag/" target="_blank">anchor</a></td>
<td>The anchor tag in HTML is used to create a hyperlink on the webpage.</td>
<td>&lt;a herf=” “&gt; …&lt;/a&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/76255ee4-7e25-4adb-8179-54089bcddad0" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-applet-tag/" target="_blank">applet</a></td>
<td>The applet tag in HTML was used to embed Java applets into any HTML document, discontinued starting from HTML 5.</td>
<td>&lt;applet&gt;….&lt;/applet&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/eb8fece5-013d-401d-9e49-f2b9751c447d" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-area-tag/" target="_blank">area</a></td>
<td>This area tag is used in an HTML document to map a portion of an image to make it clickable by the end-user.</td>
<td>&lt;area&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/2baafdb8-2565-400c-99f2-429da60d5edf" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html5-article-tag/" target="_blank">article</a></td>
<td>The &lt;article&gt; tag is one of the new sectioning element in HTML5. The tag is used to represent an article.</td>
<td>&lt;article&gt;..&lt;/article&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/ec315711-c30a-4269-aa20-155140feb255" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html5-aside-tag/" target="_blank">aside</a></td>
<td>The &lt;aside&gt; tag is used to describe the main object of the web page in a shorter way like a highlighter.</td>
<td>&lt;aside&gt;..&lt;/aside&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/bee6ba9f-7ff7-4129-9b15-5fc00dea85d8" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html5-audio/" target="_blank">audio</a></td>
<td>It is a useful tag if you want to add audio such as songs, interviews, etc. on your webpage.</td>
<td>&lt;audio&gt;..&lt;/audio&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/cf08e6ce-3a60-4712-ab76-9553e7ca621a" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-base-tag/" target="_blank">base</a></td>
<td>The HTML base tag is used to specify a base URI, or URL, for relative links. This URL will be the base URL for every link on the page.</td>
<td>&lt;base href = ” “&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/b176d4c5-fbf3-47ca-a81b-4c5d0f92af5a" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-basefont-tag/" target="_blank">basefont</a></td>
<td>This tag is used to set the default text-color, font-size, &amp; font-family of all the text in the browser. Not supported in HTML5.</td>
<td>&lt;basefont&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/3f49f812-4c1d-49ca-a28f-d5a94c3458e6" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-5-bdi-tag/" target="_blank">bdi</a></td>
<td>The bdi tag refers to Bi-Directional Isolation. It differentiates a text from other text that may be formatted in a different direction.</td>
<td>&lt;bdi&gt; … &lt;/bdi&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/56b68eca-4c08-4bb6-a213-6ae91507c98e" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-bdo-tag/" target="_blank">bdo</a></td>
<td>The bdo stands for Bi-Directional Override. This tag is used to specify the text direction or used to change the current direction.</td>
<td>&lt;bdo dir&gt; Contents… &lt;/bdo&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/1eedb320-3f03-4604-b7be-92428a429e6b" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-bgsound-tag/" target="_blank">bgsound</a></td>
<td>The bgsound tag is used to play the soundtrack in the background.</td>
<td>&lt;bgsound src=””&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/884628bb-8082-47d6-9980-18bfc1b1099f" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-big-tag/" target="_blank">big</a></td>
<td>The big tag in HTML is used to increase the selected text size by one larger than the surrounding text. In HTML 5.</td>
<td>&lt;big&gt; Contents… &lt;/big&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/fb04bf93-e46a-4b0c-b7e7-4f1526c32cef" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-blockquote-tag/" target="_blank">blockquote</a></td>
<td>The blockquote tag in HTML is used to display the long quotations (a section that is quoted from another source).</td>
<td>&lt;blockquote&gt; Contents… &lt;/blockquote&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/1172477a-5475-4316-a513-b24530b36b42" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-body-tag/" target="_blank">body</a></td>
<td>The body tag in HTML is used to define the main content present inside an HTML page.</td>
<td>&lt;body&gt; Contents… &lt;/body&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/3e50b8b9-02b7-472d-beb0-118ca8cd66f0" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-b-tag/" target="_blank">bold</a></td>
<td>The bold tag in HTML is used to specify the bold text without any extra importance.</td>
<td>&lt;b&gt;… &lt;/b&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/2b6b5739-75d5-41e2-b3c8-a7376c80b312" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-br-tag/" target="_blank">break</a></td>
<td>The break tag inserts a single carriage return or breaks in the document. This element has no end tag.</td>
<td>&lt;br&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/b40f4dd0-e365-493f-8861-d17630662e80" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-button-tag/" target="_blank">button</a></td>
<td>The button tag in HTML is used to define the clickable button. &lt;button&gt; tag is used to submit the content.</td>
<td>&lt;button type = “button”&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/c8885468-2806-4343-9a74-cf3c5819329d" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-caption-tag/" target="_blank">caption</a></td>
<td>The caption tag is used to specify the caption of a table. Only one caption can be specified for one table.</td>
<td>&lt;caption align = “value”&gt;&lt;/caption&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/b15cf3aa-162c-4cad-8dc3-2072db70539d" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-canvas-tag/" target="_blank">canvas</a></td>
<td>It can be used to draw paths, boxes, texts, gradients, and add images.</td>
<td>&lt;canvas id = “script”&gt; Contents&lt;/canvas&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/1e827cf7-6b58-499d-89f8-a612c0f14823" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-center-tag/" target="_blank">center</a></td>
<td>The center tag in HTML is used to set the alignment of text in the center. Not supported in HTML5.</td>
<td>&lt;center&gt; Contents.&lt;/center&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/72a3f541-063d-4720-bc7d-0ca776faaa49" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-cite-tag/" target="_blank">cite</a></td>
<td>The cite tag in HTML is used to define the title of a work. It displays the text in italic format.</td>
<td>&lt;cite&gt;Content&lt;/cite&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/15b9917f-ea49-4813-b695-4f9f83a7e281" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-code-tag/" target="_blank">code</a></td>
<td>The code tag in HTML is used to define the piece of computer code.</td>
<td>&lt;code&gt;Contents&lt;/code&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/d02a0558-5a22-4aa7-ba0d-360d00e2dca1" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-colgroup-tag/" target="_blank">colgroup</a></td>
<td>It is useful for applying styles to entire columns, instead of repeating the styles for each column, and for each row</td>
<td>&lt;colgroup&gt; Column lists &lt;/colgroup&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/c0a600b3-a654-4326-8a0d-56ec00986c99" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-col-tag/" target="_blank">column</a></td>
<td>The col tag in HTML is used to set the column properties for each column within a colgroup tag.</td>
<td>&lt;col attribute = “value”&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/84a095ca-a8f0-486b-b7e8-09f01f7702b5" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-comments/" target="_blank">comment</a></td>
<td>The comment tag is used to insert comments in the HTML code.</td>
<td>&lt;!–…–&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/cca7d706-a84d-4564-90d6-43ee41658ce0" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-data-tag/" target="_blank">data</a></td>
<td>The data element gives an address to a given content with a machine-readable translator.</td>
<td>&lt;data value=””&gt; Contents &lt;/data&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/1a0ec5c9-3a3a-4040-8fa7-b3e97fb71bf2" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-datalist-tag/" target="_blank">datalist</a></td>
<td>The datalist tag is used to provide autocomplete feature &amp; used with an input tag so that users can easily fill the data in the forms using select the data.</td>
<td>&lt;datalist&gt;Contents&lt;/datalist&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/b52cb9c9-96c5-4cd2-b0b5-29c890a43c71" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-dd-tag/" target="_blank">dd</a></td>
<td>The dd tag is used to denote the description or definition of an item in a description list.</td>
<td>&lt;dd&gt;Contents&lt;/dd&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/1a0ec5c9-3a3a-4040-8fa7-b3e97fb71bf2" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-dfn-tag/" target="_blank">define</a></td>
<td>The define tag in HTML represents the definition element and is used to represent a defining instance in HTML.</td>
<td>&lt;dfn&gt;Contents&lt;/dfn&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/02aa783f-ce0a-4387-8005-b7a323eedaf4" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-del-tag/" target="_blank">delete</a></td>
<td>Delete tag is used to mark a portion of text which has been deleted from the document.</td>
<td>&lt;del&gt;Contents&lt;/del&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/bfa5ead6-6e8b-4726-8431-60b71a75e7b8" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html5-details-tag/" target="_blank">details</a></td>
<td>This tag is used to create an interactive widget that the user can open or close.</td>
<td>&lt;details&gt;Contents&lt;/details&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/f87f1e56-3a9c-4f95-9630-7190025880ba" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html5-dialog-tag/" target="_blank">dialog</a></td>
<td>This tag is used to create a popup dialog and models on a web page. This tag is new in HTML5.</td>
<td>&lt;dialog open&gt; Contents… &lt;/dialog&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/7b8a656c-151e-470d-bf55-d30545f661f7" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-dir-tag/" target="_blank">dir</a></td>
<td>The dir tag is used to make a list of directory titles. It is not supported in HTML 5 &lt;ul&gt; or CSS are used instead of &lt;dir&gt; tag.</td>
<td>&lt;dir&gt; Lists… &lt;/dir&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/874df67b-bffa-46f3-9e00-f80e4545d71b" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/div-tag-html/" target="_blank">div</a></td>
<td>The div tag is used in HTML to make divisions of content in the web page (text, images, header, footer, navigation bar, etc).</td>
<td>&lt;div&gt;Content&lt;/div&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/8438f64f-6fa1-4b5a-b870-72476ebab7b1" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-dl-tag/" target="_blank">dl</a></td>
<td>The dl tag in HTML is used to represent the description list. In HTML4.1, it defines definition list and in HTML5, it defines description list.</td>
<td>&lt;dl&gt; Contents… &lt;/dl&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/a474fa79-b809-4ffc-800f-dc28352d2af5" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-dt-tag/" target="_blank">dt</a></td>
<td>The dt tag in HTML is used to specify the description list. It is used inside the &lt;dl&gt; element. It is usually followed by a &lt;dd&gt; tag.</td>
<td>&lt;dt&gt; Content… &lt;/dt&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/fc66bf22-dbb9-4919-b332-ac25c9ea9c27" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-embed-tag/" target="_blank">embed</a></td>
<td>It is used as a container for embedding plug-ins such as flash animations.</td>
<td>&lt;embed attributes&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/fc66bf22-dbb9-4919-b332-ac25c9ea9c27" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html5-fieldset-tag/" target="_blank">fieldset</a></td>
<td>The fieldset tag in HTML5 is used to make a group of related elements in the form, and it creates the box over the elements.</td>
<td>&lt;fieldset&gt;Contents&lt;/fieldset&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/8d82f9b2-eb1d-405b-9de9-53f0bfb70b7f" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html5-figcaption-tag/" target="_blank">figcaption</a></td>
<td>The figurecaption tag in HTML is used to set a caption to the figure element in a document. This tag is new in HTML5.</td>
<td>&lt;figcaption&gt; Figure caption &lt;/figcaption&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/08df89d6-49d1-42c0-b87e-a11a5673f7c6" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html5-figure-tag/" target="_blank">figure</a></td>
<td>The figure tag in HTML is used to add self-contained content like illustrations, diagrams, photos, or codes listed in a document.</td>
<td>&lt;figure&gt; Image content… &lt;/figure&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/14302f74-3ab1-44d8-8735-c14189e3a89a" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-font-tag/" target="_blank">font</a></td>
<td>The font tag in HTML plays an important role in the web page to create an attractive and readable web page.</td>
<td>&lt;font attribute = “value”&gt; Content &lt;/font&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/988c87fd-00a8-4c05-9889-c4d2ac5ee772" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html5-footer-tag/" target="_blank">footer</a></td>
<td>The footer tag in HTML is used to define a footer of HTML document. This section contains the footer information.</td>
<td>&lt;footer&gt; … &lt;/footer&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/09e0b34a-630b-4e82-8a5a-6ee67decaa7a" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-form-tag/" target="_blank">form</a></td>
<td>This form is used basically for the registration process, logging into your profile on a website or creating your profile on a website, etc …</td>
<td>&lt;form&gt; Form Content… &lt;/form&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/6968cec2-cf34-43a7-bac7-c48a492ac57a" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-frame-tag/" target="_blank">frame</a></td>
<td>HTML Frames are used to divide the web browser window into multiple sections. Not supported in HTML5.</td>
<td>&lt;frame/&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/856360ac-a0d4-4f84-99d8-5b7e3d092154" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-frameset-tag/" target="_blank">frameset</a></td>
<td>The frameset element contains one or more frame elements. It is used to specify the number of rows and columns in a frameset with their pixel of spaces.</td>
<td>&lt;frameset cols = “pixels|%|*”&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/856360ac-a0d4-4f84-99d8-5b7e3d092154" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-head-tag/" target="_blank">head</a></td>
<td>The head tag in HTML is used to define the head portion of the document which contains information related to the document.</td>
<td>&lt;head&gt;…&lt;/head&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/142b2a94-1c99-4f1d-aa8d-a91c7bf30945" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-5-header-tag/" target="_blank">header</a></td>
<td>The header tag is used to contain the information related to the title and heading of the related content.</td>
<td>&lt;header&gt; …&lt;/header&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/416900f9-168c-482b-808f-34227d545506" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-heading/" target="_blank">heading</a></td>
<td>An HTML heading tag is used to define the headings of a page. These 6 heading elements are h1, h2, h3, h4, h5, and h6; with h1 being the highest level and h6 being the least.</td>
<td>&lt;h1&gt;Heading1&lt;/h1&gt;&lt;h2&gt;Heading2&lt;/h2&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/08de951d-a1a6-464d-8b2e-9101b8d6974f" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-hgroup-tag/" target="_blank">hgroup</a></td>
<td>The hgroup tag in HTML is used to wrap one or more heading elements from &lt;h1&gt; to &lt;h6&gt;, such as the headings and sub-headings.</td>
<td>&lt;hgroup&gt; … &lt;/hgroup&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/08a3f91f-b2e5-44a3-855c-3bb90d3cb9ae" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-hr-tag/" target="_blank">hr</a></td>
<td>The hr tag in HTML stands for horizontal rule and is used to insert a horizontal rule.</td>
<td>&lt;hr&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/aff72766-3d13-43ed-8f2b-42fcd96537e6" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-html-tag/" target="_blank">html</a></td>
<td>The html tag in HTML is used to define the root of HTML and XHTML documents.</td>
<td>&lt;html&gt; Contents &lt;/html&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/3760b6df-3bcb-4c35-a221-3857a0b05331" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-iframes/" target="_blank">Iframes</a></td>
<td>The iframe tag defines a rectangular region within the document in which the browser can display a separate document.</td>
<td>&lt;iframe src=”URL” title=”description”&gt;&lt;/iframe&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/2ec908b1-af76-4d23-b99e-07830f7cd340" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-images/" target="_blank">image</a></td>
<td>HTML Image, how to add the image in HTML. In earlier times, the web pages only contains textual content, which made them appear quite boring and uninteresting.</td>
<td>&lt;img src=”url” alt=”some_text” width=”” height=””&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/0aeb35b6-8a61-4886-9017-50f0fd8175ba" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-input-tag/" target="_blank">input</a></td>
<td>The input tag is used within &lt; form&gt; element to declare input controls that allow users to input data.</td>
<td>&lt;input type = “value” …. /&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/6884f9fa-b64c-4c8a-9cef-490adf449728" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-ins-tag/" target="_blank">ins</a></td>
<td>The ins tag is typically used to mark a range of text that has been added to the document.</td>
<td>&lt;ins&gt; Contents… &lt;/ins&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/787de785-5956-40ae-98fb-9c0c038c5828" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-isindex-tag/" target="_blank">isindex</a></td>
<td>The index tag is used to query any document through a text field.</td>
<td>&lt;isindex prompt=”search”&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/7e101a37-222f-466e-b9a7-efcf010e6f65" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-i-tag/" target="_blank">italic</a></td>
<td>This tag is generally used to display a technical term, phrase, the important word in a different language.</td>
<td>&lt;i&gt; Contents&lt;/i&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/aba6bcf1-8106-405f-a5db-55971c1c7196" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-kbd-tag/" target="_blank">kbd</a></td>
<td>The text enclosed within kbd tag is typically displayed in the browser’s default monospace font.</td>
<td>&lt;kbd&gt; text content … &lt;/kbd&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/fd6d3fb5-7024-45d9-b6fd-624df448dd0e" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-keygen-tag/" target="_blank">keygen</a></td>
<td>The keygen tag in HTML is used to specify a key-pair generator field in a form. When a form is submitted then two keys are generated, the private key and a public key.</td>
<td>&lt;keygen name = “name”&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/ee647689-73fd-46db-bb29-7473f0f6ac00" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-label-tag/" target="_blank">label</a></td>
<td>The label tag in HTML is used to provide a usability improvement for mouse users.</td>
<td>&lt;label&gt; form content… &lt;/label&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/8aec82cc-1582-45f3-bdc3-5c55588d5342" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-legend-tag/" target="_blank">legend</a></td>
<td>The legend tag is used to define the title for the child contents. The legend elements are the parent element.</td>
<td>&lt;legend&gt; Text &lt;/legend&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/ac2699e9-8260-4c77-81de-eb816ba637be" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-li-tag/" target="_blank">list</a></td>
<td>The list tag in HTML is used to define the list item in an HTML document. It is used within an Ordered List &lt;ol&gt; or Unordered List &lt;ul&gt;.</td>
<td>&lt;li&gt; List Items &lt;/li&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/10315c0e-506e-493f-9f09-ebba2021110f" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-main-tag/" target="_blank">main</a></td>
<td>The main tag is used to give the main information of a document. The content inside the &lt;main&gt; element should be unique for the document.</td>
<td>&lt;main&gt;Coontents&lt;/main&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/47713a7a-0f01-43a0-9c2f-4feded9b8c93" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html5-mark-tag/" target="_blank">mark</a></td>
<td>The mark tag in HTML is used to define the marked text. It is used to highlight the part of the text in a paragraph.</td>
<td>&lt;mark&gt; Contents… &lt;/mark&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/670592ce-96ed-423d-8dbf-79e8bbb000c0" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-marquee-tag/" target="_blank">marquee</a></td>
<td>The marquee tag in HTML is used to create scrolling text or images on a webpage. It scrolls either horizontally or vertically.</td>
<td>&lt;marquee&gt;Contents&lt;/marquee&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/c3ce613f-99de-43ee-b8d2-f2ff5d34c592" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-menuitem-tag/" target="_blank">menuitem</a></td>
<td>The menuitem tag is used to define a command or menu that the user can utilize from the popup item. Not supported in HTML5.</td>
<td>&lt;menuitem label=”” icon=”” type&gt; &lt;/menuitem&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/e557bf2b-d58c-414c-92ed-c08f0d309602" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-meta-tag/" target="_blank">meta</a></td>
<td>The meta tag is regularly used to give watchwords, portrayals, author data, and other metadata that might be utilized by the program to deliver the document accurately or in simple words, it provides important information about a document.</td>
<td>&lt;meta attribute-name=”value”&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/0d4d528e-2aca-40aa-b6b8-ddb84f504c2a" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html5-meter-tag/" target="_blank">meter</a></td>
<td>It is used to define the scale for measurement in a well-defined range and also supports a fractional value.</td>
<td>&lt;meter attributes…&gt; &lt;/meter&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/40df23f7-4da4-45b0-9bd5-7787e465d519" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-nav-tag/" target="_blank">nav</a></td>
<td>The nav tag is used for declaring the navigational section in HTML documents. Websites typically have sections dedicated to navigational links, which enables users to navigate the site.</td>
<td>&lt;nav&gt; Links… &lt;/nav&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/0fcec69a-4f9b-4f78-b298-7cb0d0582213" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-nobr-tag/" target="_blank">nobreak</a></td>
<td>The no break tag is used to create a single line text, that does not matter how long the statement is, this tag is used with &lt;wbr&gt; tag.</td>
<td>&lt;nobr&gt; Statement &lt;/nobr&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/d4b329c3-62db-451c-b8a4-b1f44d05e345" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-noembed-tag/" target="_blank">noembed</a></td>
<td>The noembed tag is used to show that the browser is not supported by &lt;embed&gt; tag.</td>
<td>&lt;noembed&gt; Element &lt;/noembed&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/b7d7d48c-3e55-4d53-9647-bfccef591b12" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-noscript-tag/" target="_blank">noscript</a></td>
<td>The noscript tag in HTML is used to display the text for those browsers which does not support the script tag or the browsers disable the script by the user.</td>
<td>&lt;noscript&gt; Contents… &lt;/noscript&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/9934890b-4c44-4e05-b040-6feba8a0c63b" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-object-tag/" target="_blank">object</a></td>
<td>The object tag is an HTML tag used to display multimedia like audio, videos, images, PDFs, and Flash on web pages.</td>
<td>&lt;object&gt;…&lt;/object&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/c2650a7f-67ac-4e8a-b059-9df42924dade" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-optgroup-tag/" target="_blank">optgroup</a></td>
<td>This tag is used to create a group of the same category options in a drop-down list.</td>
<td>&lt;optgroup&gt;…&lt;/optgroup&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/25c7c7d6-5ced-47a1-8488-74c666a54179" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-option-tag/" target="_blank">option</a></td>
<td>The option tag in HTML is used to choose an option from a Drop-Down menu.</td>
<td>&lt;option&gt; Contents… &lt;/option&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/d7e289a1-a7b0-4925-949c-d0525fb6426b" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-output-tag/" target="_blank">output</a></td>
<td>The output tag in HTML is used to represent the result of a calculation performed by the client-side script such as JavaScript.</td>
<td>&lt;output&gt; Results… &lt;/output&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/01fa63ff-0a35-4b55-9ed0-4346915ea1c7" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-paragraph/" target="_blank">paragraphs</a></td>
<td>The &lt;p&gt; tag in HTML defines a paragraph. These have both opening and closing tags.</td>
<td>&lt;p&gt; Content &lt;/p&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/f65a6826-6bd9-4f2a-9f14-c66f56eee3dc" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-param-tag/" target="_blank">param</a></td>
<td>The param tag in HTML is used to define a parameter for plug-ins which is associated with &lt;object&gt; element.</td>
<td>&lt;param name=”” value=””&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/9ec2e81f-f82c-46f3-a8fc-fab89dd824a0" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-phrase-tag/" target="_blank">phrase</a></td>
<td>In HTML, phrase tag is used to indicate the structural meaning of a block of text.</td>
<td>&lt;em&gt; Text Content &lt;/em&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/57aea2f7-20c9-41c5-a3b8-c4bff097625b" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-pre-tag/" target="_blank">pre</a></td>
<td>The pre tag in HTML is used to define the block of preformatted text which preserves the text spaces.</td>
<td>&lt;pre&gt; Contents… &lt;/pre&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/8a82b6f3-1a74-411a-8349-c09d8fef290b" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-5-progress-tag/" target="_blank">progress</a></td>
<td>It is used to represent the progress of a task. It is also defined how much work is done and how much is left to download a thing.</td>
<td>&lt;progress attributes…&gt; &lt;/progress&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/0e3a00dd-58fd-4824-80da-d99e4e07d051" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-q-tag/" target="_blank">q</a></td>
<td>The q tag is a standard quotation tag and is used for short quotations.</td>
<td>&lt;q&gt; Contents… &lt;/q&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/a2232354-3234-46b3-8a8e-553552588d62" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html5-rp-tag/" target="_blank">rp</a></td>
<td>The rp tag in HTML is used to provide parentheses around a ruby main text which defines the information.</td>
<td>&lt;rp&gt;[&lt;/rp&gt; Explaination… &lt;rp&gt;]&lt;/rp&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/14a826c4-6c7c-42c2-9bae-81acb8c036dd" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html5-rt-tag/" target="_blank">rt</a></td>
<td>The rt tag in HTML is used to define the explanation of the ruby annotation which is a small text, attached to the main text.</td>
<td>&lt;rt&gt; Explanation… &lt;/rt&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/652ad943-4cb5-47b8-b8fd-331b6a20085a" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html5-ruby-tag/" target="_blank">ruby</a></td>
<td>The ruby tag in HTML is used to specify the ruby annotation which is a small text, attached with the main text to specify the meaning of the main text.</td>
<td>&lt;ruby attributes&gt; Contents… &lt;/ruby&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/bde0ad4c-6c41-4536-83d7-bcf4ae2094d2" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-s-tag/" target="_blank">s</a></td>
<td>This tag is used to specify that the text content is no longer correct or accurate. This tag is similar but slightly different from &lt;del&gt; tag.</td>
<td>&lt;s&gt; Contents… &lt;/s&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/dfb08f74-a158-4e56-a78f-268f62ff2051" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-samp-tag/" target="_blank">samp</a></td>
<td>It is a phrase tag used to define the sample output text from a computer program.</td>
<td>&lt;samp&gt; Contents… &lt;/samp&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/3032e433-f18d-4d43-9235-af1f13571b10" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-script-tag/" target="_blank">script</a></td>
<td>The script tag in HTML is used to define the client-side script.</td>
<td>&lt;script&gt; Script Contents… &lt;/script&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/7885d953-3e8d-4a5e-b2e1-6c833e56ad71" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-section-tag/" target="_blank">section</a></td>
<td>Section tag defines the section of documents such as chapters, headers, footers, or any other sections.</td>
<td>&lt;section&gt; Section Contents &lt;/section&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/d44504ce-d72c-444f-94d6-84501f6a40ab" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-small-tag/" target="_blank">small</a></td>
<td>The small tag in HTML is used to set small font sizes. It decreases the font size by one size (from medium to small, from x-large to large).</td>
<td>&lt;small&gt; Contents… &lt;/small&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/b6fe48cc-5e42-4547-888c-0f2278f76b15" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-source-tag/" target="_blank">source</a></td>
<td>The source tag in HTML is used to attach multimedia files like audio, video, and pictures.</td>
<td>&lt;source src=”” type=””&gt;&nbsp;&lt;/source&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/69384add-947c-473a-a020-2d6da2f7fec6" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-spacer-tag/" target="_blank">spacer</a></td>
<td>The spacer tag is used to create some white space. Not-supporte in HTML5 .</td>
<td>&lt;spacer type=”” size=””&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/b070be34-acb8-4f4d-b999-fc3ed0f57c33" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/span-tag-html/" target="_blank">span</a></td>
<td>The HTML span element is a generic inline container for inline elements and content.</td>
<td>&lt;span class=””&gt;Some Text&lt;/span&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/b52a2a8d-ac9a-4ffa-8dfc-740cee344f4c" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-strike-tag/" target="_blank">strike</a></td>
<td>HTML strike tag, along with understanding its implementation through the example. The &lt;strike&gt; tag defines a strike or line through Text.</td>
<td>&lt;strike&gt; Contents &lt;/strike&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/69287866-5f92-4f9f-b72e-86b6b1327859" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-strong-tag/" target="_blank">strong</a></td>
<td>The strong tag in HTML is the parsed tag and is used to show the importance of the text. Make that text bold.</td>
<td>&lt;strong&gt; Contents… &lt;/strong&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/9d3d2a46-6629-4101-9445-bb7a2e25f4aa" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-style-tag/" target="_blank">style</a></td>
<td>The style tag in HTML helps us to design the web page.</td>
<td>&lt;tagname style=”property:value;”&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/25f3c0cd-3134-48d1-9566-f5016c6c752c" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-subscript-superscript-tags/" target="_blank">sub and sup Tags</a></td>
<td>
<ul>
<li>The sub-tag is used to add a subscript text to the HTML document.</li>
<li>The &lt;sup&gt; tag is used to add superscript text to the HTML document.</li>
</ul>
</td>
<td>&lt;sub&gt;subscript text&lt;/sub&gt;&lt;sup&gt;superscript text&lt;/sup&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/0ce61fab-6cf4-492a-b68f-bb208d3988d3" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-5-summary-tag/" target="_blank">summary</a></td>
<td>The &lt;summary&gt; tag in HTML is used to define a summary for the &lt;details&gt; element.</td>
<td>&lt;summary&gt; Content &lt;/summary&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/8faf02ea-d616-4be2-84f5-d83a5236c14b" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-svg-basics/" target="_blank">svg</a></td>
<td>HTML SVG Basics, &amp; their implementation through the examples. SVG stands for Scalable Vector Graphics.</td>
<td>&lt;svg height=”” width=””&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/02108d7c-b358-478d-8283-15e14b3f9935" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-tables/" target="_blank">table</a></td>
<td>HTML Table, various ways to implement it, &amp; will also understand its usage through the examples. HTML Table is an arrangement of data in rows and columns, or possibly in a more complex structure.</td>
<td>&lt;table&gt;… &lt;/table&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/d02cc077-0f6f-49a6-bf46-b148652e5bf5" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-tbody-tag/" target="_blank">tbody</a></td>
<td>The tbody tag in HTML is used to make a group of the same type of content of the body element.</td>
<td>&lt;tbody&gt; // Table contents &lt;/tbody&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/eb38b287-673c-43b9-b7b4-9580bf53d554" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-td-tag/" target="_blank">td</a></td>
<td>The table data tag is used to define a standard cell in an HTML table.</td>
<td>&lt;td&gt;……..&lt;/td&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/ec62e4b4-98a6-4b67-a9f7-373da09f103e" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-template-tag/" target="_blank">template</a></td>
<td>The template tag in HTML is used to store the HTML code fragments, which can be cloned and inserted in an HTML document.</td>
<td>&lt;template&gt; Contents &lt;/template&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/1270335f-ced9-45c7-b890-d73d8aac9101" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-tfoot-tag/" target="_blank">tfoot</a></td>
<td>This tag is used in HTML table with header and body which is known as “thead” and “tbody”.</td>
<td>&lt;tfoot&gt; // Table footer contents… &lt;/tfoot&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/45810107-c722-4163-9266-23a98bb4279d" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-th-tag/" target="_blank">th</a></td>
<td>The table header tag in HTML is used to set the header cell of a table. Two types of cells in the HTML table Header &amp; Standard.</td>
<td>&lt;th&gt; Contents… &lt;/th&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/55f2314c-df64-4164-95a7-ee5d295fb763" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-thead-tag/" target="_blank">thead</a></td>
<td>This tag is used in HTML tables as head and body which are known as thead and tbody.</td>
<td>&lt;thead&gt;Table head Contents…&lt;/thead&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/d8c49b3a-8c92-44a3-ba2f-23dad08a86ba" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-time-tag/" target="_blank">time</a></td>
<td>The time tag is used to display the human-readable date/time. It can also be used to encode dates and times in a machine-readable form.</td>
<td>&lt;time attribute&gt; Time… &lt;/time&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/d7dd20f2-232f-4b3e-86bf-70d55c36d401" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-title-tag/" target="_blank">title</a></td>
<td>The title tag in HTML is used to define the title of HTML document. It sets the title in the browser toolbar.</td>
<td>&lt;title&gt; Title name &lt;/title&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/78bd6104-c40d-4235-9d71-a5c9ae3364c1" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-tr-tag/" target="_blank">tr</a></td>
<td>The table row tag is used to define a row in an HTML table. The &lt;tr&gt; element contains multiple &lt;th&gt; or &lt;td&gt; elements.</td>
<td>&lt;tr&gt;…..&lt;/tr&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/6b86eba5-dd3c-43fb-813b-2f91075845bc" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-track-tag/" target="_blank">track</a></td>
<td>The tracking tag specifies text tracks for media components audio and video.</td>
<td>&lt;track attribute&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/f0604f54-8e19-4531-92c1-d78820f589ca" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-tt-tag/" target="_blank">tt</a></td>
<td>The tt tag is the abbreviation of teletype text. This tag is depreciated from HTML 5. It was used for marking Keyboard input.</td>
<td>&lt;tt&gt; Contents… &lt;/tt&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/114c14e3-1ddf-4292-8ad2-88b8b24d48de" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-u-tag/" target="_blank">underline</a></td>
<td>The underline tag in HTML stands for underline, and it’s used to underline the text enclosed within the &lt;u&gt; tag.</td>
<td>&lt;u&gt; Contents… &lt;/u&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/f5ec3fbc-48b2-44dc-b12c-aee567691954" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-var-tag/" target="_blank">var</a></td>
<td>It is a phrase tag used to specify the variable in a mathematical equation or in a computer program.</td>
<td>&lt;var&gt; Contents… &lt;/var&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/8bf8223b-6e2e-4713-8fd6-70061d04cfea" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html5-video/" target="_blank">video</a></td>
<td>HTML5 Video, along with knowing the different ways to add the videos to the HTML page.</td>
<td>&lt;video src=”” controls&gt; &lt;/video&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/469d139f-22c3-4c68-b6c4-e29e0e7587cb" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-5-wbr-tag/" target="_blank">wbr</a></td>
<td>The wbr tag is used to define the position within the text which is treated as a line break by the browser.</td>
<td>&lt;wbr&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/07f80b12-5710-4943-8a1a-2fdd7ff367ab" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
<tr>
<td><a href="https://www.geeksforgeeks.org/html-xmp-tag/" target="_blank">xmp</a></td>
<td>The XMP tag is used to create any content in letter format.</td>
<td>&lt;xmp&gt; statement &lt;/xmp&gt;</td>
<td>
<div class="practiceLinkDiv-banner">
<div class="redirect"><a href="https://ide.geeksforgeeks.org/tryit.php/4fcee4c0-cbed-4629-b213-bc6b6cf374d2" target="_blank">Try</a></div>
<p></p></div>
</td>
</tr>
</tbody>
</table>


