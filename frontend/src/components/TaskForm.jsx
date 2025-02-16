import React, { useState } from 'react';

const TaskForm = ({ onTaskAdded }) => {
  const [task, setTask] = useState({
    title: '',
    priority: 'Средний',
    dueDate: ''
  });
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [error, setError] = useState('');

  const handleSubmit = async (e) => {
    e.preventDefault();
    if (!task.title.trim() || isSubmitting) return;

    try {
      setIsSubmitting(true);
      setError('');
      
      const formattedDate = task.dueDate ? task.dueDate : '';
      
      await window.go.main.TodoApp.AddTask(
        task.title.trim(),
        task.priority,
        formattedDate
      );
      
      setTask({ title: '', priority: 'Средний', dueDate: '' });
      
      if (onTaskAdded) {
        await onTaskAdded();
      }
    } catch (error) {
      setError(error.message || 'Произошла ошибка при добавлении задачи');
    } finally {
      setIsSubmitting(false);
    }
  };

  return (
    <form onSubmit={handleSubmit} className="bg-white p-6 rounded-lg shadow-sm border">
      {error && (
        <div className="mb-4 p-3 bg-red-50 border border-red-200 text-red-700 rounded">
          {error}
        </div>
      )}
      
      <div className="space-y-4">
        <div>
          <input
            type="text"
            value={task.title}
            onChange={(e) => setTask({ ...task, title: e.target.value })}
            placeholder="Что нужно сделать?"
            className="w-full p-3 border border-gray-200 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            disabled={isSubmitting}
          />
        </div>
        
        <div className="flex space-x-4">
          <select
            value={task.priority}
            onChange={(e) => setTask({ ...task, priority: e.target.value })}
            className="flex-1 p-3 border border-gray-200 rounded-lg bg-white"
            disabled={isSubmitting}
          >
            <option value="Низкий">🟢 Низкий приоритет</option>
            <option value="Средний">🟡 Средний приоритет</option>
            <option value="Высокий">🔴 Высокий приоритет</option>
          </select>
          
          <input
            type="date"
            value={task.dueDate}
            onChange={(e) => setTask({ ...task, dueDate: e.target.value })}
            className="flex-1 p-3 border border-gray-200 rounded-lg"
            disabled={isSubmitting}
          />
        </div>
      </div>
      
      <button
        type="submit"
        disabled={isSubmitting || !task.title.trim()}
        className={`
          w-full mt-4 p-3 rounded-lg font-medium transition-all
          ${isSubmitting || !task.title.trim()
            ? 'bg-gray-100 text-gray-400 cursor-not-allowed'
            : 'bg-blue-600 text-white hover:bg-blue-700 active:bg-blue-800'}
        `}
      >
        {isSubmitting ? 'Добавление...' : 'Добавить задачу'}
      </button>
    </form>
  );
};

export default TaskForm; 