export type Props = {
  text: string;
  href: string;
}

export const Blink: React.FC<Props> = ({ text, href }): JSX.Element => {
  const content = (
    <div className="relative group">
      <div className="absolute -inset-0.5 bg-gradient-to-r from-pink-600 to-purple-600 rounded-lg blur opacity-75 group-hover:opacity-100 transition duration-1000 group-hover:duration-200 animate-tilt"></div>
      <li className="relative bg-black rounded-lg leading-0 px-5 py-2 text-indigo-400 hover:text-gray-200 transition duration-1000 group-hover:duration-200">
        <a href={href}>
          {text}
        </a>
      </li>
    </div>
  );

  return content;
};

export const SpecialBlink: React.FC<Props> = ({ text, href }): JSX.Element => {
  const content = (
    <div className="relative group">
      <div className="absolute -inset-1 bg-gradient-to-r from-pink-600 to-purple-600 rounded-lg blur opacity-75 group-hover:opacity-100 transition duration-1000 group-hover:duration-200 animate-tilt"></div>
      <li className="relative bg-gradient-to-r from-indigo-500 to-pink-500 rounded-lg leading-0 px-5 py-2 text-black hover:text-gray-200 transition duration-1000 group-hover:duration-200">
        <a href={href}>
          {text}
        </a>
      </li>
    </div>
  );

  return content;
};
