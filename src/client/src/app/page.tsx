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
      <div className="set flex flex-row items-center font-bold text-transparent bg-clip-text bg-[linear-gradient(to_right,theme(colors.indigo.400),theme(colors.indigo.100),theme(colors.sky.400),theme(colors.fuchsia.400),theme(colors.sky.400),theme(colors.indigo.100),theme(colors.indigo.400))]">
        <SearchIcon style={{ color: 'white', fontSize: '200px'}}/>
        <h1 className="text-7xl "> Wiki Game Solver </h1>
      </div>

      <InputForm/>

      <Result />

    </main>
  );
}
