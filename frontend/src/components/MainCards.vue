<template>
  <div class="container">
    <div class="row row-cols-1 row-cols-md-3 g-3">
      <MainCard v-for="item in items" :key="item.id" v-bind:item="item"/>
    </div>
  </div>
</template>

<script>
import MainCard from "@/components/MainCard.vue";
import axios from "axios";

export default {
  name: 'MainCards',
  components: {MainCard},
  data() {
    return {
      items: null
    };
  },
  mounted() {
    axios
        .get('http://0.0.0.0:8080/api/v1/card/list')
        .then(response => (this.items = response.data.data));
  }
}
</script>

<style scoped>
.btn-left {
  float: left;
  font-size: 12px;
  width: 100%;
}

.btn-right {
  float: right;
  font-size: 12px;
  width: 100%;
}

.container {
  margin-top: 10px;
}
</style>
