// 完了タスク表示フラグ
let showCompleted = true;

// チェックボックスをクリックした時に完了状態を更新する関数
function toggleComplete(todoID) {
    const checkbox = document.getElementById(`check-${todoID}`);
    const completed = checkbox.checked;

    // エンドポイントに送信
    fetch(`/todo/complete/${todoID}`, {
        method: 'PATCH',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ completed: completed })
    })
    .then(response => response.json())
    .then(data => {
        // 完了状態に応じて、取り消し線を適用
        const label = document.getElementById(`label-${todoID}`);
        if (completed) {
            label.style.textDecoration = "line-through";
        } else {
            label.style.textDecoration = "none";
        }

        // data-completed 属性を更新
        const container = document.getElementById(`todo-${todoID}`);
        container.setAttribute("data-completed", completed ? "true" : "false");
    })
    .catch(error => {
        console.error('Error:', error);
    });
}

// タスクを削除する関数
function deleteTodo(todoID) {
    // エンドポイントに送信
    fetch(`/todo/delete/${todoID}/`, {
        method: 'DELETE',
        headers: {
            'Content-Type': 'application/json',
        },
    })
    .then(res => res.json())
    .then(data => {
        // 削除状況に応じて、要素を削除
        if (data.deleted) {
            document.getElementById('todo-' + todoID).remove();
        }
    });
}

function toggleCompleted() {
    showCompleted = !showCompleted;
    document.getElementById("toggleCompleted").innerText = showCompleted ? "Hide Completed" : "Show All";

    document.querySelectorAll('[data-completed="true"]').forEach(el => {
        el.classList.toggle("d-none", !showCompleted);
    });
}