import { useState, useEffect } from "react";
import CloseIcon from "@mui/icons-material/Close";
import SearchIcon from '@mui/icons-material/Search';

function Input(props){
    const { type, placeholder, name } = props;

    const [search, setSearch] = useState("");
    const [searchData, setSearchData] = useState([])
    const [selectedItem, setSelectedItem] = useState(-1)

    const handleChange = e => {
        setSearch(e.target.value)
    }

    const handleClose = () => {
        setSearch("")
        setSearchData([])
        setSelectedItem(-1)
    }

    const handleKeyDown = e => {
        if(selectedItem < searchData.length) {
            if(e.key === "ArrowUp" && selectedItem > 0) {
                setSelectedItem(prev => prev - 1)
            } else if (e.key === "ArrowDown" && selectedItem < searchData.length - 1) {
                setSelectedItem(prev => prev + 1)
            }
            else if (e.key === "Enter" && selectedItem >= 0) {
                setSearch(options[selectedItem].label);
                setSearchData([])
                setSelectedItem(-1)
            }
        } else {
            setSelectedItem(-1)
        }
        
    }

    useEffect(() => {
        if(search !== "") {
            fetch(`https://en.wikipedia.org/w/api.php?origin=*&action=opensearch&format=json&limit=4&namespace=0&search=${search}`)
                .then(res => res.json())
                .then(data => setSearchData(data));
        }
   }, [search]); 

    const [searchTerm, titles, emptyArray, urls] = searchData;

    const options = titles && urls ? titles.map((title, index) => ({
        label: title,
        value: urls[index]
    })) : [];

    const handleOptionClick = (value) => {
        setSearch(value);
    }

    return (
        <div className="p-5 w-full h-16 border-b border-white border-b-0 flex flex-col items-center relative">
            <input 
                type={type} 
                className="text-sm border rounded w-full py-2 px-3 ext-slate-700" 
                placeholder={placeholder} 
                name={name} 
                autoComplete="off"
                onChange={handleChange}
                value={search}
                onKeyDown={handleKeyDown}
            />
            <div className = "bg-gray-300 px-4 h-full flex items-center cursor-pointer">
                {
                    search === "" ? (
                        <SearchIcon />
                    ) : (
                        <CloseIcon onClick = {handleClose}/>
                    )
                }
            </div>
            
            <div className="search_result bg-white absolute top-full mt-5 w-full z-10">
                {
                search !== "" &&(options.map((data, index) => {
                    return (
                    <div
                        key={index}
                        // className="px-5 py-2 cursor-pointer text-xl block"
                        className = {selectedItem === index ?
                            "px-5 py-2 cursor-pointer text-xl block hover:bg-gray-400 bg-blue-200" :
                            "px-5 py-2 cursor-pointer text-xl block hover:bg-gray-400"
                        }
                        onClick={() => handleOptionClick(data.label)}
                    > 
                        {data.label}
                    </div>
                    )
                }))
                }
            </div>
        </div>

    );
}

export default Input