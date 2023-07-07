'use client';

import { getAllTasks } from "@/lib/task/api/getTask";
import { TaskPost } from "@/lib/task/types/task";
import Link from "next/link";

export default async function Page() {
  const taskData: Promise<TaskPost[]> = getAllTasks("tasks");
  const tasks = await taskData;

  const content = (
    <main className="flex min-h-screen flex-col items-center justify-between p-24">
      <h1>Hello</h1>

      <button
        onClick={createForm}
        className="flex top-0 left-0"
      >
        Create
      </button>

      {tasks.map(task => {
        return (
          taskCard(task)
        );
      })}
    </main>
  );

  return content;
}

const taskCard = (task: TaskPost): JSX.Element => {
  return (
    <Link 
      href={`/wall/${task.id}`}
      key={task.id}
    >
      {task.name} : {task.description} : {task.deadline}
    </Link>
  );
};

const createForm = (): JSX.Element => {
  console.log("create form")

  return (
    <form>
      
    </form>
  );
};
