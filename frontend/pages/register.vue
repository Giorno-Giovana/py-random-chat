<template>
  <div class="content">
    <div class="form">
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
        <input type="file" @change="fileInput" accept="image/*" />
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
      photo: null,
      photoURL: null,
    };
  },
  methods: {
    async fileInput(e) {
      e.preventDefault();
      this.photo = e.target.files[0];
    },

    async register() {
      const response = await this.$fire.auth.createUserWithEmailAndPassword(
        this.email,
        this.password
      );

      if (this.photo && this.photo.name) {
        const metadata = {
          contentType: this.photo.type,
        };

        const task = await this.$fire.storage
          .ref(`photos/${Date.now()}-${this.photo.name}`)
          .put(this.photo, metadata);
        this.photoURL = await task.ref.getDownloadURL();
      }

      await response.user.updateProfile({
        displayName: this.username,
        photoURL: this.photoURL,
      });

      this.$router.push({ path: "/" });
    },
  },
};
</script>
