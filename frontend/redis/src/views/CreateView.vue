<template>
  <div class="createPepl">
    <!-- <CreateRepl title="Welcome to Redis Replication App"/> -->
     <CreateRepl title="Create a Replication"/>
    <!-- <h1>Hello Redis</h1> -->
  </div>
</template>

<script>
// @ is an alias to /src
import CreateRepl from '@/components/Create.vue' 

export default {
  name: 'CreateRepl ',
  components: {
    CreateRepl 
  }
}
</script>