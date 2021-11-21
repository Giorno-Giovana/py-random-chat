const getParticipant = (userId, participationMode) => {
    const socket = new WebSocket(`ws://sshamanism.ru:8080/api/tinder/next?uid=${userId}&participation_mode=${participationMode}`);
    socket.onmessage = (event) => console.log(event)
};