import InputForm from "@/components/Elements/Input/form"
import SearchIcon from '@mui/icons-material/Search';
import Result from "@/components/Elements/Result/Result";

export default function Home() {
    return (
        <main className="flex flex-col items-center justify-center p-12">
            <div className="set flex flex-row items-center font-bold]">
                <SearchIcon style={{ color: 'white', fontSize: '200px'}}/>
                <h1 className="text-7xl text-white font-bold"> Wiki Game Solver </h1>
            </div>

            <InputForm/>

            <Result />
        </main>
    );
}
