<template>
  <div id="app">
    <login/>
    <img alt="Vue logo" src="./assets/logo.png">
  </div>
</template>

<script>
//import HelloWorld from './components/HelloWorld.vue'
//import xmain from './components/xmain.vue'
import login from './components/login.vue'
export default {
  name: 'App',
  components: {
    login
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
  margin-top: 60px;
}
</style>
