'use client'
import InputForm from "@/components/Elements/Input/form"
import SearchIcon from '@mui/icons-material/Search';
import Result from "@/components/Elements/Result/Result";
import StarryBackground from "@/components/Elements/Back/Background"
import { useState } from "react";

export default function Home() {
    const [isLoading, setIsLoading] = useState<boolean>(false);
    const [startTime, setStartTime] = useState(0);

    return (
        <main className="flex flex-col items-center justify-center p-12">
            <StarryBackground/>
            <div className="judul set text-white flex flex-row items-center font-bold]">
                <SearchIcon fontSize="inherit" color="inherit"/>
                <h1 className="font-bold"> Wiki Game Solver </h1>
            </div>
            <InputForm isLoading = {isLoading} setIsLoading = {setIsLoading} startTime = {startTime} setStartTime = {setStartTime}/>
            <Result isLoading = {isLoading} setIsLoading = {setIsLoading} startTime = {startTime}/>
        </main>

    );
}
