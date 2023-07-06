import { getTask } from "@/lib/task/api/getTask";
import { Task } from "@/lib/task/types/task";
import Link from "next/link";

type Params = {
  params: {
    Table: string;
    Id: number;
  };
};

export default async function Page({ params: { Table, Id } }: Params) {
  //const taskData: Promise<Task> = getTask(Table, Id);
  //const task = await taskData;
//
  //console.log(task);

  const content = (
      <main className="flex min-h-screen flex-col items-center justify-between p-24">
        <h1>Task {Id} {Table}</h1>

        {/* {makeTaskCard(task)} */}
      </main>
  );

  return content;
}

const makeTaskCard = (task: Task): JSX.Element => {
  return (
    <Link href={`/wall`}>
      {task.Name} : {task.Description} : {task.Deadline}
    </Link>
  );
};