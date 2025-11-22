const ctx = document.getElementById('weight-chart')

const rawData = [
    { id: 3, user_id: 1, weight: 150, category: "daily", entry_date: "2025-11-02" },
    { id: 2, user_id: 1, weight: 149, category: "daily", entry_date: "2025-11-03" },
    { id: 4, user_id: 1, weight: 151, category: "daily", entry_date: "2025-11-05" },
    { id: 4, user_id: 1, weight: 147, category: "daily", entry_date: "2025-11-06" },
    { id: 4, user_id: 1, weight: 143, category: "daily", entry_date: "2025-11-07" }
];

// Convert dates to MM/DD format and sort them chronologically
const sorted = rawData.sort((a, b) => new Date(a.entry_date) - new Date(b.entry_date));

const labels = sorted.map(item => {
    const d = new Date(item.entry_date);
    return `${(d.getMonth() + 1).toString().padStart(2, '0')}/${d.getDate().toString().padStart(2, '0')}`
});

const weights = sorted.map(item => item.weight);

const minWeight = Math.min(...weights);
const maxWeight = Math.max(...weights);

new Chart(ctx, {
    type: 'line',
    data: {
        labels: labels,
        datasets: [{
            label: 'Weight',
            data: weights,
            borderWidth: 3.5,
            // tension: 0.1, // slight curve to the line
            fill: false,
            pointRadius: 5,
            borderColor: '#4C82FF',
            pointBackgroundColor: '#4C82FF',
            pointBorderColor: '#4C82FF',
        }]
    },
    options: {
        plugins: {
            legend: {
                display: false
            }
        },
        scales: {
            y: {
                beginAtZero: false,
                grid: {
                    display: true,
                    color: 'grey'
                },
                suggestedMin: minWeight - 5,   // buffer below
                suggestedMax: maxWeight + 5,   // buffer above
                ticks: {
                    stepSize: 5                // increment every 5 lbs
                },
            }
        }
    }
});