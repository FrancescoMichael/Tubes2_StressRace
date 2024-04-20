"use client";
import Link from "next/link";
import React, { useState } from "react";
import { FaBars, FaTimes } from "react-icons/fa";
import {roboto} from "@/app/ui/font"
import { usePathname } from "next/navigation";
import clsx from 'clsx';

const Navbar = () => {
  const [nav, setNav] = useState(false);
  const pathname = usePathname()

  const links = [
    { name: "Home", link: "/" },
    { name: "WikiRace", link: "wiki-race"},
    { name: "About Us", link: "about" },
  ];

  return (
    <div className="sticky backdrop-blur-sm top-0">
      <nav className={ `${roboto.className} flex flex-wrap w-[100%] items-center justify-between bg-[transparent] md:p-0` }  >
      <h1 className="p-3 px-4 text-2xl text-white">
        <Link href="/" className="link-underline link-underline-black">
          StressRace
        </Link>
      </h1>

      <ul className="hidden md:flex">
        {links.map(({ name, link }) => (
          <li key={name} >
            <Link href={`/${link}`}
              className={clsx(
                "px-4 cursor-pointer capitalize font-medium text-gray-500 hover:text-white focus:text-white duration-200 link-underline",
                {
                  'text-white': pathname === link,
                },
              )}
            >{name}</Link>
          </li>
        ))
        }
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
    </div>
    
  );
};

export default Navbar;
