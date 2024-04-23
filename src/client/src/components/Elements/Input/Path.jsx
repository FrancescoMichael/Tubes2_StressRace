"use client"
import React, { useState } from "react";
import {roboto} from "@/app/ui/font"

function Path(props) {
    const { setPathChoice } = props;
    const [value, setValue] = useState(1)

    const options = [
        {label: "Single Path", value: 1},
        {label: "Multiple Path", value: 2},
    ]

    function handleSelect(event) {
        setValue(event.target.value)
        setPathChoice(event.target.value)
    }

    return (
        <div className= "d-flex justify-content-center mt-5 mb-5">
            <div style={{ color: 'white', fontSize: '20px', display: 'inline-block'}}>Select Your Number of Path</div>
            <div className={ `${roboto.className} flex flex-wrap w-[100%] items-center text-center justify-between bg-[transparent] md:p-0 mt-2`}>
                
                <select className="form-select flex flex-wrap w-[100%] items-center justify-between rounded-xl border-gray-300 p-2" onChange={handleSelect}>
                    {options.map(option => (
                        <option key={option.value} value={option.value}>{option.label}</option>
                    ))}
                </select>
            </div>
        </div>
    );
}

export default Path;