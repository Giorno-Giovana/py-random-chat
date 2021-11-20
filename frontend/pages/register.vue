<template>
  <div id="wrapper">
    <Header />
    <div class="content">
      <form method="post" @submit.prevent="register">
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
        <input id="submit" type="submit" value="Register" />
      </form>
    </div>
    <Footer />
  </div>
</template>

<script>
import Header from "~/components/Header";
import Footer from "~/components/Footer";
import { FirebaseStorage } from "@/plugins/firebase.js";

export default {
  components: {
    Header,
    Footer,
  },

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
      this.$store
        .dispatch("signUp", {
          email: this.email,
          password: this.password,
        })
        .then((data) => {
          if (this.image && this.image.name) {
            const storageRef = FirebaseStorage.ref(
              `${Date.now()}-${this.image.name}`
            ).put(this.image);
            storageRef.on(
              `state_changed`,
              (snapshot) => {
                this.uploadValue =
                  (snapshot.bytesTransferred / snapshot.totalBytes) * 100;
              },
              (error) => {
                console.log(error.message);
              },
              () => {
                storageRef.snapshot.ref.getDownloadURL().then((url) => {
                  data.user
                    .updateProfile({
                      displayName: this.username,
                      photoURL: url,
                    })
                    .then(() => {
                      this.$router.replace({ name: "index" });
                    });
                });
              }
            );
          }
          data.user
            .updateProfile({
              displayName: this.username,
            })
            .then(() => {
              this.$router.replace({ name: "index" });
            });
        })
        .catch((err) => {
          alert(err.message);
        });
    },
  },
};
</script>

