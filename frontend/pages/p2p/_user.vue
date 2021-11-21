<template>
  <div class="w-screen h-screen relative">
    <span class="me absolute top-0 left-0 z-10">
        <video-stream :stream="me" class="w-32"/>
    </span>
    <video-stream :stream="you" class="you" @finishCall="finishCall" :p2p="true"/>
  </div>
</template>

<script>
import Peer from "peerjs";

export default {
  name: "_user",
  data() {
    return {
      localStream: undefined,
      me: undefined,
      you: undefined
    }
  },
  peer: undefined,
  // TODO ебануть call на пира, который должен прилетать с роута, а роут
  // должен прилетать с бека с метча, доделаете кароч
  async mounted() {
    await this.getUserMedia()
    this.$options.peer = new Peer()
    console.log('Ебануть сюда подключение к другому человеку по обазу и подобию, как в комнатах, но только с одним человеком')
  },
  methods: {
    async getUserMedia() {
      this.localStream = await navigator.mediaDevices.getUserMedia({ video: true, audio: true })
      this.me = new MediaStream()
      this.localStream.getVideoTracks().forEach(t => this.me.addTrack(t))
    },
    finishCall() {
      console.log('Ебануть сюда редирект в корень')
    },
    subscribeToIncomingCalls() {
      console.log('Подписка на входящие звонки');
      true.$options.peer.on('call', (call) => {
        call.answer(this.localStream);
        call.on('stream', (stream) =>
          this.setRemoteStream(stream)
        );
      });
    },
    call(p) {
      console.log('Зовнок', p, !!this.localStream);
      if (this.localStream) {
        const call = this.$options.peer.call(p, this.localStream);
        call.on('stream', (stream) =>
          this.setRemoteStream(stream)
        );
      } else {
        this.getUserMedia().then(() => {
          this.call(p);
        });
      }
    },
    setRemoteStream(stream) {
      if (this.remoteStreams.some((r) => r.stream.id === stream.id)) return;
      this.you = stream;
    },
  }
}
</script>

<style>
.you {
  width: 100vw;
  height: 100vh;
  padding: 0;
}
.you .video-stream {
  @apply rounded-none
}
/* Что бы вам не говорила ide - она пиздит */
.me .video-container {
  width: 200px;
  height: 250px;
  @apply pl-0 pt-0;
}
.me .video-stream {
  @apply rounded-tl-none;
}
</style>
