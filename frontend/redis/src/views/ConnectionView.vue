<template>
  <div class="ConnectInfo">
    <!-- <CreateRepl title="Welcome to Redis Replication App"/> -->
     <ConnectInfo title="Get connection info"/>
    <!-- <h1>Hello Redis</h1> -->
     <!-- <section>
        <h1>Hello Redis</h1>
     </section> -->
  </div>
</template>

<script>
// @ is an alias to /src
import ConnectInfo from '@/components/Connection.vue' 

export default {
  name: 'ConnectInfo ',
  components: {
    ConnectInfo 
  }
}
</script>