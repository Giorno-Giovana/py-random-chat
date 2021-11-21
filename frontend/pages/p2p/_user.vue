<template>
  <div class="w-screen h-screen relative">
    <span class="me absolute top-0 left-0 z-10">
        <video-stream :stream="me" class="w-32"/>
    </span>
    <video-stream :stream="me" class="you" @finishCall="finishCall" :p2p="true"/>
  </div>
</template>

<script>
export default {
  name: "_user",
  data() {
    return {
      localStream: undefined,
      me: undefined
    }
  },
  async mounted() {
    this.localStream = await navigator.mediaDevices.getUserMedia({ video: true, audio: true })
    this.me = new MediaStream()
    this.localStream.getVideoTracks().forEach(t => this.me.addTrack(t))
    console.log('Ебануть сюда подключение к другому человеку по обазу и подобию, как в комнатах, но только с одним человеком')
  },
  methods: {
    finishCall() {
      console.log('Ебануть сюда редирект в корень')
    }
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

.me .video-container {
  width: 200px;
  height: 250px;
  @apply pl-0 pt-0;
}
.me .video-stream {
  @apply rounded-tl-none;
}
</style>
