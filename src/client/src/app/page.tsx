import Image from "next/image";
import Input from "@/components/Elements/Input/input"
import InputForm from "@/components/Elements/Input/form"


export default function Home() {
  return (
    <main className="flex min-h-screen flex-col items-center justify-center p-24">
      <div className="flex flex-row items-center">
        <Image src="/magnifyingGlass.svg" width={100} height={100} alt="magnifyingGlass" className="mag"></Image>
        <h1 className="text-6xl text-white"> Wiki Game Solver </h1>
      </div>
      <InputForm>
      </InputForm>

    </main>
  );
}
