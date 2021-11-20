<template>
  <div id="wrapper">
    <Header />
    <div class="content">
      <form method="post" @submit.prevent="login">
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
        <input id="submit" type="submit" value="Log in" />
      </form>
    </div>
    <Footer />
  </div>
</template>

<script>
import Header from "~/components/Header";
import Footer from "~/components/Footer";

export default {
  components: {
    Header,
    Footer,
  },

  data: function () {
    return {
      email: "",
      password: "",
    };
  },

  methods: {
    login: function (err) {
      this.$store
        .dispatch("signInWithEmail", {
          email: this.email,
          password: this.password,
        })
        .then(() => {
          this.$router.replace({ name: "index" });
        })
        .catch((err) => {
          alert(err.message);
        });
    },
  },
};
</script>
