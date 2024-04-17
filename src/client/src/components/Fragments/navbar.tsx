"use client";
import Link from "next/link";
import React, { useState } from "react";
import { FaBars, FaTimes } from "react-icons/fa";
import {josefin} from "@/app/ui/font"

const Navbar = () => {
  const [nav, setNav] = useState(false);

  const links = [
    { name: "Home", link: "/" },
    { name: "WikiRace", link: "wiki-race"},
    { name: "About Us", link: "about" },
   
  ];

  return (
    <nav className={ `${josefin.className} flex flex-wrap items-center justify-between bg-[#31363F] md:p-0` }  >
      <h1 className="p-3 px-4 text-2xl text-white">
        <Link href="/" className="link-underline link-underline-black">
          StressRace
        </Link>
      </h1>

      <ul className="hidden md:flex">
        {links.map(({ name, link }) => (
          <li key={name} className="px-4 cursor-pointer capitalize font-medium text-gray-500 hover:text-white focus:text-white duration-200 link-underline">
            <Link href={`/${link}`}>{name}</Link>
          </li>
        ))}
      </ul>

      <button
        aria-label="Toggle navigation"
        aria-expanded={nav}
        onClick={() => setNav(!nav)}
        className="cursor-pointer pr-4 z-10 text-gray-500 md:hidden"
      >
        {nav ? <FaTimes size={30} /> : <FaBars size={30} />}
      </button>

      {nav && (
        <ul className="flex flex-col justify-center items-center absolute top-0 left-0 w-full h-screen bg-gradient-to-b from-black to-gray-800 text-gray-500">
          {links.map(({name, link }) => (
            <li key={name} className="px-4 cursor-pointer capitalize py-6 text-4xl">
              <Link href={`/${link}`} onClick={() => setNav(false)}>
                {link}
              </Link>
            </li>
          ))}
        </ul>
      )}
    </nav>
  );
};

export default Navbar;
