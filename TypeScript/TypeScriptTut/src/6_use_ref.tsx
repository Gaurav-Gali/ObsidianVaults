import React, { useEffect, useRef } from "react";

const _6_use_ref = () => {
    const inputRef = useRef<HTMLInputElement>(null);
    // Focusses on the input field when the component mounts.
    useEffect(() => {
        inputRef.current?.focus();
    }, []);
    
    return (
        <div>
            <input type="text" ref={inputRef} />
        </div>
    );
};

export default _6_use_ref;
