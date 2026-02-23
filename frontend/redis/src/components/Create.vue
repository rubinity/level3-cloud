
<template>
  <div class="createRepl">
    <h1>{{ title }}</h1>
    <form v-on:submit.prevent="create">
    <p>Replication name: <input type="text" placeholder="replication name" required v-model="name"></p>
    <p>Replication namespace: <input type="text" placeholder=namespace v-model=namespace readonly></p>
    <p>Cluster size: 1 - 3 <input type="text" placeholder="replication size" required v-model="size"></p>
    <button type="submit" action="create()"> Create </button>
    </form>
     <p>{{ status }}</p>
     <p>{{ note }} <router-link to={{ aurthurl }}>{{ authname }}</router-link></p>
  </div>
</template>

<script>
import axios from '/Users/mariia.rubina13/Projects/cloud/week5/vue/node_modules/axios'
export default {
  name: 'CreateRepl',
  props: {
    title: String,
  },
  data() {
    return {
      size: 3,
      name: null,
      namespace: "test2",
      status: null,
      code: 0,
      req: {
        "size": this.size, "namespace": this.namespace, "name": this.name
      },
      note: null,
      authurl: null,
      authname: null,
    }
  },
  methods: {
    create() {
      axios
      .post("https://rubinity.stackit.gg/api/create",{
        "size": this.size, "namespace": this.namespace, "name": this.name
      })
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
            this.status = e
          }
        this.status = e
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
