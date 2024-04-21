import Image from "next/image";

export default function Home() {
  return (
    <main className="flex min-h-screen flex-col items-center p-24">
      <div className="bg-white">
        <img src="https://images.unsplash.com/photo-1566396223585-c8fbf7fa6b6d?q=80&w=1898&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D  " alt="TAI" className="h-64 w-full object-cover object-center"/>
      </div>
      <div className="text-white text-lg text-justify pt-20">
        <p>
        WikiRace atau Wiki Game adalah permainan yang melibatkan Wikipedia, sebuah ensiklopedia daring gratis yang dikelola oleh berbagai relawan di dunia, dimana pemain mulai pada suatu artikel Wikipedia dan harus menelusuri artikel-artikel lain pada Wikipedia (dengan mengeklik tautan di dalam setiap artikel) untuk menuju suatu artikel lain yang telah ditentukan sebelumnya dalam waktu paling singkat atau klik (artikel) paling sedikit.
        </p>  
      </div>  
    </main>
  );
}
