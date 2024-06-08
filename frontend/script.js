const handle = document.getElementById('handle');
const value = document.getElementById('value');

let isDragging = false;

handle.addEventListener('mousedown', (e) => {
  isDragging = true;
  e.preventDefault();
});

document.addEventListener('mousemove', (e) => {
  if (isDragging) {
    const counter = document.getElementById('counter');
    const rect = counter.getBoundingClientRect();
    const newPosition = e.clientX - rect.left;
    const maxValue = counter.offsetWidth - handle.offsetWidth;

    const newValue = Math.max(0, Math.min(newPosition, maxValue));
    handle.style.left = newValue + 'px';
    value.textContent = Math.round((newValue / maxValue) * 100);
  }
});

document.addEventListener('mouseup', () => {
  isDragging = false;
});
