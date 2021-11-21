const getParticipant = (userId, participationMode) => {
    const socket = new WebSocket(`wss://sshamanism.ru:8080/api/api/tinder/next?uid=${userId}&participation_mode=${participationMode}`);
    socket.onmessage = (event) => console.log(event)
};