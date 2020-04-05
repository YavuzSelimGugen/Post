<template>
  <div id="xmain">
    <h1>Login</h1>
    <input type="text" name="username" v-model="input.username" placeholder="Username" />
    <input type="password" name="password" v-model="input.password" placeholder="Password" />
    <button type="button" v-on:click="getPosts()">Get Posts</button>
    <div id="app">{{ info }}</div>
  </div>
</template>

<script>
import Axios from "axios";
export default {
  name: "xmain",
  data() {
    return {
      input: {
        username: "",
        password: ""
      },
      info: null
    };
  },
  methods: {
    getPosts() {
      Axios("http://localhost:8082/posts").then(response => {this.info = response});
    },
    login(){

    }
  }
};
</script>

<style scoped>
</style>