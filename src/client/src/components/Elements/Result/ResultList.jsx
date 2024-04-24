import { duration } from "@mui/material";
import Aos from "aos";
import "aos/dist/aos.css"
import React, { useEffect } from "react";

export default function ResultList({ dataResults }) {
    useEffect(() => {
        Aos.init({duration: 1000});
    }, []);

    return (
        <div className="d-flex justify-content-center mt-5" style={{ width: '90%' }}>
            {dataResults && dataResults.length > 0 && !dataResults.some(item => item.title === null || item.url === null) ? (
                <div className="container-result text-white">
                    {dataResults.map((item) => (
                        <div key={item.id} data-aos="fade-right" className="box-result flex flex-col justify-content-center mt-10 b-10 border-white border-2 p-3" style={{ width: '30%' }}>
                            <h2 className="mx-auto" style={{ color: 'white', fontSize: '20px', display: 'inline-block' }}>PATH {item.id}</h2>
                            <ul>
                                {item.title.map((title, index) => (
                                    <div key={index} className="flex flex-col justify-content-center mt-5" style={{ width: '100%' }}>
                                        <li>
                                            <a className="mx-auto hover:text-sky-400 duration-200" style={{ fontSize: '20px', display: 'inline-block' }} href={item.url[index]}>{title}</a>
                                        </li>
                                    </div>
                                ))}
                            </ul>
                        </div>
                    ))}
                </div>
            ) : (
                <div>
                    <h1 style={{ color: 'white', fontSize: '20px', display: 'inline-block' }}>NO RESULT</h1>
                </div>
            )}
        </div>
    );
}
