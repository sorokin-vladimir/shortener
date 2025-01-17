document.body.addEventListener("htmx:afterSettle", function(event) {
  const target = event.target;
  if (target.id !== "partial-content") {
    return;
  }

  const newButton = document.getElementById("copy-short");
  if (!newButton) {
    return;
  }

  newButton.addEventListener("click", copyHandler)
})

function copyHandler() {
  if (!navigator.clipboard) {
    const input = document.getElementById("short-input");
    input.select();
    input.setSelectionRange(0, input.value.length);

    try {
        document.execCommand("copy");
    } catch (err) {
        console.error("Failed to copy text: ", err);
    }
    input.blur();
    return;
  }

  const textToCopy = document.getElementById("short").textContent;
  navigator.clipboard.writeText(textToCopy)
}
