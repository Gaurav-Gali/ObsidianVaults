import React from "react";

const _4_events = () => {
    const handleInput = (e: React.ChangeEvent<HTMLInputElement>) => {
        console.log(e.target.value);
    };

    const handleButton = (e: React.MouseEvent<HTMLButtonElement>) => {
        e.preventDefault();
        console.log("Button clicked");
    };
    return (
        <div>
            <input type="text" onChange={(e) => handleInput(e)} />
            <button type="submit" onClick={(e) => handleButton(e)}>Submit</button>
        </div>
    );
};

export default _4_events;
