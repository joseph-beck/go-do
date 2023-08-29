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
        const taskElements = tasksData.map((task: Task, index: number) => taskCard(task, index));

        const pageContent = (
          <div className="flex min-h-screen flex-col">
            <header className="bg-blue-200 p-4">Header</header>
            <div className="flex flex-1 flex-row">
              <main className="flex-1 bg-blue-50 p-4">
                <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
                  {taskElements}
                </div>
              </main>
              <nav className="order-first w-32 bg-blue-100 p-4">Navigation</nav>
              <aside className="w-32 bg-blue-100 p-4">Side</aside>
            </div>
            <footer className="bg-blue-200 p-4">Footer</footer>
          </div>
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

const taskCard = (task: Task, index: number): JSX.Element => {
  return (
    <TaskCard
      index={index}
      title={task.name}
      body={task.description}
      complete={task.complete}
    />
  );
};

const createForm = (): JSX.Element => {
  console.log("create form")

  return (
    <form>

    </form>
  );
};
