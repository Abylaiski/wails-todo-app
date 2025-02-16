import React from 'react';

const TaskList = ({ tasks, onToggleComplete, onDeleteTask }) => {
  if (tasks.length === 0) {
    return (
      <div className="text-center text-gray-500 mt-8 p-8 bg-gray-50 rounded-lg border-2 border-dashed">
        <p className="text-lg">Список задач пуст</p>
        <p className="text-sm mt-2">Добавьте новую задачу, чтобы начать</p>
      </div>
    );
  }

  return (
    <div className="space-y-3">
      {tasks.map((task) => (
        <div
          key={task.id}
          className={`group p-4 rounded-lg shadow-sm border transition-all ${
            task.completed ? 'bg-gray-50' : 'bg-white hover:shadow-md'
          }`}
        >
          <div className="flex items-center justify-between">
            <div className="flex items-center space-x-4">
              <input
                type="checkbox"
                checked={task.completed}
                onChange={() => onToggleComplete(task.id)}
                className="w-5 h-5 rounded border-gray-300 text-blue-600 focus:ring-blue-500 cursor-pointer"
              />
              <div>
                <h3 className={`font-medium ${task.completed ? 'line-through text-gray-400' : 'text-gray-900'}`}>
                  {task.title}
                </h3>
                <div className="flex items-center space-x-4 mt-1 text-sm">
                  <span className={`px-2 py-1 rounded-full ${getPriorityColor(task.priority)}`}>
                    {task.priority}
                  </span>
                  {task.due_date && (
                    <span className="text-gray-500">
                      Срок: {formatDate(task.due_date)}
                    </span>
                  )}
                </div>
              </div>
            </div>
            <button
              onClick={() => onDeleteTask(task.id)}
              className="opacity-0 group-hover:opacity-100 transition-opacity px-3 py-1 text-red-600 hover:bg-red-50 rounded-md"
            >
              Удалить
            </button>
          </div>
        </div>
      ))}
    </div>
  );
};

const getPriorityColor = (priority) => {
  switch (priority) {
    case 'Высокий':
      return 'bg-red-100 text-red-800';
    case 'Средний':
      return 'bg-yellow-100 text-yellow-800';
    case 'Низкий':
      return 'bg-green-100 text-green-800';
    default:
      return 'bg-gray-100 text-gray-800';
  }
};

const formatDate = (dateString) => {
  if (!dateString) return '';
  try {
    const date = new Date(dateString);
    if (isNaN(date.getTime())) return dateString;
    
    return new Intl.DateTimeFormat('ru-RU', {
      year: 'numeric',
      month: 'long',
      day: 'numeric'
    }).format(date);
  } catch (error) {
    console.error('Ошибка форматирования даты:', error);
    return dateString;
  }
};

export default TaskList; 