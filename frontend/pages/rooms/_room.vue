<template>
  <div>
    <div>Hi mark</div>
    <VideoStream :stream="localStream" />
    <VideoStream :stream="rs" v-for="rs in remoteStreams" :key="rs.id" />
    <input class="border border-black" type="text" v-model="callerId" />
    <button class="border border-black" @click="call">Call</button>
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
      callerId: "",
      id: peer.id,
    };
  },
  room: undefined,
  beforeDestroy() {
    peer.disconnect();
  },
  async mounted() {
    const rooms = await this.$fire.firestore
      .collection("rooms")
      .doc(this.$route.params.room);
    this.$options.peers = rooms.collection("peers");

    this.openPeerConnection();
    this.localStream = await window.navigator.mediaDevices.getUserMedia({
      video: true,
      audio: true,
    });
    this.subscribeToIncomingCalls();
  },
  methods: {
    call() {
      console.log("Попытка позвонить", this.callerId);
      const call = peer.call(this.callerId, this.localStream);
      call.on("stream", this.addStreamToRemotes);
    },
    openPeerConnection() {
      peer.on("open", async (id) => {
        this.$options.peers.add({ peer: this.id });
        this.id = id;
        console.log("Локальный id", this.id);
      });
    },
    subscribeToIncomingCalls() {
      console.log("Подписка на входящие звонки");
      peer.on("call", (call) => {
        console.log("Совершена попытка звонка", call);
        call.answer(this.localStream);
        call.on("stream", this.addStreamToRemotes);
      });
    },
    addStreamToRemotes(stream) {
      if (this.remoteStreams.some((s) => s.id === stream.id)) return;
      console.log("Оп, подрубочка", stream);
      this.remoteStreams.push(stream);
    },
  },
};
</script>
