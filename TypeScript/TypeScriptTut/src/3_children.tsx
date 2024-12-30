import React from "react";

const _3_children = (props: { children: React.ReactNode }) => {
    return (
        <div>
            <h1>Parent</h1>
            {props.children}
            <h2>Child 3</h2>
        </div>
    );
};

export default _3_children;
