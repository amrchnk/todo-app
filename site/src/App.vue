<template>
  <div id="app">
    <Header/>

    <router-view/>
  </div>
</template>

<script>
import Header from '@/components/Header'

export default {
  name: 'App',
  components: {
    Header
  },
  data() {
    return {
      cookie: document.cookie
    }
  },
  beforeUpdate() {
    this.cookie = document.cookie
  }
}
</script>

<style>
#app {
  font-family: 'Roboto', sans-serif;
}

@import url('https://fonts.googleapis.com/css2?family=Roboto:wght@100;300;400;500;700&display=swap');
</style>
