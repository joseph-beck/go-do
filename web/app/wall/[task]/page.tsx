import { getTask } from "@/lib/api/taskGet";
import { Task } from "@/lib/types/task";
import Link from "next/link";

type Params = {
  params: {
    Table: string;
    Id: number;
  };
};

export default async function Page({ params: { Table, Id } }: Params) {
  console.log(Table, Id)
  // const taskData: Promise<Task> = getTask("tasks", Id);
  // const task = await taskData;
  // console.log(task);

  const content = (
      <main className="flex min-h-screen flex-col items-center justify-between p-24">
        <h1>Task {Id} {Table}</h1>

        {/* { makeTaskCard(task) } */}
      </main>
  );

  return content;
}

const makeTaskCard = (task: Task): JSX.Element => {
  return (
    <div>
      <h1>{task.name}</h1>
      <p>{task.description}</p>
    </div>
  );
};

const makeEditButton = (): JSX.Element => {
  const click = () => {
    console.log("edit");
  };

  return (
    <button
      onClick={click}
    >
      Edit
    </button>
  );
};

const makeDeleteButton = (): JSX.Element => {
  const click = () => {
    console.log("delete");
  };

  return (
    <button
      onClick={click}
    >
      Delete
    </button>
  );
};
