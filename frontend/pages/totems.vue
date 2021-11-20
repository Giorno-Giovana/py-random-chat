<template>
    <div>
        <div v-for="totem in totems" key="totem">
            <h1>{{ totem.discussion_topic }}</h1>
            {{ totem.description }}
            Users online: {{ parseInt(totem.num_attendees) }}
        </div>
    </div>
</template>

<script>
export default {
    data() {
        return { totems: [] };
    },
    async fetch() {
        this.totems = (
            await this.$fire.firestore.collection('totems').get()
        ).docs.map((doc) => doc.data());
    },
};
</script>
