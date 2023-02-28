<template>
  <div class="col">
    <div class="card">
      <img v-bind:src="item.image" class="card-img-top" alt="...">
      <div class="card-body">
        <h5 class="card-title text-start"><a v-bind:href="item.code" class="link-card">{{ item.name }}</a></h5>
        <p class="text-start lh-sm">{{ item.description }}</p>
      </div>
      <div class="card-body">
        <div class="d-flex justify-content-between align-items-center">
          <div class="btn-group">
            <button type="button" class="btn btn-sm btn-danger">Button 1</button>
            <button type="button" class="btn btn-sm btn-success">Button 2</button>
          </div>
          <small class="text-muted"></small>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'MainCard',
  props: ['item']
}
</script>

<style scoped>
.link-card {
  text-decoration: none;
  color: black;
  font-weight: 700;

}
</style>
