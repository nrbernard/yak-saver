import type { Task } from "./tasks";

export function TaskItem({
  task,
  addChildTask,
  removeTask,
  toggleTaskCompletion,
}: {
  task: Task;
  addChildTask: (taskId: string) => void;
  removeTask: (taskId: string) => void;
  toggleTaskCompletion: (taskId: string, isCompleted: boolean) => void;
}) {
  return (
    <li key={task.id}>
      <div className="flex items-center justify-between">
        <input
          type="checkbox"
          checked={!!task.completedAt}
          onChange={() => toggleTaskCompletion(task.id, !!task.completedAt)}
          className="mr-2 cursor-pointer accent-banana"
        />

        <div className={`flex-1 ${task.completedAt ? "line-through" : ""}`}>
          {task.link ? (
            <a href={task.link} target="_blank" className="flex-1">
              {task.content}
            </a>
          ) : (
            <span>{task.content}</span>
          )}
        </div>

        <div className="flex items-center gap-4">
          <button
            className="justify-self-end text-xl cursor-pointer hover:text-banana"
            onClick={() => addChildTask(task.id)}
          >
            +
          </button>
          <button
            className="justify-self-end text-xl cursor-pointer hover:text-banana"
            onClick={() => removeTask(task.id)}
          >
            -
          </button>
        </div>
      </div>
      {task.children.length > 0 && (
        <ul className="mt-2 ml-4">
          {task.children.map((child) => (
            <TaskItem
              key={child.id}
              task={child}
              addChildTask={addChildTask}
              removeTask={removeTask}
              toggleTaskCompletion={toggleTaskCompletion}
            />
          ))}
        </ul>
      )}
    </li>
  );
}
