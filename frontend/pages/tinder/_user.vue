<template>
    <v-card class="h-screen w-screen">
        <v-img
            height="350"
            src="https://pbs.twimg.com/media/Er7fRCmVoAAOz06.jpg"
        />

        <v-card-title>{{ name }}</v-card-title>

        <v-card-text>
            <v-chip v-for="skill in skills">{{ skill }}</v-chip>
            <div class="mt-4">{{ description }}</div>
        </v-card-text>

        <v-divider class="mx-4" />
        <v-card-title>Links</v-card-title>
        <div class="px-4 grid grid-cols-2 gap-y-4">
            <div>
                <v-icon> mdi-github </v-icon>
                <a v-bind:href="'https://github.com/' + github">{{ github }}</a>
            </div>
            <div>
                <v-icon> mdi-facebook </v-icon>
                <a v-bind:href="'https://ru-ru.facebook.com/' + facebook">{{
                    facebook
                }}</a>
            </div>
            <div>
                <v-icon> mdi-linkedin </v-icon>
                <a href="#">linkedin</a>
            </div>
        </div>
        <div class="flex justify-around absolute bottom-0 w-full h-24">
            <div class="shadow-lg actionButton">
                <v-icon class="actionIcon"> mdi-close </v-icon>
            </div>

            <div class="bg-red-400 actionButton">
                <v-icon class="actionIcon"> mdi-heart </v-icon>
            </div>
        </div>
    </v-card>
</template>

<script>
export default {
    name: 'match',

    data() {
        return {
            name: 'Blank Blankivich',
            description:
                'I had my 15 minutes of being the new boy of pop, like lots of people before and after me. Overnight, everyone starts treating you differently, and perceives you differently.',
            skills: ['Vue', 'NodeJS', 'Golang', 'Nonassociative algebra'],
            github: 'deniskamazur',
            facebook: 'deniskamazur',
            linkedin: 'deniskamazur',
            occupation: 'SWE',
            place_of_work: 'Brelluh',
            participation_mode: 'online',
        };
    },

    async fetch() {},
};
</script>

<style scoped>
.actionButton {
    @apply rounded-full w-20 h-20 text-center;
}

.actionIcon {
    @apply pt-7 flex align-middle justify-center;
}
</style>
