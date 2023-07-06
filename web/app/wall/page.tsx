import { getAllTasks } from "@/lib/task/api/getTask";
import { Task } from "@/lib/task/types/task";
import Link from "next/link";

export default async function Page() {
  const taskData: Promise<Task[]> = getAllTasks("tasks");
  const tasks = await taskData;

  const content = (
      <main className="flex min-h-screen flex-col items-center justify-between p-24">
        <h1>Hello</h1>

        {tasks.map(task => {
          return (
            makeTaskCard(task)
          );
        })}
      </main>
  );

  return content;
}

const makeTaskCard = (task: Task): JSX.Element => {
  return (
    <Link 
      href={`/wall/${task.Id}`}
      key={task.Id}
    >
      {task.Name} : {task.Description} : {task.Deadline}
    </Link>
  );
};
