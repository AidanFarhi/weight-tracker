// const weightInput = document.getElementById('daily-weight-entry-weight-input')

// const MAX_WEIGHT_VALUE = 600

// let lastValidValue = ''

// weightInput.addEventListener('input', e => {
//     if (weightInput.value.length == 0 && e.data == '.') {
//         weightInput.value = ''
//         return
//     }
//     if (weightInput.value.length > 0 && e.data == '.') {
//         if (Number(weightInput.value) == MAX_WEIGHT_VALUE) {
//             weightInput.value = weightInput.value.split('.')[0]
//             const end = weightInput.value.length
//             weightInput.setSelectionRange(end, end)
//             return
//         }
//     }
//     if (Number(weightInput.value) > MAX_WEIGHT_VALUE) {
//         weightInput.value = lastValidValue
//         return
//     }
//     lastValidValue = weightInput.value
// })
