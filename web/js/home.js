const ctx = document.getElementById('weight-chart')

async function fetchWeightData() {
    try {
        const weightEntriesResponse = await fetch('/api/daily-weight-entries?n=7')
        if (!weightEntriesResponse.ok) {
            throw new Error(`HTTP error! status: ${weightEntriesResponse.status}`)
        }
        const weightEntriesJSON = await weightEntriesResponse.json()
        const targetWeightEntryResponse = await fetch('/api/target-weight')
        if (!targetWeightEntryResponse.ok) {
            throw new Error(`HTTP error! status: ${targetWeightEntryResponse.status}`)
        }
        const targetWeightEntryJSON = await targetWeightEntryResponse.json()
        const sorted = weightEntriesJSON.sort((a, b) => new Date(a.entryDate) - new Date(b.entryDate))
        // Convert dates to MM/DD
        const labels = sorted.map(item => {
            const d = new Date(item.entryDate)
            return `${String(d.getUTCMonth() + 1).padStart(2, '0')}/${String(d.getUTCDate()).padStart(2, '0')}`
        })
        const weights = sorted.map(item => item.weight)
        const minWeight = Math.min(...weights) 
        const maxWeight = Math.max(...weights)
        const targetWeight = targetWeightEntryJSON.weight
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
                    legend: { display: false },
                    annotation: {
                        annotations: {
                            targetLine: {
                                type: 'line',
                                yMin: targetWeight,
                                yMax: targetWeight,
                                borderColor: 'red',
                                borderWidth: 2,
                                borderDash: [6, 6],
                                label: {
                                    content: 'Target',
                                    enabled: true,
                                    position: 'start',
                                    backgroundColor: 'red',
                                    color: 'white',
                                }
                            }
                        }
                    }
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
                        suggestedMin: Math.min(minWeight, targetWeight) - 5,
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