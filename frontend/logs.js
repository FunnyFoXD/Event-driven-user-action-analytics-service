async function fetchLogs() {
    try {
        const response = await fetch('http://localhost:8080/logs');
        const logs = await response.json();

        const tableBody = document.getElementById('logs-body');
        tableBody.innerHTML = ''; // Clear

        logs.forEach(log => {
            const row = document.createElement('tr');

            const dateCell = document.createElement('td');
            dateCell.textContent = log.CreatedAt;

            const userCell = document.createElement('td');
            userCell.textContent = log.user_id;

            const actionCell = document.createElement('td');
            actionCell.textContent = log.action;

            row.appendChild(dateCell);
            row.appendChild(userCell);
            row.appendChild(actionCell);

            tableBody.appendChild(row);
        });
    } catch (err) {
        console.error('Error download logs:', err);
    }
}

document.addEventListener('DOMContentLoaded', fetchLogs);
