import React, { useState, useEffect } from "react";
import TaskList from "./components/TaskList";
import TaskForm from "./components/TaskForm";

function App() {
  const [tasks, setTasks] = useState([]);
  const [error, setError] = useState(null);

  const loadTasks = async () => {
    try {
      setError(null);
      const loadedTasks = await window.go.main.TodoApp.GetTasks();
      console.log('Загруженные задачи:', loadedTasks);
      setTasks(loadedTasks || []);
    } catch (error) {
      console.error('Ошибка при загрузке задач:', error);
      setError('Ошибка при загрузке задач');
    }
  };

  useEffect(() => {
    loadTasks();
  }, []);

  const handleToggleComplete = async (id) => {
    try {
      await window.go.main.TodoApp.ToggleTaskCompletion(id);
      await loadTasks();
    } catch (error) {
      console.error('Ошибка при изменении статуса задачи:', error);
    }
  };

  const handleDeleteTask = async (id) => {
    try {
      await window.go.main.TodoApp.RemoveTask(id);
      await loadTasks();
    } catch (error) {
      console.error('Ошибка при удалении задачи:', error);
      setError('Не удалось удалить задачу');
    }
  };

  return (
    <div className="container mx-auto p-4 max-w-2xl">
      <h1 className="text-2xl font-bold mb-6 text-center">Список задач</h1>
      {error && (
        <div className="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded mb-4">
          {error}
        </div>
      )}
      <TaskForm onTaskAdded={loadTasks} />
      <TaskList
        tasks={tasks}
        onToggleComplete={handleToggleComplete}
        onDeleteTask={handleDeleteTask}
      />
    </div>
  );
}

export default App;