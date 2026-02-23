<template>
  <div class="deleteRepl">
    <!-- <CreateRepl title="Welcome to Redis Replication App"/> -->
     <DeleteRepl title="Delete a Replication"/>
    <!-- <h1>Hello Redis</h1> -->
     <!-- <section>
        <h1>Hello Redis</h1>
     </section> -->
  </div>
</template>

<script>
// @ is an alias to /src
import DeleteRepl from '@/components/Delete.vue' 

export default {
  name: 'DeleteRepl ',
  components: {
    DeleteRepl 
  }
}
</script>