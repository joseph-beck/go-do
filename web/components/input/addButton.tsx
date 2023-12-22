export type Props = {
  size: "xs" | "s" | "m" | "l" | "xl";
}

export const AddButton: React.FC<Props> = ({ size }): JSX.Element => {
  const getButton = (buttonSize: string, iconSize: string) => (
    <div className="relative group">
      <div className="absolute -inset-0.5 bg-gradient-to-r from-pink-600 to-purple-600 rounded-lg blur opacity-75 group-hover:opacity-100 transition duration-1000 group-hover:duration-200 animate-tilt"></div>
      <button
        className={`relative bg-black px-${buttonSize} py-${buttonSize} rounded-lg leading-0 text-indigo-400 hover:text-gray-200 transition duration-1000 group-hover:duration-200`}
      >
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className={`w-${iconSize} h-${iconSize}`}>
          <path strokeLinecap="round" strokeLinejoin="round" d="M12 4.5v15m7.5-7.5h-15" />
        </svg>

      </button>
    </div>
  );

  let content: JSX.Element;

  switch (size) {
    case "xs":
      content = getButton("2.5", "5");
      break;
    case "s":
      content = getButton("3", "5");
      break;
    case "m":
      content = getButton("3.5", "6");
      break;
    case "l":
      content = getButton("4", "6");
      break;
    case "xl":
      content = getButton("5", "7");
      break;
    default:
      content = getButton("3", "5");
      break;
  }

  return content;
};