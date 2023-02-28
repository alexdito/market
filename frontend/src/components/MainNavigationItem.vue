<template>
  <li class="nav-item" :class="{'dropdown': item.children_list.length > 0 }">
    <a class="nav-link" :class="{'dropdown-toggle': item.children_list.length > 0 }" v-bind:href="item.code"
       role="button" data-bs-toggle="dropdown" aria-expanded="false">
      {{ item.name }}
    </a>
    <ul v-if="item.children_list.length" class="dropdown-menu">
      <li v-for="children in item.children_list" :key="children.id"><a class="dropdown-item"
                                                                       v-bind:href="item.code + '/' + children.code">{{
          children.name
        }}</a></li>
    </ul>
  </li>
</template>

<script>
export default {
  name: 'MainNavigationItem',
  props: ['item']
}
</script>