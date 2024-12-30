![Source](https://youtu.be/CzRQ9mnmh44?t=30743)

Basic Boilerplate
```dart
// This package allows to start working with flutter
import 'package:flutter/material.dart';
```

```dart
import 'package:flutter/material.dart';

void main() {
	// Basic hello world
	runApp(
		Text(
			"Hello World",
			textDirection: TextDirection.ltr,
		)
	);	
}
```

### Stateless widgets
```dart
import 'package:flutter/material.dart';

void main() {
	runApp(const MyApp()); //Takes a widget as an argument
}

class MyApp extends StatelessWidget {
	consy MyApp({super.key});
	
	@override
	Widget build(BuildContext context) {
		return const Text("Hello Flutter", textDirection: TextDirection.ltr);
	}
}
```
1. Make sure to override the build function in order to return useful widgets
2. Also pass the key of the abstract class via the constructor using  **super.key**
	1. A key helps flutter different different widgets in the app.
3. While calling the widget, making it a constant will make sure it doesn't re-render when not required.
4. In general declare widgets as constants in order to improve efficiency.

### Using MaterialApp
```dart
import 'package:flutter/material.dart';

void main() {
	runApp(const MyApp()); //Takes a widget as an argument
}

class MyApp extends StatelessWidget {
	const MyApp({super.key});
	
	@override
	Widget build(BuildContext context) {
		return const MaterialApp();
	}
}
```
1. The main advantage of using a design pattern like Material or Cupertino is that it allows uniformity in the UI.
2. Performing navigations and theming will be easier. 