# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

### Development
```bash
npm run dev        # Start development server
npm run build      # Build for production (runs TypeScript compiler then Vite build)
npm run preview    # Preview production build
npm run lint       # Run ESLint
```

## Architecture

### Tech Stack
- **Frontend**: React 19 with TypeScript
- **Build Tool**: Vite with SWC for Fast Refresh
- **Styling**: Tailwind CSS v4 with custom theme colors
- **State Management**: TanStack Query (React Query) for server state
- **Backend**: REST API running on `http://localhost:8080`

### Application Structure

This is a hierarchical project/task management application with a client-server architecture.

**Data Model**:
- **Projects**: Top-level entities with a name and collection of tasks
- **Tasks**: Hierarchical tree structure where tasks can have child tasks (recursive)
- Tasks optionally contain a `link` field for external references

**API Integration**:
The frontend communicates with a backend REST API via TanStack Query:
- `GET /projects` - Fetches all projects with nested tasks
- `POST /projects` - Creates a new project
- `DELETE /projects/:id` - Deletes a project
- `POST /tasks` - Creates a task (accepts `projectId`, optional `parentTaskId`, and `content`)
- `DELETE /tasks/:id` - Deletes a task

All mutations automatically invalidate the `["projects"]` query key to refetch data.

**Component Hierarchy**:
- `App.tsx` - Root component managing projects list and coordinating all CRUD operations
- `Project.tsx` - Renders individual project with tasks list
- `TaskItem.tsx` - Recursively renders tasks and their children
- `Button.tsx` - Shared button component

**Data Flow**:
- API hooks are defined in `projects.ts` and `tasks.ts`
- `App.tsx` fetches projects and passes down mutation functions as callbacks
- Child components trigger mutations which invalidate queries, causing automatic refetch

### Custom Theme

Tailwind is configured with a custom color palette in `src/main.css`:
- `tangerine` - Primary orange accent color (burnt-peach-400)
- `banana` - Secondary yellow accent (sunflower-gold-400)
- `strawberry` - Light pink accent (pastel-pink-200)

The theme also defines complete color scales for pastel-pink, sunflower-gold, and burnt-peach.

### TypeScript Configuration

Uses TypeScript project references:
- `tsconfig.app.json` - Application source code
- `tsconfig.node.json` - Vite configuration files
- `tsconfig.json` - Composite configuration

### Linting

ESLint flat config with:
- TypeScript ESLint recommended rules
- React Hooks rules
- React Refresh plugin for Vite HMR
- Prettier compatibility
