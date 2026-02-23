<template>
  <div class="about">
    <img alt="View logo" src="../assets/logo.png">
    <VueInfo msg="About Vue"/>
  </div>
</template>

<script>
// @ is an alias to /src
import VueInfo from '@/components/VueInfo.vue'

export default {
  name: 'VueView',
  components: {
    VueInfo
  }
}
</script>
