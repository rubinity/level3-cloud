
<template>
  <div class="auth">
    <h1>{{ title }}</h1>
    <form v-on:submit.prevent="auth">
    <p>Namespace: <input type="text" required v-model="namespace"></p>
    <p>Password: <input type="password" required v-model="pass"></p>
    <button type="submit"> Authorize</button>
    </form>
    <!-- <p>{{ code }}</p> -->
<p>{{ cookie }}</p>
  </div>
</template>

<script>
import axios from '/Users/mariia.rubina13/Projects/cloud/week5/vue/node_modules/axios'
export default {
  name: 'AuthUser',
  props: {
    title: String,
  },
  data() {
    return {
      name: null,
      password: null,
      // namespace: String,
      status: null,
      code: 0,
      rq: null,
      response: null,
      cookie: null
    }
  },
  methods: {
    auth() {
        this.rq = {"namespace":"test2","password": this.pass}
      axios
      .post("https://rubinity.stackit.gg/api/auth", this.rq, { withCredentials: true })
      .then((response) => {
       // JSON responses are automatically parsed.
    //    this.response = response.data
       this.code  = response.data.se
       this.cookie = response.set;
       console.log(response);
       console.log("request", this.rq);
      })
      .catch((e) => {
        console.log(e);
        this.status = e
        console.log("request", this.rq);
      });
    }
  },
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
