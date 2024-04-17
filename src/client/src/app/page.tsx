import Image from "next/image";

export default function Home() {
  return (
    <main className="flex min-h-screen flex-row items-center justify-center p-24">
      <Image src="/magnifyingGlass.svg" width={100} height={100} alt="magnifyingGlass" className="mag"></Image>
      <h1 className="text-6xl text-white"> Wiki Game Solver </h1>
    </main>
  );
}
