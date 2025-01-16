![Source](https://youtu.be/F0GQ0l2NfHA?t=1460)

**Note** : ==Generative AI pipeline is a set of steps followed to build an end to end GenAI software==

# Pipeline

### Data Acquisition
1. Check for available data(text, c.s.v, PDF, xlsx etc).
2. If data is not available, then check for other data(DB, internet, API, web scraping).
3. When data is not available, then collect your own data.
	1. The main drawback in creating your own data is that the dataset will be less.
	2. Hence, you can make use of [[1_General#LLMs|LLMs]].
4. If you have less data then the following can be done.
	1. ==Replace data with synonyms==
		1. Ex : I like apples -> I like oranges.
	2. ==Bigram Flip==
		1. This is a process of changing the data without changing it's context.
		2. Ex: Hello, I'm Gaurav -> My name is Gaurav.
	3. ==Back Translation==
		1. Translating text from target to source, this will modify the data without changing the context and also provide an improved result.
	4. ==Adding Additional Data==
		1. Ex : Hello, I am Gaurav -> Hello, I am Gaurav and I am a Student.

#### Data Preparation
1. Data Cleanup
	1. Say, the data is scrapped from the web, then we would have to remove the HTML tags and do some spelling corrections if required etc.
	2. Basic Preprocessing
		1. Tokenisation is the process of breaking down large text data to smaller units called tokens.
			1. Sentence level tokenisation 
				1. Say, data="Hello guys I'm Tom, I'm happy to introduce you to my new cat jerry"
				2. This will be broken down to \["Hello guys I'm Tom", "I'm happy to introduce you to my new cat jerry"]
			2. Word level tokenisation
				1. Say, data="Good morning all"
				2. This will beroken down to \["Good", "morning", "all"]
	3. Optional Preprocessing
		1. Stop word removal
		2. Steaming(lesser used method)
			1. This process involes connecting different verb forms to a root representation
			2. Say verbs are, \["play", "played", "plays"] -> root="play" -> root="sports"
		3. Lamatization(frequently used method)
		4. Punctuation removal
		5. Lower case
			1. Convert the given data to lowercase to prevent duplication of data.
		6. Language detection
	4. Advance Preprocessing
		1. Parts of speech tagging
		2. Parsing
		3. Coreference resolution
		![[parsing.png]]


#### Feature Engineering
1. Text Vectorisation is the process of converting raw text into mathematical or vector data to feed it to the models. This can be done using the following ways.
	1. TFIDF
	2. Bag Of Word
	3. Word to Vec
	4. One Hot Encoding
	5. Transformer models


#### Modeling
1. Choosing either open-source models or paid models to work with.
	1. Open Source Models
		1. These models can be used free of cost but comes with some infrastructural difficulties.
	2. Paid Models
		1. These models need to be paid based on token usage and can be used with ease since a lot of the load bearing will be taken care on their own servers.

#### Evaluation
1. Intrinsic Evaluation
2. Extrinsic Evaluation

#### After the development of the model
- ##### Deployment 
- ##### Monitoring and model updating
