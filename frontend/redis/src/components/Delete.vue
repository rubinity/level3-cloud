
<template>
  <div class="deleteRepl">
    <h1>{{ title }}</h1>
    <form v-on:submit.prevent="rdelete">
    <p>Replication name: <input type="text" placeholder="replication name" required v-model="name"></p>
    <button type="submit"> delete </button>
    </form>
    <p>{{ status }}</p>
    <p>{{ note }} <router-link to={{ aurthurl }}>{{ authname }}</router-link></p>
  </div>
</template>

<script>
import axios from '/Users/mariia.rubina13/Projects/cloud/week5/vue/node_modules/axios'
export default {
  name: 'DeleteRepl',
  props: {
    title: String,
  },
  data() {
    return {
      name: null,
      // namespace: String,
      status: null,
      code: 0,
      rq: null,
      note: null,
      authurl: null,
      authname: null,
    }
  },
  methods: {
    rdelete() {
        this.rq = {"namespace": "test2", "name": this.name}
      axios
      .delete("https://rubinity.stackit.gg/api/delete", {data: this.rq})
      .then((response) => {
       // JSON responses are automatically parsed.
       this.code  = response.ok
       this.status = response.data.status;
       console.log(response);
       console.log("request", this.rq);
      })
      .catch((e) => {
        console.log(e);
        if (e.response.status == "401"){
            this.status = e
            this.note = "Please login:"
            this.authurl = "/auth"
            this.authname = "Auth"
          }
          else{
            this.status = e.response.data.status
          }
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
