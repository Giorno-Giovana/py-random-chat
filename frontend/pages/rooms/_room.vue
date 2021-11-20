<template>
  <div>
    <VideoStream :stream="localStream" />
    <div class="grid-cols-4">
      <VideoStream :stream="rs.stream" v-for="rs in remoteStreams" :key="rs.stream.id" />
    </div>
  </div>
</template>

<script>
import Peer from "peerjs";
import VideoStream from "../../components/video-stream.vue";

const peer = new Peer(undefined, {
  debug: 1,
});

export default {
  components: { VideoStream },
  data() {
    return {
      localStream: undefined,
      remoteStreams: [],
      id: peer.id,
      room: this.$route.params.room,
    };
  },
  async mounted() {
    this.openPeerConnection();
    this.localStream = await window.navigator.mediaDevices.getUserMedia({
      video: true,
      audio: true,
    });
    this.subscribeToIncomingCalls();
  },
  socket: undefined,
  methods: {
    openPeerConnection() {
      peer.on("open", async (id) => {
        console.log("Локальный id", id);
        this.id = id;
        this.$options.socket = new WebSocket(`ws://51.250.16.140:8080/api/room/join?pid=${id}&rid=${this.room}`);
        this.subscribeToSocket()
      });
    },
    subscribeToSocket() {
      this.$options.socket.onopen = () => {
        console.log("[open] Соединение установлено");
      };

      this.$options.socket.onmessage = event => {
        const data = JSON.parse(event.data)
        if (data.type === 1) {
          this.call(data.peer)
        }
        if (data.type === 2) {
          this.removeUserFromCall(data.peer)
        }
      };

      this.$options.socket.onclose = (event) => {
        if (event.wasClean) {
          console.log(`[close] Соединение закрыто чисто, код=${event.code} причина=${event.reason}`);
        } else {
          console.log('[close] Соединение прервано');
        }
      };

      this.$options.socket.onerror = (error) => {
        console.log(`[error] ${error.message}`);
      };
    },
    subscribeToIncomingCalls() {
      console.log("Подписка на входящие звонки");
      peer.on("call", (call) => {
        call.answer(this.localStream);
        call.on("stream", stream => this.addStreamToRemotes(stream, call.peer));
      });
    },
    call(p) {
      console.log("Зовнок", p);
      const call = peer.call(p, this.localStream);
      call.on("stream", stream => this.addStreamToRemotes(stream, p));
    },
    addStreamToRemotes(stream, p) {
      if (this.remoteStreams.some((r) => r.stream.id === stream.id)) return;
      this.remoteStreams.push({stream, peer: p});
    },
    removeUserFromCall(peer) {
      console.log('Удалён ', peer)
      this.remoteStreams = this.remoteStreams.filter(r => r.peer !== peer)
    },
  },
};
</script>
