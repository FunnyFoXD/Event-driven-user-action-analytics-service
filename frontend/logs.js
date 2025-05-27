function formatDate(isoString) {
    const date = new Date(isoString);
    return date.toLocaleString();
}

async function fetchLogs() {
    try {
        const response = await fetch('http://localhost:8080/logs');
        const logs = await response.json();

        const tableBody = document.querySelector('#logs-table tbody');

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

async function submitEventForm(e) {
    e.preventDefault();

    const user_id = document.getElementById('user_id').value.trim();
    const action = document.getElementById('action').value.trim();

    if (!user_id || !action) {
        alert('Please fill in all fields!');
        return;
    }

    try {
        const response = await fetch('http://localhost:8080/event', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ user_id, action })
        });

        if (!response.ok) {
            throw new Error('Failed to add event');
        }

        alert('Event added successfully');
        document.getElementById('event-form').reset();
        fetchLogs();
    } catch (error) {
        alert('Error adding event: ' + error.message);
    }
}

document.addEventListener('DOMContentLoaded', () => {
    fetchLogs();

    const eventForm = document.getElementById('event-form');
    if (eventForm) {
        eventForm.addEventListener('submit', submitEventForm);
    } else {
        console.warn("No event form found on page.");
    }
});
