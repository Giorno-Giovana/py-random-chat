<template>
  <div
    @mouseenter="isToolbarVisible = true"
    @mouseleave="isToolbarVisible = false"
    class="video-container"
  >
    <img v-if="emotion" :src="emotion" class="emotion absolute z-20 scale-75" />
    <img
      :src="image"
      class="pulse w-full h-full left-0 top-0 absolute"
      :style="animation"
      alt=""
    />
    <div class="insider">
      <video
        @click="shoot"
        ref="video"
        autoplay
        class="w-full h-full object-cover inline rounded-3xl"
      />
      <div v-if="isToolbarVisible" class="absolute bottom-5 w-full flex">
        <v-icon
          class="bg-white rounded-full p-2 mx-auto"
          color="red"
          @click="thumbDown"
          v-if="webcam"
        >
          mdi-thumb-down
        </v-icon>
        <v-icon class="bg-white rounded-full p-2 mx-auto" @click="toggleMute">
          {{ muted ? "mdi-microphone-off" : "mdi-microphone" }}
        </v-icon>
        <v-icon
          class="bg-white rounded-full p-2 mx-auto"
          color="green"
          @click="thumbUp"
          v-if="webcam"
        >
          mdi-thumb-up
        </v-icon>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: ["stream", "webcam", "emotion"],
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
    async shoot(e) {
      var video = this.$refs.video;
      // const scale_factor = 0.33;

      // var rect = video.getBoundingClientRect();
      // var len = Math.max(video.videoWidth, video.videoHeight) * scale_factor;

      // var x = e.clientX - rect.left;
      // var y = e.clientY - rect.top;

      // var dx = x + len / 2;
      // var dy = y;

      var canvas = document.createElement("canvas");
      canvas.width = video.videoWidth;
      canvas.height = video.videoHeight;
      var ctx = canvas.getContext("2d");
      ctx.drawImage(video, 0, 0);
      // ctx.drawImage(video, dx, dy, len, len, 0, 0, len, len);

      var base64image = canvas.toDataURL();
      base64image = base64image.replace(/^data:image\/[a-z]+;base64,/, "");

      this.$recognition.identifyBase64Image(base64image).then((username) => {
        console.log(username);
      });
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
    thumbUp() {
      this.$emit("emotion", "like");
    },
    thumbDown() {
      this.$emit("emotion", "dislike");
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

.emotion {
  animation-duration: 2000ms;
  animation-name: like-heart-animation;
  animation-timing-function: revert;
  visibility: hidden;
  position: absolute;
  top: 0;
  left: 0;
}

@keyframes like-heart-animation {
  0% {
    opacity: 0;
    transform: scale(0);
    visibility: visible;
  }
  15% {
    opacity: 0.9;
    transform: scale(1.2);
  }
  30% {
    transform: scale(0.95);
  }
  45%,
  80% {
    opacity: 0.9;
    transform: scale(1);
  }
  100% {
    transform: scale(0);
  }
}
</style>
