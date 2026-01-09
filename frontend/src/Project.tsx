import { Button } from "./Button";
import { TaskItem } from "./TaskItem";
import type { Task } from "./tasks";

export interface ProjectProps {
  projectName: string;
  projectTasks: Task[];
  removeProject: () => void;
  addChildTask: (taskId: string) => void;
  removeTask: (taskId: string) => void;
  toggleTaskCompletion: (taskId: string, isCompleted: boolean) => void;
  addTask: () => void;
}

export function Project({
  projectName,
  projectTasks,
  removeProject,
  addChildTask,
  removeTask,
  toggleTaskCompletion,
  addTask,
}: ProjectProps) {
  return (
    <div className="shadow-md rounded-md p-4 mb-4 bg-tangerine dark:bg-slate-600">
      <div className="flex items-start justify-between border-b-2 border-strawberry pb-2 mb-2">
        <div>
          <h2 className="text-xl font-extrabold mb-1">{projectName}</h2>
          <span className="block text-sm text-banana">
            {projectTasks.length} tasks
          </span>
        </div>

        <Button onClick={removeProject}>Delete Project</Button>
      </div>

      <ul className="mb-4">
        {projectTasks.map((task) => (
          <TaskItem
            key={task.id}
            task={task}
            addChildTask={addChildTask}
            removeTask={removeTask}
            toggleTaskCompletion={toggleTaskCompletion}
          />
        ))}
      </ul>

      <div className="flex justify-end pt-2">
        <Button onClick={addTask}>Add Task</Button>
      </div>
    </div>
  );
}
