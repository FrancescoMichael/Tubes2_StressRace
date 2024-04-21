"use client";
import React , { useState }from 'react';
import Input from './input'; 
import SwapHorizIcon from '@mui/icons-material/SwapHoriz';
import Algo from './algo';

function InputForm() {
    const [notification, setNotification] = useState("")
    const [algorithm, setAlgorithm] = useState(1)

    const [searchStart, setSearchStart] = useState("")
    const [urlStart, setURLStart] = useState("")

    const [searchTarget, setSearchTarget] = useState("")
    const [urlTarget, setURLTarget] = useState("")

    const handleSubmit = (event) => {
        event.preventDefault();
        const formData = new FormData(event.target);
        // const startingPage = formData.get('starting-page');
        // const targetPage = formData.get('target-page');
    };

    const handleSwap = () => {
        let tempTitle = searchStart;
        setSearchStart(searchTarget)
        setSearchTarget(tempTitle);

        let tempURL = urlStart;
        setURLStart(urlTarget)
        setURLTarget(tempURL);
    }

    const handleSearch = async () => {
        if(searchTarget === "" || searchStart === "") {
            console.log(notification);
            setNotification("Please input your data correctly!");
        } else {
            setNotification("");
            // console.log(`Start URL : ${urlStart}\nEnd URL : ${urlTarget}\nAlgoritma : ${algorithm}`)
            const algorithmString = algorithm.toString();
            
            // Kirim data ke backend
            await fetch('http://localhost:8080/api/search', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    urlStart,
                    urlTarget,
                    algorithm: algorithmString,
                }),
                credentials: 'include',
            })
            .then(response => {
                if (!response.ok) {
                    throw new Error('Network response was not ok');
                }
                return response.json();
            })
            .then(data => {
                console.log('Response from backend:', data);
            })
            .catch(error => {
                console.error('There was a problem with your fetch operation:', error);
            });
        }
    }

    return (
        <>
            <Algo setAlgorithmChoice = {setAlgorithm}/>
            <form onSubmit={handleSubmit} id="form">
                <div className="flex container-search">
                    <div className="box">
                        <Input type="text" placeholder="Start page..." name="starting-page" className="text-white" setInputSearch={setSearchStart} setURLSearch = {setURLStart} value={searchStart}></Input>
                    </div>
                    <div className="swap">
                        <button type ="swap">
                            <SwapHorizIcon style={{ color: 'white', fontSize: '50px'}} onClick = {handleSwap}/>
                        </button>
                    </div>
                    <div className="box">
                        <Input type="text" placeholder="Target page..." name="target-page" className="text-white" setInputSearch={setSearchTarget} setURLSearch = {setURLTarget} value={searchTarget}></Input>
                    </div>
                </div>

                <div className="flex flex-col items-center justify-center">
                    <button type="submit" className="glow-on-hover bg-white text-black px-4 py-2 text-xl rounded-xl font-medium" onClick = {handleSearch}>
                    {/* <button type="submit" className="bg-gradient-to-r from-blue-500 to-red-500 text-white px-4 py-2 text-xl rounded-xl font-medium focus:ring ring-black ring-opacity-10 gradient element-to-rotate" onClick = {handleSearch}> */}
                        Search!
                    </button>
                    {notification && (
                        <div className="text-red-500">{notification}</div>
                    )}
                </div>
            </form>
        </>
    );
}

export default InputForm;
