"use client"
import React, { useState } from "react";
import ResultList from "./ResultList"
import Pagination from "./Pagination"

export default function Result() {
    const data = [
        {
          "id": "1",
          "title": [
            "Hampi",
            "Hampi (town)",
            "Hampi Express",
            'Michael Jordan'
          ],
          "url": [
            "https://en.wikipedia.org/wiki/Hampi",
            "https://en.wikipedia.org/wiki/Hampi_(town)",
            "https://en.wikipedia.org/wiki/Hampi_Express",
            'https://en.wikipedia.org/wiki/Michael_Jordan'
          ]
        },
        {
          "id": "1",
          "title": [
            "Hampi",
            "Hampi (town)",
            "Hampi Express",
            'Michael Jordan'
          ],
          "url": [
            "https://en.wikipedia.org/wiki/Hampi",
            "https://en.wikipedia.org/wiki/Hampi_(town)",
            "https://en.wikipedia.org/wiki/Hampi_Express",
            'https://en.wikipedia.org/wiki/Michael_Jordan'
          ]
        },
        {
          "id": "1",
          "title": [
            "Hampi",
            "Hampi (town)",
            "Hampi Express",
            'Michael Jordan'
          ],
          "url": [
            "https://en.wikipedia.org/wiki/Hampi",
            "https://en.wikipedia.org/wiki/Hampi_(town)",
            "https://en.wikipedia.org/wiki/Hampi_Express",
            'https://en.wikipedia.org/wiki/Michael_Jordan'
          ]
        },
        {
          "id": "1",
          "title": [
            "Hampi",
            "Hampi (town)",
            "Hampi Express",
            'Michael Jordan'
          ],
          "url": [
            "https://en.wikipedia.org/wiki/Hampi",
            "https://en.wikipedia.org/wiki/Hampi_(town)",
            "https://en.wikipedia.org/wiki/Hampi_Express",
            'https://en.wikipedia.org/wiki/Michael_Jordan'
          ]
        },
        {
          "id": "1",
          "title": [
            "Hampi",
            "Hampi (town)",
            "Hampi Express",
            'Michael Jordan'
          ],
          "url": [
            "https://en.wikipedia.org/wiki/Hampi",
            "https://en.wikipedia.org/wiki/Hampi_(town)",
            "https://en.wikipedia.org/wiki/Hampi_Express",
            'https://en.wikipedia.org/wiki/Michael_Jordan'
          ]
        },
        {
          "id": "1",
          "title": [
            "Hampi",
            "Hampi (town)",
            "Hampi Express",
            'Michael Jordan'
          ],
          "url": [
            "https://en.wikipedia.org/wiki/Hampi",
            "https://en.wikipedia.org/wiki/Hampi_(town)",
            "https://en.wikipedia.org/wiki/Hampi_Express",
            'https://en.wikipedia.org/wiki/Michael_Jordan'
          ]
        },
        {
          "id": "1",
          "title": [
            "Hampi",
            "Hampi (town)",
            "Hampi Express",
            'Michael Jordan'
          ],
          "url": [
            "https://en.wikipedia.org/wiki/Hampi",
            "https://en.wikipedia.org/wiki/Hampi_(town)",
            "https://en.wikipedia.org/wiki/Hampi_Express",
            'https://en.wikipedia.org/wiki/Michael_Jordan'
          ]
        },
        {
          "id": "1",
          "title": [
            "Hampi",
            "Hampi (town)",
            "Hampi Express",
            'Michael Jordan'
          ],
          "url": [
            "https://en.wikipedia.org/wiki/Hampi",
            "https://en.wikipedia.org/wiki/Hampi_(town)",
            "https://en.wikipedia.org/wiki/Hampi_Express",
            'https://en.wikipedia.org/wiki/Michael_Jordan'
          ]
        },
        {
          "id": "1",
          "title": [
            "Hampi",
            "Hampi (town)",
            "Hampi Express",
            'Michael Jordan'
          ],
          "url": [
            "https://en.wikipedia.org/wiki/Hampi",
            "https://en.wikipedia.org/wiki/Hampi_(town)",
            "https://en.wikipedia.org/wiki/Hampi_Express",
            'https://en.wikipedia.org/wiki/Michael_Jordan'
          ]
        },
        
        {
          "id": "1",
          "title": [
            "Hampi",
            "Hampi (town)",
            "Hampi Express",
            'Michael Jordan'
          ],
          "url": [
            "https://en.wikipedia.org/wiki/Hampi",
            "https://en.wikipedia.org/wiki/Hampi_(town)",
            "https://en.wikipedia.org/wiki/Hampi_Express",
            'https://en.wikipedia.org/wiki/Michael_Jordan'
          ]
        },
        {
          "id": "2",
          "title": [
            'Michael',
            'Michael Jackson',
            'Michael Jordan',
            'Michael Jordan'
          ],
          "url": [
            'https://en.wikipedia.org/wiki/Michael',
            'https://en.wikipedia.org/wiki/Michael_Jackson',
            'https://en.wikipedia.org/wiki/Michael_Jordan',
            'https://en.wikipedia.org/wiki/Michael_Jordan'
          ]
        }
    ];

    const mappedData = data.map(item => {
        const options = item.title.map((title, index) => ({
          label: title,
          value: item.url[index],
        }));
        return options; 
    });

    const [currentPage, setCurrentPage] = useState(1);
    const [postPerPage, setPostPerPage] = useState(6);

    const lastPostIndex = currentPage * postPerPage;
    const firstPostIndex = lastPostIndex - postPerPage;
    const currentPost = mappedData.slice(firstPostIndex, lastPostIndex);

    return (
        <>
            <ResultList dataResults = {currentPost}/>
            <Pagination 
                totalPosts = {mappedData.length} 
                postPerPage = {postPerPage}
                setCurrentPage={setCurrentPage}
                currentPage={currentPage}
            />
        </>
        
    );
}