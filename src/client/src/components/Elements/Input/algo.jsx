"use client";
import React, { useState } from "react";
import {roboto} from "@/app/ui/font"

function Algo() {
    const [value, setValue] = useState(1)

    const options = [
        {label: "BFS", value: 1},
        {label: "IDS", value: 2},
    ]

    function handleSelect(event) {
        setValue(event.target.value)
    }

    return (
        <div className= "d-flex justify-content-center mt-5">
            <div style={{ color: 'white', fontSize: '20px', display: 'inline-block'}}>SELECT YOUR ALGORITHM</div>
            <div className={ `${roboto.className} flex flex-wrap w-[100%] items-center text-center justify-between bg-[transparent] md:p-0`}>
                
                <select className="form-select flex flex-wrap w-[100%] items-center justify-between" onChange={handleSelect}>
                    {options.map(option => (
                        <option key={option.value} value={option.value}>{option.label}</option>
                    ))}
                </select>
                <p>{value}</p>
            </div>
        </div>
    );
}

export default Algo;