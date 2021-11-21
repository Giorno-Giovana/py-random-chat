<template>
  <div class="header">
    <div class="logo">
      <NuxtLink :to="{ name: 'index' }">duck</NuxtLink>
    </div>

    <div class="user">
      <template v-if="isLoggedIn">
        <img class="profile-photo" :src="authUser.photoURL" alt="img" />
        <div class="info">
          <a class="user" href="#"> {{ authUser.username }} </a>
          <div class="list">
            <button @click="signOut">Sign out</button>
          </div>
        </div>
      </template>
      <template v-else>
        <div class="info">
          <div class="list">
            <NuxtLink :to="{ name: 'signin' }">Sign in</NuxtLink>
            <NuxtLink :to="{ name: 'signup' }">Sign up</NuxtLink>
          </div>
        </div>
      </template>
    </div>
  </div>
</template>

<script>
import { mapState, mapGetters } from "vuex";

export default {
  computed: {
    ...mapState({
      authUser: (state) => state.authUser,
    }),
    ...mapGetters({
      isLoggedIn: "isLoggedIn",
    }),
  },
  methods: {
    async signOut() {
      try {
        this.$fire.firestore
          .collection("users")
          .doc(this.authUser.uid)
          .update({ is_online: false });
        await this.$fire.auth.signOut();
      } catch (e) {
        alert(e);
      }
    },
  },
};
</script>
