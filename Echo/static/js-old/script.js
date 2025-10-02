function changeMessage () {
    const messageElement = document.getElementById("message");
    const currentMessage = messageElement.textContent;

    const alternateMessages = [
        "Hola de nuevo!",
        "Bienvenido de vuelta!",
        "Quetengas un buen dia!",
    ]

    let newMessage;

    do {
        newMessage = alternateMessages[Math.floor(Math.random()* alternateMessages.length)]
    } while (newMessage === currentMessage);

    messageElement.textContent = newMessage;

}
