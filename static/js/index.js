const backButton = document.getElementById('back-button');

if (backButton) {
  backButton.addEventListener('click', () => {
    window.history.back();
  });
}
