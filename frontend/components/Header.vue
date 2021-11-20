<template>
  <div class="header">
    <div class="logo">
      <NuxtLink :to="{ name: 'index' }">duck</NuxtLink>
    </div>

    <div class="user">
      <template v-if="$store.state.user">
        <img
          class="profile-photo"
          :src="$store.state.user.photoURL"
          alt="img"
        />
        <div class="info">
          <a class="user" href="#">{{ $store.state.user.displayName }}</a>
          <div class="list">
            <a @click="signOut" href="#">Log out</a>
          </div>
        </div>
      </template>
      <template v-else>
        <div class="info">
          <div class="list">
            <NuxtLink :to="{ name: 'login' }">Log in</NuxtLink>
            <NuxtLink :to="{ name: 'register' }">Register</NuxtLink>
          </div>
        </div>
      </template>
    </div>
  </div>
</template>

<script>
import { mapGetters } from "vuex";

export default {
  computed: {
    ...mapGetters({
      user: "user",
    }),
  },

  methods: {
    signOut: function (err) {
      this.$store.dispatch("signOut").catch((err) => {
        alert(err.message);
      });
    },
  },
};
</script>