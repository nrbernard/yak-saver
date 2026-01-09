import { useCreateProject, useDeleteProject, useProjects } from "./projects";
import { useCreateTask, useDeleteTask, useUpdateTask } from "./tasks";
import { Project } from "./Project";

export default function App() {
  const { data: serverProjects } = useProjects();
  const { mutate: createTask } = useCreateTask();
  const { mutate: deleteTask } = useDeleteTask();
  const { mutate: updateTask } = useUpdateTask();
  const { mutate: createProject } = useCreateProject();
  const { mutate: deleteProject } = useDeleteProject();

  const addChildTask = (projectId: string) => (taskId: string) => {
    const childContent = prompt("Enter the child task content:");

    if (childContent) {
      const project = serverProjects?.find(
        (project) => project.id === projectId
      );

      if (!project) {
        throw new Error(`Project with id ${projectId} not found`);
      }

      createTask({ projectId, parentTaskId: taskId, content: childContent });
    }
  };

  const removeTask = (taskId: string) => {
    deleteTask({ taskId });
  };

  const toggleTaskCompletion = (taskId: string, isCompleted: boolean) => {
    updateTask({ taskId, completed: !isCompleted });
  };

  const addProject = () => {
    const projectName = prompt("Enter the project name:");
    if (projectName) {
      createProject({ name: projectName });
    }
  };

  const addTask = (projectId: string) => () => {
    const project = serverProjects?.find((project) => project.id === projectId);
    if (!project) {
      throw new Error(`Project with id ${projectId} not found`);
    }
    const taskContent = prompt("Enter the task content:");

    if (taskContent) {
      createTask({ projectId: project.id, content: taskContent });
    }
  };

  const removeProject = (projectId: string) => () => {
    if (confirm("Are you sure you want to delete this project?")) {
      deleteProject({ projectId });
    }
  };

  return (
    <main className="p-4 max-w-3xl mx-auto">
      <div className="flex justify-between items-center mb-4">
        <h1 className="text-2xl font-bold">Projects</h1>
        <button
          onClick={addProject}
          className="px-4 py-1 rounded-md hover:cursor-pointer bg-tangerine text-text-primary hover:text-banana"
        >
          Add Project
        </button>
      </div>

      <div>
        {serverProjects?.map((project) => (
          <Project
            key={project.id}
            projectName={project.name}
            projectTasks={project.tasks ?? []}
            removeProject={removeProject(project.id)}
            addChildTask={addChildTask(project.id)}
            removeTask={removeTask}
            toggleTaskCompletion={toggleTaskCompletion}
            addTask={addTask(project.id)}
          />
        ))}
      </div>
    </main>
  );
}
