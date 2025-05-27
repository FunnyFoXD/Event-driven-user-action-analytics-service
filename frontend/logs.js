function formatDate(isoString) {
    const date = new Date(isoString);
    return date.toLocaleString();
}

async function fetchLogs() {
    try {
        const response = await fetch('http://localhost:8080/logs');
        const logs = await response.json();

        const tableBody = document.querySelector('#logs-table tbody');

        // if (!tableBody) {
        //     console.error("Table body not found! Check HTML structure.");
        //     return;
        // }

        tableBody.innerHTML = ''; // Clear

        logs.forEach(log => {
            const row = document.createElement('tr');
            row.innerHTML = `
                <td>${formatDate(log.CreatedAt)}</td>
                <td>${log.user_id}</td>
                <td>${log.action}</td>
            `;
            tableBody.appendChild(row);
        });
    } catch (error) {
        console.error('Error download logs:', error);
    }
}

document.addEventListener('DOMContentLoaded', fetchLogs);
