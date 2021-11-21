<template>
  <div>
    <div class="grid grid-cols-3">
      <div class="col-span-2" v-if="remoteStreams.length">
        <VideoStream
          :stream="remoteStreams[0].stream"
          :emotion="remoteStreams[0].emotion"
          style="width: 100%; height: 100vh; margin: 0"
        />
      </div>
      <div class="flex md:justify-around xl:justify-center flex-wrap">
        <VideoStream
          :stream="webcam"
          @mute="toggleSelfMute"
          :webcam="true"
          @emotion="showSelfEmotion"
          :emotion="selfEmotion"
        />
        <VideoStream
          v-for="rs in remoteStreams.slice(1)"
          :stream="rs.stream"
          :emotion="rs.emotion"
          :key="rs.stream.id"
        />
      </div>
    </div>
  </div>
</template>

<script>
import Peer from "peerjs";
import VideoStream from "../../components/video-stream.vue";

let peer;

export default {
  components: { VideoStream },
  layout: false,
  data() {
    return {
      selfEmotion: "",
      webcam: undefined,
      localStream: undefined,
      remoteStreams: [],
      id: undefined,
      room: this.$route.params.room,
      isSelfMuted: true,
    };
  },
  async mounted() {
    await this.getUserMedia();
    peer = new Peer(undefined, {
      debug: 1,
    });
    this.id = peer.id;
    this.openPeerConnection();
    this.subscribeToIncomingCalls();
  },
  socket: undefined,
  methods: {
    async showSelfEmotion(emotion) {
      this.selfEmotion = "";
      await this.$nextTick();
      switch (emotion) {
        case "dislike":
          this.selfEmotion = "/dislike.png";
          this.$options.socket.send("/dislike.png");
          break;
        case "like":
          this.selfEmotion = "/heart.png";
          this.$options.socket.send("/heart.png");
      }
    },
    async toggleSelfMute(muted) {
      this.localStream.getAudioTracks()[0].enabled = !muted;
    },
    async getUserMedia() {
      this.localStream = await window.navigator.mediaDevices.getUserMedia({
        video: true,
        audio: true,
      });

      this.webcam = new MediaStream();
      this.localStream.getTracks().forEach((track) => {
        if (track.kind === "audio") return;
        this.webcam.addTrack(track);
      });
    },
    openPeerConnection() {
      peer.on("open", async (id) => {
        console.log("Локальный id", id);
        this.id = id;
        this.$options.socket = new WebSocket(
          `ws://51.250.16.140:8080/api/room/join?pid=${id}&rid=${this.room}`
        );
        this.subscribeToSocket();
      });
    },
    subscribeToSocket() {
      this.$options.socket.onopen = () => {
        console.log("[open] Соединение установлено");
      };

      this.$options.socket.onmessage = async (event) => {
        const data = JSON.parse(event.data);
        console.log("[SOCKET]", data);
        if (data.type === 1) {
          this.call(data.peer);
        }
        if (data.type === 2) {
          this.removeUserFromCall(data.peer);
        }
        if (data.type === 3) {
          const screamerIndex = this.remoteStreams.findIndex(
            (s) => s.peer === data.peer
          );
          const screamerValue = this.remoteStreams[screamerIndex];
          this.$set(this.remoteStreams, screamerIndex, {
            ...screamerValue,
            emotion: "",
          });
          await this.$nextTick();
          this.$set(this.remoteStreams, screamerIndex, {
            ...screamerValue,
            emotion: data.message,
          });
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
        call.on("stream", (stream) =>
          this.addStreamToRemotes(stream, call.peer)
        );
      });
    },
    call(p) {
      console.log("Зовнок", p, !!this.localStream);
      if (this.localStream) {
        const call = peer.call(p, this.localStream);
        call.on("stream", (stream) => this.addStreamToRemotes(stream, p));
      } else {
        this.getUserMedia().then(() => {
          this.call(p);
        });
      }
    },
    addStreamToRemotes(stream, p) {
      if (this.remoteStreams.some((r) => r.stream.id === stream.id)) return;
      this.remoteStreams.push({ stream, peer: p, emotion: "" });
    },
    removeUserFromCall(peer) {
      console.log("Удалён ", peer);
      this.remoteStreams = this.remoteStreams.filter((r) => r.peer !== peer);
    },
  },
};
</script>
