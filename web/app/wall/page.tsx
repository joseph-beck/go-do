'use client';

import { useState, useEffect } from 'react';
import { listTasks } from "@/lib/api/taskGet";
import { Task } from "@/lib/types/task";
import { TaskCard } from '@/components/cards/taskCard';

export default function Page() {
  const [content, setContent] = useState<JSX.Element | null>(null);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const tasksData = await listTasks("tasks");
        const taskElements = tasksData.map((task: Task) => taskCard(task));

        const pageContent = (
          <main className="flex min-h-screen flex-col items-center justify-between p-24">
            <h1>Hello</h1>
            <button onClick={createForm} className="flex top-0 left-0">Create</button>
            {taskElements}
          </main>
        );

        setContent(pageContent);
      } catch (error) {
        console.error('Error fetching tasks:', error)

        setContent(<p>Error fetching tasks</p>);
      }
    };

    fetchData();
  }, []);

  return content;
}

const taskCard = (task: Task): JSX.Element => {
  return (
    <div>
      <TaskCard
        title={task.name}
        body={task.description}
      />
    </div>
  );
};

const createForm = (): JSX.Element => {
  console.log("create form")

  return (
    <form>

    </form>
  );
};
