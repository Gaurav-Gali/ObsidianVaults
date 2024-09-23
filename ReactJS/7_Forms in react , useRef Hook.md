![Source](https://youtu.be/eILUmCJhl64?t=29058)

1. State Management : Each input's state is stored in the component's state.
2. Handling Changes : Use **onChange** to detect input changes.
3. Submissions : Use **onSubmit** to handle submissions , to prevent the default behaviour of reloading the page when submitting the form , use ==event.preventDefault()==
4. Validation : Use custom validation or third-party libraries.

<hr>

### useRef Hook
1. This hook is used to access to DOM elements and retain mutable values without re-renders.
2. Used with the ref attribute for direct DOM interactions.
3. Can hold previous state or prop values.
4. It is not limited to DOM references hence it can hold any value.
5. Refs can be passed as props as well.

```js
import {useRef} from 'react';

// 1.creating a useRef hook
const refObject = useRef(0);

// 2.accessing the current value stored
console.log(refObejct.current);

// 3.using the ref attribute in inputs
<input ref={refObject} />
```