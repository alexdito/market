<template>
  <div class="nav-item">
    <nav class="navbar navbar-expand-lg bg-body-tertiary">
      <div class="container-fluid">
        <a class="navbar-brand" href="/">Marketplace</a>
        <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
          <span class="navbar-toggler-icon"></span>
        </button>
        <div class="collapse navbar-collapse" id="navbarSupportedContent">
          <ul class="navbar-nav me-auto mb-2 mb-lg-0">
            <MainNavigationItem v-for="item in items" :key="item.id" v-bind:item="item"/>
          </ul>
          <form class="d-flex" role="search">
            <input class="form-control me-2" type="search" placeholder="Поиск" aria-label="Search">
            <button class="btn btn-outline-success" type="submit">Найти</button>
          </form>
        </div>
      </div>
    </nav>
  </div>
</template>

<script>
import axios from "axios";
import MainNavigationItem from "@/components/MainNavigationItem.vue";

export default {
  name: 'MainNavigation',
  components: {MainNavigationItem},
  data() {
    return {
      items: null
    };
  },
  mounted() {
    axios
        .get('http://0.0.0.0:8080/api/v1/navigation/main-structure')
        .then(response => (this.items = response.data.data));
  }
}
</script>

<style scoped>
  .nav-item {
    margin: 0 4px
  }
</style>
