import React, { useState } from "react";

type UserType = {
    name: string;
    age: number;
};

const _5_use_state = () => {
    const [user, setUser] = useState<UserType | null>(null);
    return <div>_5_use_state</div>;
};

export default _5_use_state;
