import Image from "next/image";
import Input from "@/components/Elements/Input/input"
import InputForm from "@/components/Elements/Input/form"
import Algo from "@/components/Elements/Input/algo"
import MyComponent from "@/components/Elements/Input/lee"
import Result from "@/components/Elements/Result/Result";


export default function Home() {
  return (
    <main className="flex flex-col items-center justify-center p-12">
      <div className="flex flex-row items-center">
        <Image src="/magnifyingGlass.svg" width={120} height={120} alt="magnifyingGlass" className="mag"></Image>
        <h1 className="text-7xl text-white font-bold"> Wiki Game Solver </h1>
      </div>

      <Algo>
      </Algo>

      <InputForm>
      </InputForm>      
      <Result />

    </main>
  );
}
