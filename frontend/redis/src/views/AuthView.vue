<template>
  <div class="auth">
    <!-- <CreateRepl title="Welcome to Redis Replication App"/> -->
     <AuthUser title="Authorization"/>
    <!-- <h1>Hello Redis</h1> -->
     <!-- <section>
        <h1>Hello Redis</h1>
     </section> -->
  </div>
</template>

<script>
// @ is an alias to /src
import AuthUser from '@/components/Auth.vue' 

export default {
  name: 'AuthUser ',
  components: {
    AuthUser 
  }
}
</script>