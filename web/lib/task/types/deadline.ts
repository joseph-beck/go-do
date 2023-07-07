import { Date } from "./date"
import { Time } from "./time"

export type Deadline = {
    date: Date;
    time: Time;
};

export function makeDeadline(deadlineString: string): Deadline {
  var halves = deadlineString.split("-");
  var date = halves[0].split("/");
  var time = halves[1].split(":");

  const deadline: Deadline = {
    date: {
      day: date[0],
      month: date[1],
      year: date[2],
    },
    time: {
      minute: time[0],
      hour: time[1],
    }, 
  };

  return deadline;
}
