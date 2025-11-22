const ctx = document.getElementById('weight-chart');

// Fetch data from your API
async function fetchWeightData() {
    try {
        const response = await fetch('/api/daily-weights'); // Replace with your API endpoint
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        const rawData = await response.json();

        // Sort chronologically
        const sorted = rawData.sort((a, b) => new Date(a.entryDate) - new Date(b.entryDate));

        // Convert dates to MM/DD
        const labels = sorted.map(item => {
            const d = new Date(item.entryDate);
            return `${(d.getMonth() + 1).toString().padStart(2, '0')}/${d.getDate().toString().padStart(2, '0')}`;
        });

        const weights = sorted.map(item => item.weight);

        const minWeight = Math.min(...weights);
        const maxWeight = Math.max(...weights);

        // Create the chart
        new Chart(ctx, {
            type: 'line',
            data: {
                labels: labels,
                datasets: [{
                    label: 'Weight',
                    data: weights,
                    borderWidth: 3.5,
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
                        suggestedMin: minWeight - 5,
                        suggestedMax: maxWeight + 5,
                        ticks: {
                            stepSize: 5
                        },
                    }
                }
            }
        });
    } catch (error) {
        console.error('Error fetching weight data:', error);
    }
}

// Call the function to fetch data and render chart
fetchWeightData();