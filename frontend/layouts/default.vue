<template>
  <div id="wrapper">
    <!--    <Header />-->
    <Nuxt />
    <!--    <Footer />-->
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
};
</script>

<style>
@import url(//db.onlinewebfonts.com/c/737bc13f79a2e683d7ce615d7f5358c7?family=ALS+Artemius+Sans);
@font-face {
  font-family: "ALS Artemius Sans";
  src: url("//db.onlinewebfonts.com/t/737bc13f79a2e683d7ce615d7f5358c7.eot");
  src: url("//db.onlinewebfonts.com/t/737bc13f79a2e683d7ce615d7f5358c7.eot?#iefix")
      format("embedded-opentype"),
    url("//db.onlinewebfonts.com/t/737bc13f79a2e683d7ce615d7f5358c7.woff2")
      format("woff2"),
    url("//db.onlinewebfonts.com/t/737bc13f79a2e683d7ce615d7f5358c7.woff")
      format("woff"),
    url("//db.onlinewebfonts.com/t/737bc13f79a2e683d7ce615d7f5358c7.ttf")
      format("truetype"),
    url("//db.onlinewebfonts.com/t/737bc13f79a2e683d7ce615d7f5358c7.svg#ALS Artemius Sans")
      format("svg");
}
</style>
