<template>
  <img class="img-renge" alt="Vue logo" src="./assets/miyauchi-renge.jpg">
  <HelloWorld />
</template>

<script>
import HelloWorld from './components/HelloWorld.vue'

export default {
  name: 'App',
  components: {
    HelloWorld
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #e3e6e9;
  height: 100vh;
  background: url("./assets/nyanpasu.jpg") no-repeat center center fixed; 
  background-size: cover;
  background-attachment: fixed;
  /* margin-top: 60px; */
}

.img-renge {
  border-radius: 50%;
  height: 200px;
  width: 400px;
  margin-bottom: 10px;
}
</style>
