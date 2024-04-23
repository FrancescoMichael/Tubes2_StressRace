"use client";
import React , { useState }from 'react';
import Input from './input'; 
import SwapHorizIcon from '@mui/icons-material/SwapHoriz';
import Algo from './algo';
import { WindowSharp } from '@mui/icons-material';

function InputForm({ isLoading, setIsLoading, startTime, setStartTime}) {
    const [notification, setNotification] = useState("")
    const [algorithm, setAlgorithm] = useState(1)

    const [searchStart, setSearchStart] = useState("")
    const [urlStart, setURLStart] = useState("")

    const [searchTarget, setSearchTarget] = useState("")
    const [urlTarget, setURLTarget] = useState("")

    

    const handleSubmit = (event) => {
        event.preventDefault();
        const formData = new FormData(event.target);
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
            setNotification(" Please input your data correctly!");
            setTimeout(() => {
                setNotification("");
            }, 2000);
        } else {
            setIsLoading(true);
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
            setStartTime(performance.now());
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
                        <button type="swap">
                            <SwapHorizIcon style={{fontSize: '50px'}} onClick = {handleSwap}/>
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
                        <div className="fixed inset-0 bg-black bg-opacity-30 backdrop-blur-sm flex justify-center items-center text-red-500">
                            <div className="bg-red-100 border border-red-400 text-red-700 px-6 py-4 rounded relative text-lg" role="alert">
                                <strong class="font-bold">Holy smokes!</strong>
                                <span class="block sm:inline">{notification}</span>

                            </div>
                        </div>
                )}
                </div>
            </form>
        </>
    );
}

export default InputForm;
