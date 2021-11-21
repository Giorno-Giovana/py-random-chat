<template>
  <div
    @mouseenter="isToolbarVisible = true"
    @mouseleave="isToolbarVisible = false"
    class="video-container"
  >
    <img
      :src="image"
      class="pulse w-full h-full left-0 top-0 absolute"
      :style="animation"
      alt=""
    />
    <div class="insider">
      <video
        @click="handleVideoClick"
        ref="video"
        autoplay
        class="w-full h-full object-cover inline rounded-3xl"
      />
      <div v-if="isToolbarVisible" class="absolute bottom-5 w-full flex">
        <v-icon class="bg-white rounded-full p-1 mx-auto" @click="toggleMute">
          {{ muted ? "mdi-microphone-off" : "mdi-microphone" }}
        </v-icon>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: ["stream"],
  data() {
    return {
      muted: false,
      isToolbarVisible: false,
      image: `/Vector-${Math.floor(Math.random() * 20)}.png`,
      animation: `animation-duration: ${Math.floor(Math.random() * 10) + 60}s;`,
    };
  },
  mounted() {
    if (this.stream) {
      this.setStream(this.stream);
    }
  },
  methods: {
    async handleVideoClick(e) {
      var video = e.target;
      const scale_factor = 0.33;

      var rect = video.getBoundingClientRect();
      var len = video.videoWidth * scale_factor;

      var x = e.clientX - rect.left;
      var y = e.clientY - rect.top;

      var dx = x + len / 2;
      var dy = y;

      var canvas = document.createElement("canvas");
      canvas.width = len;
      canvas.height = len;
      var ctx = canvas.getContext("2d");
      ctx.drawImage(video, dx, dy, len, len, 0, 0, len, len);

      var base64image = canvas.toDataURL();
      base64image = base64image.replace(/^data:image\/[a-z]+;base64,/, "");

      const username = await this.$recognition.identifyBase64Image(base64image);
      console.log(username);
    },

    toggleMute() {
      if (!this.muted) {
        this.muteVideo();
      } else {
        this.unMuteVideo();
      }

      this.muted = !this.muted;
      this.$emit("mute", this.muted);
    },
    setStream(stream) {
      this.$refs.video.srcObject = stream;
    },
    muteVideo() {
      const stream = this.$refs.video.srcObject;
      const newStream = new MediaStream();
      stream.getVideoTracks().forEach((t) => newStream.addTrack(t));
      this.setStream(newStream);
    },
    unMuteVideo() {
      this.setStream(this.stream);
    },
  },
  watch: {
    stream(newStream) {
      this.setStream(newStream);
    },
  },
};
</script>
<style>
.video-container {
  width: 300px;
  height: 300px;
  position: relative;
  display: flex;
  padding: 30px;
}

.insider {
  width: 100%;
  height: 100%;
  position: relative;
}

.pulse {
  width: 40rem;
  animation: pulsar cubic-bezier(0.4, 0, 0.6, 0.7) infinite;
}

@keyframes pulsar {
  0% {
    transform: rotate(0deg);
    opacity: 0.7;
  }
  20%,
  80% {
    opacity: 1;
  }
  50% {
    opacity: 0.7;
    transform: rotate(60deg);
  }
  100% {
    transform: rotate(0deg);
    opacity: 0.7;
  }
}
</style>
