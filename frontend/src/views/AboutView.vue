<template>
  <div id="app">
    <div class="container-fluid">
      <MainNavigation/>
      <MainCards/>
      <MainFooter/>
    </div>
  </div>
</template>

<script>
import MainNavigation from "@/components/MainNavigation.vue";
import MainCards from "@/components/MainCards.vue";
import MainFooter from "@/components/MainFooter.vue";

export default {
  name: 'App',
  components: {
    MainCards,
    MainFooter,
    MainNavigation,
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;

}

.container-fluid {
  margin: 0;
  padding: 0;
}

html {
  margin: 0;
  padding: 0;
}
</style>
