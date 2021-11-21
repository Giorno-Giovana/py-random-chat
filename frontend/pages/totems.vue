<template>
  <div>
    <h1 class="text-center text-4xl my-16">Available rooms</h1>
    <v-divider class="my-10 mx-8"/>
    <div class="flex flex-wrap">
      <NuxtLink
        v-for="r in rooms"
        :to="r.link"
        :key="r.stream.id"
      >
        <video-stream
          :stream="r.stream"
          class="m-5"
        >
          <template v-slot:tooltip>
            <div
              class="w-full mt-2 rounded-lg px-4 py-1 z-20"
              :class="r.color"
            >
              Users in room: 123
            </div>
          </template>
        </video-stream>
      </NuxtLink>

    </div>
  </div>
</template>

<script>
import Peer from "peerjs";

export default {
  data() {
    return {
      rooms: [],
      localStream: undefined,
      me: undefined,
    };
  },
  async mounted() {
    await this.getUserMedia()
    this.$options.peer = new Peer()
  },
  methods: {
    // Это чтобы локально видосы для тестов были, так то он не нужен
    async getUserMedia() {
      this.localStream = await navigator.mediaDevices.getUserMedia({video: true, audio: true})
      const stream = new MediaStream()
      this.localStream.getVideoTracks().forEach(t => stream.addTrack(t))
      const colors = [
          'blue',
          'red',
          'yellow',
          'green',
          'purple',
          'pink'
        ]
      //TODO Добавить сюда урлы
      this.rooms = this.rooms.concat(
        [
          {link: '#', stream, color: `bg-${colors[Math.floor(Math.random() * 6)]}-400`},
          {link: '#', stream, color: `bg-${colors[Math.floor(Math.random() * 6)]}-400`},
          {link: '#', stream, color: `bg-${colors[Math.floor(Math.random() * 6)]}-400`},
          {link: '#', stream, color: `bg-${colors[Math.floor(Math.random() * 6)]}-400`},
          {link: '#', stream, color: `bg-${colors[Math.floor(Math.random() * 6)]}-400`},
          {link: '#', stream, color: `bg-${colors[Math.floor(Math.random() * 6)]}-400`},
        ])
    },
  }
};
</script>
