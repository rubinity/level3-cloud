<template>
  <div class="home">
    <img alt="Redis logo" src="../assets/redis.png">
    <!-- <HelloWorld msg="Welcome to Redis Replication App"/> -->
  </div>
  <!-- <nav class="center">
    <p><router-link to="/create">Create</router-link></p>
    <p><router-link to="/delete">Delete</router-link></p>
    <p><router-link to="/connection">Connection</router-link></p>
    <p><router-link to="/list">List</router-link></p>
    <p><router-link to="/auth">Auth</router-link></p>
  </nav> -->
  <router-view/>
</template>

<script>
// @ is an alias to /src
// import HelloWorld from '@/components/HelloWorld.vue'

export default {
  name: 'HomeView',
  components: {
    // HelloWorld
  }
}
</script>
