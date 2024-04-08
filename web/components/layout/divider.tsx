interface Props {
  width: "1/4" | "1/2" | "3/4" | "full";
}

export const Divider: React.FC<Props> = ({ width }): JSX.Element => {
  const content = (
    <div className={`w-${width} mx-auto`}>
      <span className="relative flex justify-center">
        <div
          className="absolute inset-x-0 top-1/2 h-px -translate-y-1/2 bg-transparent bg-gradient-to-r from-transparent via-gray-500 to-transparent opacity-75"
        />
      </span>
    </div>
  );

  return content
};
