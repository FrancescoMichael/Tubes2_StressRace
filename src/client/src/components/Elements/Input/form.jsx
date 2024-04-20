"use client";
import React from 'react';
import Input from './input'; 
import SwapHorizIcon from '@mui/icons-material/SwapHoriz';

function InputForm() {
    const handleSubmit = (event) => {
        event.preventDefault();
        const formData = new FormData(event.target);
        const startingPage = formData.get('starting-page');
        const targetPage = formData.get('target-page');
        console.log('Starting Page:', startingPage, 'Target Page:', targetPage);
    };

    return (
        <form onSubmit={handleSubmit} className="space-y-4" id="form">
            <div className="container-search">
                <div className="box">
                    <Input type="text" placeholder="Start page..." name="starting-page"></Input>
                </div>
                <div className="swap">
                    <button type ="swap">
                        <SwapHorizIcon fontSize="large" style={{ color: 'white' }} />
                    </button>
                </div>
                <div className="box">
                    <Input type="text" placeholder="Target page..." name="target-page"></Input>
                </div>
            </div>


            <div class="flex flex-col items-center justify-center">
                <button type="submit" className="bg-gradient-to-r from-blue-500 to-red-500 text-white px-4 py-2 text-xl rounded font-medium focus:ring ring-black ring-opacity-10 gradient element-to-rotate">
                    Search!
                </button>
            </div>
        </form>
    );
}

export default InputForm;
