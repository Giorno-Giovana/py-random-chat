<template>
    <div>
        <h1 class="text-center text-4xl my-16">Available rooms</h1>
        <v-divider class="my-10 mx-8" />
        <div class="flex flex-wrap">
            <NuxtLink v-for="r in rooms" :to="r.link" :key="r.stream.id">
                <video-stream :stream="r.stream" class="m-5">
                    <template v-slot:tooltip>
                        <div
                            class="w-full mt-2 rounded-lg px-4 py-1 z-20"
                            :class="r.color"
                        >
                            Users in room: {{ r.num_participants }}
                        </div>
                    </template>
                </video-stream>
            </NuxtLink>
        </div>
    </div>
</template>

<script>
import Peer from 'peerjs';

export default {
    data() {
        return {
            rooms: [],
            localStream: undefined,
            me: undefined,
        };
    },
    async mounted() {
        await this.getUserMedia();
        this.$options.peer = new Peer();
    },
    methods: {
        // Это чтобы локально видосы для тестов были, так то он не нужен
        async getUserMedia() {
            this.localStream = await navigator.mediaDevices.getUserMedia({
                video: true,
                audio: true,
            });
            const stream = new MediaStream();
            this.localStream
                .getVideoTracks()
                .forEach((t) => stream.addTrack(t));
            const colors = ['blue', 'red', 'yellow', 'green', 'purple', 'pink'];

            const roomList = (
                await (
                    await fetch('https://sshamanism.ru/api/api/room/list')
                ).json()
            ).list;

            this.rooms = this.rooms.concat([
                {
                    link: '/rooms/1',
                    stream,
                    color: `bg-${colors[Math.floor(Math.random() * 6)]}-400`,
                    num_participants: roomList[0].num_participants,
                },
                {
                    link: '/rooms/2',
                    stream,
                    color: `bg-${colors[Math.floor(Math.random() * 6)]}-400`,
                    num_participants: roomList[1].num_participants,
                },
                {
                    link: '/rooms/3',
                    stream,
                    color: `bg-${colors[Math.floor(Math.random() * 6)]}-400`,
                    num_participants: roomList[2].num_participants,
                },
            ]);
        },
    },
};
</script>
