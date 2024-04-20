import Image from "next/image";
import Input from "@/components/Elements/Input/input"
import InputForm from "@/components/Elements/Input/form"
import Algo from "@/components/Elements/Input/algo"
import MyComponent from "@/components/Elements/Input/lee"
import SearchIcon from '@mui/icons-material/Search';
import Result from "@/components/Elements/Result/Result";


export default function Home() {
  return (
    <main className="flex flex-col items-center justify-center p-12">
      <div className="flex flex-row items-center">
        <SearchIcon style={{ color: 'white', fontSize: '200px'}}/>
        <h1 className="text-7xl text-white font-bold"> Wiki Game Solver </h1>
      </div>

      <Algo>
      </Algo>

      <InputForm>
      </InputForm>      

      <InputForm>
      </InputForm>
      
      <InputForm>
      </InputForm>
      
      <InputForm>
      </InputForm>
      <Result />

    </main>
  );
}
