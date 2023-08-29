import { CheckBox } from "../input/checkBox";

interface Props {
  index: number;
  title: string;
  body: string;
  complete: boolean;
}

export const TaskCard: React.FC<Props> = ({ index, title, body, complete }) => {
  return (
  <div className="mx-auto max-w-sm rounded-lg bg-white shadow">
    <div className="p-4">
      <h3 className="text-xl font-medium text-gray-900">{title}</h3>
      <p className="mt-1 text-gray-500">{body}</p>
      <CheckBox
        id={index.toString()}
        text="complete"
        ticked={complete}
      />
    </div>
  </div>
  );
};
