![Source](https://youtu.be/Zq5fmkH0T78?t=5571)

**Note** : ==Never use a CDN for fonts , it would take time to load and cause problems with the UX.
Instead download and ship the font directly.==

#### Using Google fonts in next/font
```tsx
import type { Metadata } from "next";
import "./globals.css";

import { Work_Sans } from "next/font/google";

const workSans = Work_Sans({
	weight: "400",
	subsets: ["latin"],
	display: "swap",
});

export const metadata: Metadata = {
	title: "StartNow",
	description: "A platform to explore startups and new business ideas",
};

export default function RootLayout({
	children,
}: Readonly<{
	children: React.ReactNode;
}>) {

	return (
		<html lang="en">
			<body className={`${workSans.className}`}>{children}</body>
		</html>
	);

}
```

