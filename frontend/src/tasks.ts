import { useMutation, useQueryClient } from "@tanstack/react-query";
import type { Project } from "./projects";
import { API_BASE_URL } from "./config";

export interface Task {
  id: string;
  content: string;
  children: Task[];
  link?: string;
  completedAt?: string;
}

export function useCreateTask() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: async ({
      projectId,
      parentTaskId,
      content,
    }: {
      projectId: string;
      parentTaskId?: string;
      content: string;
    }) => {
      const response = await fetch(`${API_BASE_URL}/tasks`, {
        method: "POST",
        body: JSON.stringify({ projectId, parentTaskId, content }),
        headers: {
          "Content-Type": "application/json",
        },
      });
      if (!response.ok) {
        throw new Error("Failed to fetch projects");
      }
      return response.json();
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["projects"] });
    },
  });
}

export function useDeleteTask() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: async ({ taskId }: { taskId: string }) => {
      await fetch(`${API_BASE_URL}/tasks/${taskId}`, {
        method: "DELETE",
      });
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["projects"] });
    },
  });
}

// Helper function to recursively update task completedAt
function updateTaskRecursively(
  tasks: Task[],
  taskId: string,
  completed: boolean
): Task[] {
  return tasks.map((task) => {
    if (task.id === taskId) {
      return {
        ...task,
        completedAt: completed ? new Date().toISOString() : undefined,
      };
    }
    if (task.children.length > 0) {
      return {
        ...task,
        children: updateTaskRecursively(task.children, taskId, completed),
      };
    }
    return task;
  });
}

export function useUpdateTask() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: async ({
      taskId,
      completed,
    }: {
      taskId: string;
      completed: boolean;
    }) => {
      const response = await fetch(`${API_BASE_URL}/tasks/${taskId}`, {
        method: "PATCH",
        body: JSON.stringify({ completed }),
        headers: {
          "Content-Type": "application/json",
        },
      });
      if (!response.ok) {
        throw new Error("Failed to update task");
      }
      return response.json();
    },

    // Optimistically update UI immediately before server responds
    onMutate: async ({ taskId, completed }) => {
      // Cancel outgoing refetches to avoid overwriting optimistic update
      await queryClient.cancelQueries({ queryKey: ["projects"] });

      // Snapshot current state for rollback
      const previousProjects = queryClient.getQueryData<Project[]>(["projects"]);

      // Optimistically update the task's completedAt field
      queryClient.setQueryData<Project[]>(["projects"], (old) => {
        if (!old) return old;

        return old.map((project) => ({
          ...project,
          tasks: updateTaskRecursively(project.tasks ?? [], taskId, completed),
        }));
      });

      // Return context with snapshot for rollback
      return { previousProjects };
    },

    // Rollback on error
    onError: (_err, _variables, context) => {
      if (context?.previousProjects) {
        queryClient.setQueryData(["projects"], context.previousProjects);
      }
    },

    // Refetch to ensure sync with server
    onSettled: () => {
      queryClient.invalidateQueries({ queryKey: ["projects"] });
    },
  });
}
