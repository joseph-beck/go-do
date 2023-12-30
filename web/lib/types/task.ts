import { Deadline, makeDeadline } from "./deadline";

export type Task = {
  id: number;
  name: string;
  description: string;
  complete: boolean;
  deadline: Deadline;
};

export type Tasks = {
  tasks: Task[];
  task_count: number;
};