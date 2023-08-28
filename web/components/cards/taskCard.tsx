interface Props {
  title: string;
  body: string;
}

export const TaskCard: React.FC<Props> = ({ title, body }) => {
  return (
  <div className="mx-auto max-w-md rounded-lg bg-white shadow">
    <div className="p-4">
      <h3 className="text-xl font-medium text-gray-900">{title}</h3>
      <p className="mt-1 text-gray-500">{body}</p>
    </div>
  </div>
  );
};
