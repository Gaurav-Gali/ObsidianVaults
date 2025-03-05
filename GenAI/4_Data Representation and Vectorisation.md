![Soruce](https://youtu.be/F0GQ0l2NfHA?t=5175)

1. Feature extractions from text/images.
	1. It is the conversion of textual data to numerical or vector data.
2. Why we need it ?
	1. The models only take in numerical or vector data.
	2. Hence converting of raw data to text is crucial.
3. Why it is so difficult ?
	1. It involves a lot of processing
4. What is the core idea ?
5. Some techniques.

## Methods to vectorise raw data

### One Hot Encoding
Let's consider the following text :

People like watching movies.
They  like watching sports.

Here there are 4 unique words (People, movies, They, sports).
==Hence N = 4==

**Step 1 :** Write each row of  a paragraph as new line in a table.

<table>
	<tr>
		<td>D1</td>
		<td>People like watching movies.</td>
	</tr>
	<tr>
		<td>D2</td>
		<td>They  like watching sports.</td>
	</tr>
</table>

**Step 2 :** Write down a table with the unique words and represent the given text as binary.
The rules are such that, if the unique word is present in that line, then mark 1 else 0.

For D1

<table>
	<tr>
		<td>People</td>
		<td>movies</td>
		<td>They</td>
		<td>sports</td>
	</tr>
	<tr>
		<td>1</td>
		<td>0</td>
		<td>0</td>
		<td>0</td>
	</tr>
	<tr>
		<td>0</td>
		<td>1</td>
		<td>0</td>
		<td>0</td>
	</tr>
</table>

For D2

<table>
	<tr>
		<td>People</td>
		<td>movies</td>
		<td>They</td>
		<td>sports</td>
	</tr>
	<tr>
		<td>0</td>
		<td>0</td>
		<td>1</td>
		<td>0</td>
	</tr>
	<tr>
		<td>0</td>
		<td>0</td>
		<td>0</td>
		<td>1</td>
	</tr>
</table>

**Step 3 :** Create a 2D vector for each Data

```py
D1 = [
	[1,0,0,0],
	[0,1,0,0]
]

D2 = [
	[0,0,1,0],
	[0,1,0,1]
]
```

**Step 4 :** Using the generated 2D vectors, create a neural network with the input size as N which is the number of unique words.

**The drawbacks of One Hot Encoding** :
1. Sparsity, which is the unwanted number of 0's that increase the complexity and size of the vector.
2. No Fixed Size
3. OOV, which is Out Of Vocabolary issue.
4. Not capturing semantic meaning.

Hence this method is not widely used.

<hr>
### BOW (Bag Of Words)
Let's consider the following text :

People like watching movies.
They  like watching sports.

Here there are 4 unique words (People, movies, They, sports).
==Hence N = 4==

<table>
	<tr>
		<td>D1</td>
		<td>People like watching movies.</td>
	</tr>
	<tr>
		<td>D2</td>
		<td>They  like watching sports.</td>
	</tr>
</table>

This method works with the count of words.

**Step 1 :** Construct a table with the unique words and its count per line

<table>
	<tr>
		<td>People</td>
		<td>movies</td>
		<td>They</td>
		<td>sports</td>
	</tr>
	<tr>
		<td>1</td>
		<td>1</td>
		<td>0</td>
		<td>0</td>
	</tr>
	<tr>
		<td>0</td>
		<td>0</td>
		<td>1</td>
		<td>1</td>
</table>


==**Note :** This method is suitable for sentimental analysis.==

<hr>

#### Some modern techniques for vectorisation
1. TFIDF
	1. This is a statistical way of vectorisation.
2. Word2Vec
	1. This is a Deep Learning way of vectorisation.
3. Transformer
	1. Extracting features using transformer models.
	2. Used in LLM's

