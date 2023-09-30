document.addEventListener("DOMContentLoaded", function () {
    const dynamicContent = document.getElementById("dynamic-content");
    const changeTextButton = document.getElementById("change-text");
 
    changeTextButton.addEventListener("click", function () {
        const randomText = generateRandomText();
        dynamicContent.textContent = randomText;
    });
 
    function generateRandomText() {
        const characters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789";
        const textLength = 10; // Change this to the desired length of random text
        let randomText = "";
 
        for (let i = 0; i < textLength; i++) {
            const randomIndex = Math.floor(Math.random() * characters.length);
            randomText += characters.charAt(randomIndex);
        }
 
        return randomText;
    }
 });
 