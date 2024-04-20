"use client";
import React, { useState } from "react";

function Algo() {
    const [value, setValue] = useState('')

    const options = [
        {label: "BFS", value: 1},
        {label: "IDS", value: 2},
    ]

    function handleSelect(event) {
        setValue(event.target.value)
    }

    return (
        <div className="d-flex justify-content-center mt-5">
            <div className="w-50 p-3 border rounded">
                <h3>ALGO</h3>
                <select className="form-select" onChange={handleSelect}>
                    {options.map(option => (
                        <option key={option.value} value={option.value}>{option.label}</option>
                    ))}
                </select>

            </div>
        </div>
    );
}

export default Algo;