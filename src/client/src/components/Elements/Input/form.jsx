"use client";
import React from 'react';
import Input from './input'; 
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
            <div className="pt-10">
                <Input type="text" placeholder="start page" name="starting-page"></Input>
            </div>
            <div>
                <Input type="text" placeholder="target page" name="target-page"></Input>
            </div>
            <button type="submit" className="h-12 px-8 font-semibold rounded bg-[#31363F] text-gray-500">
                Search!
            </button>
        </form>
    );
}

export default InputForm;
