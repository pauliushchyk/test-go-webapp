const backButton = document.getElementById('back-button');

if (backButton) {
  backButton.addEventListener('click', () => {
    window.history.back();
  });
}

const inputs = document.querySelectorAll('input[name=birth_date]');

if (inputs) {
  inputs.forEach(input => {
    const date = new Date();

    input.setAttribute('max', `${date.getFullYear() - 18}-12-31`);

    input.setAttribute('min', `${date.getFullYear() - 60}-01-01`);
  });
}
