function changeMessage() {
    const messageElement = document.getElementById("message");
    const currentMessage = messageElement.textContent;
    const alternateMessage = [
        "¡Hola de Nuevo!",
        "¡Bienvenido de vuelta!",
        "¡Que tengas un buen día!"
    ];

    let newMessage;
    do {
        newMessage = alternateMessage[Math.floor(Math.random() * alternateMessage.length)];
    } while (newMessage === currentMessage);
    messageElement.textContent = newMessage;
}