import React from "react"
import { Blink, SpecialBlink } from "../input/blink";
import { Divider } from "../layout/divider";
import Link from "next/link";

export type Props = {

}

type Navs = {
  href: string;
  name: string;
}

export const Footer: React.FC<Props> = (): JSX.Element => {

  const footerNavs: Navs[] = [
    { href: '/features', name: 'features' },
    { href: '/about', name: 'about' },
    { href: '/help', name: 'help' },
  ]

  const content = (
    <footer className="bg-black pt-10">
      <div className="max-w-screen-xl mx-auto px-4 text-gray-600 md:px-8">
        <div className="space-y-6 sm:max-w-md sm:mx-auto sm:text-center">
          <p>
            explore what godo has to offer
          </p>
          <div className="items-center gap-x-6 space-y-3 sm:flex sm:justify-center sm:space-y-0">
            <Blink text="get started" href="/signup" />
            <SpecialBlink text="login &rarr;" href="/login" />
          </div>
        </div>
        <div className="mt-10">
          <Divider width="3/4" />
        </div>
        <div className="py-10 items-center justify-between sm:flex">
          <p>godo</p>
          <ul className="flex flex-wrap items-center gap-4 mt-6 sm:text-sm sm:mt-0">
            {
              footerNavs.map((item, idx) => (
                <li key={idx} className="text-gray-800 hover:text-gray-500 duration-150">
                  <Link href={item.href}>
                    {item.name}
                  </Link>
                </li>
              ))
            }
          </ul>
        </div>
      </div>
    </footer>
  );

  return content
};
