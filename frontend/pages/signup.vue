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
        <label class="label">Occupation</label>
        <input
          type="text"
          class="input"
          name="occupation"
          v-model="occupation"
        />
      </div>
      <div class="field-wrapper">
        <label class="label">Place of Work</label>
        <input
          type="text"
          class="input"
          name="place_of_work"
          v-model="place_of_work"
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
      <button @click="submit" id="submit" type="submit">Register</button>
    </div>
  </div>
</template>
<script>
export default {
  data: function () {
    return {
      email: null,
      username: null,
      password: null,
      occupation: null,
      place_of_work: null,
      photo: null,
      photoURL: null,
    };
  },
  methods: {
    async fileInput(e) {
      e.preventDefault();
      this.photo = e.target.files[0];
    },

    async getBase64(file) {
      return new Promise((resolve, reject) => {
        const reader = new FileReader();
        reader.readAsDataURL(file);
        reader.onload = () => resolve(reader.result);
        reader.onerror = (error) => reject(error);
      });
    },

    async submit() {
      try {
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

          // upload image to recognition
          this.getBase64(this.photo).then((base64image) => {
            base64image = base64image.replace(
              /^data:image\/[a-z]+;base64,/,
              ""
            );

            this.$recognition
              .uploadBase64Image(this.username, base64image)
              .catch((error) => {
                console.error(error);
              });
          });
        }

        await response.user.updateProfile({
          displayName: this.username,
          photoURL: this.photoURL,
        });

        this.$fire.firestore.collection("users").add({
          uid: response.user.uid,
          is_online: true,
          occupation: this.occupation,
          place_of_work: this.place_of_work,
        });

        this.$router.push({ path: "/" });
      } catch (error) {
        alert(error);
      }
    },
  },
};
</script>
