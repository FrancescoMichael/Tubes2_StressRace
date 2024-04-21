import React from "react";

export default function ResultList({ dataResults }) {
    return (
        <div className="d-flex justify-content-center mt-5" style={{ width: '90%' }}>
            {dataResults.length > 0 ? (
                <div>
                    <div className="container-result text-white">
                        {dataResults.map((path, i) => (
                            <div key={i} className="box-result flex flex-col justify-content-center mt-10 b-10 border-white border-2 p-3" style={{ width: '30%' }}>
                                <h2 className="mx-auto" style={{ color: 'white', fontSize: '20px', display: 'inline-block' }}>PATH {i + 1}</h2>
                                <ul>
                                    {path.map((link, j) => (
                                        <div key={j} className="flex flex-col justify-content-center mt-5" style={{ width: '100%' }}>
                                            <li>
                                                <a className="mx-auto hover:text-sky-400 duration-200" style={{ fontSize: '20px', display: 'inline-block' }} href={link.value}>{link.label}</a>
                                            </li>
                                        </div>
                                    ))}
                                </ul>
                            </div>
                        ))}
                    </div>
                </div>
            ) : (
                <div>
                    <h1 style={{ color: 'white', fontSize: '20px', display: 'inline-block' }}>NO RESULT</h1>
                </div>
            )}
        </div>
    );
}
