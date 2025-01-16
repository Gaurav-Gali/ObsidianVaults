![Source](https://youtu.be/O0UGlA1YVUI)

**Note** : The fetch() function in NextJS15 is slightly modified.

1. The data is only only fetched once and the result is cached. This is to avoid making multiple api calls and minimise the cost and enhance performance.
2. We can override this feature by the following way
```ts
fetch("https://...", {cache: 'no-store'});
```

Since by default the fetch() function only fetches the data from the api only once, it is not required to use the useEffect() hook and instead it can be done far more easily.
```ts
const response = await fetch("https://...");
const data = await response.json();
```

**Note** : ==The data will not be cached when used inside a server action or a route handler.==
#### Incremental Static Regeneration (ISR)
- The problem in the above case is that the data from the api can be fetched with once and be cached or keep requesting it each time the page reloads.
- To overcome this we make use of ISR. It is simply revalidating the request in a given interval or when demanded to.

Implementation
```ts
const response = await fetch("https://...", {next:{revalidate: 3600}})
const data = await response.json();
```
Here the data will be fetched and cached every 3600 seconds or 1 hour.

**Note** : ==The page will be built as a static page if the data has been cached. To prevent this default behaviour we can do the following==
```ts
const response = await fetch("https://...", {next:{revalidate: 3600}})
const data = await response.json();

export const dynamic = 'force-dynamic'; //To prevent the static page generation.
```