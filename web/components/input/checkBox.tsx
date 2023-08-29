interface Props {
  id: string;
  text: string;
  ticked: boolean;
}

export const CheckBox: React.FC<Props> = ({ id, text, ticked }): JSX.Element => {
  return (
    <div className="flex items-center space-x-2">
      <input
        type="checkbox"
        id={id}
        className="h-4 w-4 rounded border-gray-300 text-primary-600 shadow-sm focus:border-primary-300 focus:ring focus:ring-primary-200 focus:ring-opacity-50 focus:ring-offset-0 disabled:cursor-not-allowed disabled:text-gray-400"
      />
      <label
        htmlFor={id}
        className="text-sm font-medium text-gray-700"
      >
        {text}
      </label>
    </div>
  );
}