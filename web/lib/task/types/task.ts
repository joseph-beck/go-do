import { Deadline, makeDeadline } from "./deadline";

export type Task = {
  id: number;
  name: string;
  description: string;
  complete: boolean;
  deadline: Deadline;
};

export type TaskPost = {
  id: number;
  name: string;
  description: string;
  complete: boolean;
  deadline: string;
};

export function toTask(taskPost: TaskPost): Task {
  const task: Task = {
    id: taskPost.id,
    name: taskPost.name,
    description: taskPost.description,
    complete: taskPost.complete,
    deadline: makeDeadline(taskPost.deadline),
  };

  return task;
} 
