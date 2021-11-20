<template>
  <div class="content">
    <div>
      <div class="field-wrapper">
        <label>Username</label>
        <input
          type="text"
          class="input"
          name="username"
          v-model="username"
          required
        />
      </div>
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
        <input type="file" @change="imageInput" accept="image/*" />
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
      <button @click="register" id="submit" type="submit">Register</button>
    </div>
  </div>
</template>
<script>
export default {
  data: function () {
    return {
      email: "",
      username: "",
      password: "",
      image: null,
      imageURL: null,
    };
  },
  methods: {
    async imageInput(e) {
      e.preventDefault();
      this.image = e.target.files[0];
    },

    async register() {
      console.log(this.$fire.auth);
    },
  },
};
</script>
