import Image from "next/image";
import React, { useState } from "react";
import {roboto} from "@/app/ui/font"
import img1 from '../../../../public/images/1.png';

export default function Home() {
  return (
    <main className="flex min-h-screen flex-col items-center p-24 flex-wrap">
      <div className="flex flex-row flex-wrap justify-center gap-4 items-center h-full animate-fadeIn">
        <div>
          <img src="https://images.unsplash.com/photo-1566396223585-c8fbf7fa6b6d?q=80&w=1898&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D  " alt="TAI" className="h-48 w-96 object-cover object-center"/>
        </div>
        <span className="text-white text-lg text-justify w-8/12">
        Wikipedia berawal sebagai proyek sampingan Nupedia berupa ensiklopedia bebas daring yang artikelnya ditulis oleh para ahli. Larry Sanger, yang mendirikan Nupedia bersama Jimmy Wales, melontarkan ide mengenai ensiklopedia berbasis wiki pada 10 Januari 2001 di milis Nupedia. Kemudian pada 15 Januari 2001, Wikipedia secara resmi diluncurkan di situs web www.wikipedia.com. Isi Wikipedia dapat diciptakan oleh penggunanya. Pengunjung di Wikipedia juga dapat mengubah artikel, dan banyak yang melakukannya. Halaman-halaman selalu diubah, jadi, tidak ada artikel yang pernah selesai. Dan karena itu pula di Wikipedia sering terjadi "kesulitan" yang unik. Tetapi ia pun memiliki sistem "memperbaiki sendiri" atau otomatis.
        </span>  
      </div>
      <div className="flex flex-row flex-wrap justify-center gap-4 items-center h-full mt-12 animate-fadeIn">
        <span className="text-white text-lg text-justify w-8/12">
        WikiRace atau Wiki Game adalah permainan yang melibatkan Wikipedia, sebuah ensiklopedia daring gratis yang dikelola oleh berbagai relawan di dunia, dimana pemain mulai pada suatu artikel Wikipedia dan harus menelusuri artikel-artikel lain pada Wikipedia untuk menuju suatu artikel lain yang telah ditentukan sebelumnya dalam waktu paling singkat atau klik paling sedikit. Dalam bahasa Go yang mengimplementasikan algoritma IDS dan BFS untuk menyelesaikan permainan WikiRace. Program menerima masukan berupa jenis algoritma, judul artikel awal, dan judul artikel tujuan. Program memberikan keluaran berupa jumlah artikel yang diperiksa, jumlah artikel yang dilalui, rute penjelajahan artikel, dan waktu pencarian.
        </span>    
        <div>
          <img src="https://miro.medium.com/v2/resize:fit:1400/1*jxmEbVn2FFWybZsIicJCWQ.png" alt="TAI" className="h-48 w-96 object-cover object-center"/>
        </div>
      </div>
      <div className="text-white text-lg text-justify mt-12">
        <span className="text-3xl font-bold">
          HOW TO USE
        </span>
      </div>

      <div className="text-white text-lg text-center mt-4 max-w-xl">
        <p className="text-xl">
          1 
        </p>
        <p className="text-xl mt-4">
          Pilih terlebih dahulu jenis algoritma yang anda ingin gunakan
        </p>
      </div>
      <div>
        <img src="images/1.png" className="w-48 mt-8" alt=""/>
      </div>

      <div className="text-white text-lg text-center mt-4 max-w-xl">
      <p className="text-xl">
          2 
        </p>
        <p className="text-xl mt-4">
          Masukkan judul laman awal pencarian pada box seperti dibawah ini
        </p>
      </div>
      <div>
        <img src="images/2.png" className="w-96 mt-8" alt=""/>
      </div>

      <div className="text-white text-lg text-center mt-4 max-w-xl">
        <p className="text-xl">
          3 
        </p>
        <p className="text-xl mt-4">
          Masukkan judul laman akhir pencarian pada box seperti dibawah ini
        </p>
      </div>
      <div>
        <img src="images/3.png" className="w-96 mt-8" alt=""/>
      </div>

      <div className="text-white text-lg text-center mt-4 max-w-xl">
      <p className="text-xl">
          4 
        </p>
        <p className="text-xl mt-4">
          Anda dapat menukar judul laman awal dengan judul laman akhir dengan tombol dibawah
        </p>
      </div>
      <div>
        <img src="images/4.png" className="mt-8" alt=""/>
      </div>

      
      <div className="text-white text-lg text-center mt-4 max-w-xl">
        <p className="text-xl">
          5 
        </p>
        <p className="text-xl mt-4">
          Mulai pencarian dengan menekan tombol Search! seperti dibawah
        </p>
      </div>
      <div>
        <img src="images/5.png" className="mt-8" alt=""/>
      </div>
            
    </main>
  );
}
