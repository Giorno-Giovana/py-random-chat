<template>
  <div class="content">
    <div class="form">
      <div class="field-wrapper">
        <label class="label">Email</label>
        <input
          type="email"
          class="input"
          name="email"
          v-model="email"
          required
        />
      </div>
      <div class="field-wrapper">
        <label class="label">Password</label>
        <input
          type="password"
          class="input"
          name="password"
          v-model="password"
          required
        />
      </div>
      <button @click="submit" id="submit" type="submit">Log in</button>
    </div>
  </div>
</template>

<script>
export default {
  data: function () {
    return {
      email: null,
      password: null,
    };
  },
  methods: {
    async submit() {
      try {
        await this.$fire.auth.signInWithEmailAndPassword(
          this.email,
          this.password
        );
        this.$router.push({ path: "/" });
      } catch (e) {
        alert(e);
      }
    },
  },
};
</script>
