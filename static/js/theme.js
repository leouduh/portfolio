const toggleButton = document.getElementById('theme-toggle');

toggleButton.addEventListener('click', () => {
    const isDark = document.documentElement.classList.toggle('dark');
    localStorage.setItem('theme', isDark ? 'dark': 'light');
});
