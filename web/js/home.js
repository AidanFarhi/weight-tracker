const ctx = document.getElementById('weight-chart')

async function fetchWeightData() {
    try {
        const response = await fetch('/api/daily-weights')
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`)
        }
        const rawData = await response.json()
        const sorted = rawData.sort((a, b) => new Date(a.entryDate) - new Date(b.entryDate))
        // Convert dates to MM/DD
        const labels = sorted.map(item => {
            const d = new Date(item.entryDate + "T00:00:00")
            return `${(d.getMonth() + 1).toString().padStart(2, '0')}/${d.getDate().toString().padStart(2, '0')}`
        })
        const weights = sorted.map(item => item.weight)
        const minWeight = Math.min(...weights) 
        const maxWeight = Math.max(...weights)
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
                borderColor: 'transparent',
                plugins: {
                    legend: { display: false }
                },
                scales: {
                    x: {
                        grid: {
                            display: true,
                            color: 'transparent',
                            drawBorder: false,
                            drawTicks: false
                        },
                        border: {
                            display: false,
                            color: 'transparent'
                        },
                        ticks: {
                            color: '#666'
                        }
                    },
                    y: {
                        grid: {
                            display: true,
                            color: 'grey',
                            drawBorder: false,
                            drawTicks: true
                        },
                        suggestedMin: minWeight - 5,
                        suggestedMax: maxWeight + 5,
                        border: {
                            display: false,
                            color: 'transparent'
                        },
                        ticks: {
                            stepSize: 5
                        }
                    }
                }
            }

        })
    } catch (error) {
        console.error('Error fetching weight data:', error)
    }
}

fetchWeightData()