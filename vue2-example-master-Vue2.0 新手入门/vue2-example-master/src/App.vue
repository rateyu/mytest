<template>
  <div id="app2">
    <h1>第一个工程</h1>
    <firstcomponent></firstcomponent>
    <ul>
        <li><router-link to="/first">部署说明</router-link></li>
        <li><router-link to="/second">升级规则</router-link></li>
      </ul>
    <router-view class="view"></router-view>
  </div>
</template>

<script>
import firstcomponent from './component/firstcomponent.vue'
export default {
  data () {
    return {
      msg: 'Hello Vue!'
    }
  },
  components: { firstcomponent }
}


</script>

<style>
body {
  font-family: Helvetica, sans-serif;
}
</style>
