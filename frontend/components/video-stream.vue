<template>
  <v-card
    @mouseenter="isToolbarVisible = true"
    @mouseleave="isToolbarVisible = false"
    class="m-2 w-96 h-96 object-cover inline"
  >
    <video ref='video' autoplay class="w-96 h-96 object-cover inline" />
    <div v-if="isToolbarVisible" class="absolute bottom-5 w-full flex">
        <v-icon
          class="bg-white rounded-full p-1 mx-auto"
          @click="toggleMute"
        >
          {{ muted ? 'mdi-microphone-off' : 'mdi-microphone' }}
        </v-icon>
    </div>
  </v-card>
</template>

<script>
export default {
    props: ['stream'],
    data() {
      return {
        muted: false,
        isToolbarVisible: false,
      }
    },
    mounted() {
      if (this.stream) {
        this.setStream(this.stream)
      }
    },
    methods:{
      toggleMute() {

        if (!this.muted) {
          this.muteVideo()
        } else {
          this.unMuteVideo()
        }

        this.muted = !this.muted
        this.$emit('mute', this.muted)
      },
      setStream(stream) {
        this.$refs.video.srcObject = stream
      },
      muteVideo() {
        const stream = this.$refs.video.srcObject
        const newStream = new MediaStream()
        stream.getVideoTracks().forEach(t => newStream.addTrack(t))
        this.setStream(newStream)
      },
      unMuteVideo() {
        this.setStream(this.stream)
      }
    },
    watch: {
      stream(newStream) {
        this.setStream(newStream)
      },
    }
}
</script>
