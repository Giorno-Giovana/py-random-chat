<template>
    <span><video ref='video' autoplay class="inline-block" /></span>
</template>

<script>
export default {
    props: ['stream'],
    mounted() {
        if (this.stream){
            this.$refs.video.srcObject = this.stream
        }
    },
    watch: {
        stream(newStream) {
            this.$refs.video.srcObject = newStream
        }
    }
}
</script>