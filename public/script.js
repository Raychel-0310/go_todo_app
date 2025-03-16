document.addEventListener("DOMContentLoaded", function () {
    const taskList = document.getElementById("task-list");
    const taskInput = document.getElementById("task-input");
    const addTaskButton = document.getElementById("add-task-button");

    // APIからタスクを取得
    function fetchTasks() {
        fetch("http://localhost:8080/todos")
            .then(response => response.json())
            .then(data => {
                taskList.innerHTML = "";  // リストをリセット
                data.forEach(task => {
                    const li = document.createElement("li");

                    // 完了したタスクには横線を引くためにCSSを利用
                    const taskName = task.done ? `<span class="completed">${task.task}</span>` : task.task;

                    li.innerHTML = `
                        <span>${taskName}</span>
                        ${task.done ? "" : `<button onclick="editTask(${task.id})">編集</button>
                                            <button onclick="deleteTask(${task.id})">削除</button>
                                            <button onclick="toggleDone(${task.id}, ${task.done})">完了</button>`}
                    `;

                    taskList.appendChild(li);
                });
            })
            .catch(error => {
                console.error("タスクの取得に失敗しました:", error);
            });
    }

    // 新しいタスクを追加
    addTaskButton.addEventListener("click", function () {
        const newTask = taskInput.value.trim();
        if (newTask) {
            const task = { task: newTask, done: false };

            fetch("http://localhost:8080/todos", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify(task)
            })
                .then(response => response.json())
                .then(() => {
                    taskInput.value = "";  // 入力フィールドをリセット
                    fetchTasks(); // タスクを再取得して更新
                })
                .catch(error => {
                    console.error("タスクの追加に失敗しました:", error);
                });
        }
    });

    // タスクの編集
    window.editTask = function (id) {
        const newTask = prompt("新しいタスク名を入力してください:");
        if (newTask) {
            const updatedTask = { task: newTask, done: false };

            fetch(`http://localhost:8080/todos/${id}`, {
                method: "PUT",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify(updatedTask)
            })
                .then(response => response.json())
                .then(() => {
                    fetchTasks(); // タスクを再取得して更新
                })
                .catch(error => {
                    console.error("タスクの更新に失敗しました:", error);
                });
        }
    };

    // タスクの削除
    window.deleteTask = function (id) {
        if (confirm("本当に削除しますか？")) {
            fetch(`http://localhost:8080/todos/${id}`, {
                method: "DELETE"
            })
                .then(() => {
                    fetchTasks(); // タスクを再取得して更新
                })
                .catch(error => {
                    console.error("タスクの削除に失敗しました:", error);
                });
        }
    };

    // タスクの完了状態を切り替え
    window.toggleDone = function (id, currentStatus) {
        const updatedTask = { task: "", done: !currentStatus };

        fetch(`http://localhost:8080/todos/${id}`, {
            method: "PUT",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(updatedTask)
        })
            .then(response => response.json())
            .then(() => {
                fetchTasks(); // タスクを再取得して更新
            })
            .catch(error => {
                console.error("タスクの完了状態の切り替えに失敗しました:", error);
            });
    };

    // 初回にタスクを取得
    fetchTasks();
});
