
<template>
  <div class="ConnectInfo">
    <h1>{{ title }}</h1>
    <form v-on:submit.prevent="getConnection">
    <p>Replication name: <input type="text" placeholder="replication name" required v-model="name"></p>
    <button type="submit"> get</button>
    </form>
    <p>host: {{ host}} </p>
    <p>port: {{ port }}</p>
    <!-- <p>{{ status }}</p> -->
    <p>{{ note }} <router-link to={{ aurthurl }}>{{ authname }}</router-link></p>
    
  </div>
</template>

<script>
import axios from '/Users/mariia.rubina13/Projects/cloud/week5/vue/node_modules/axios'
export default {
  name: 'ConnectInfo',
  props: {
    title: String,
  },
  data() {
    return {
      name: null,
      namespace: "test2",
      status: null,
      code: 0,
      host: null,
      note: null,
      port: null,
      res: {},
      authurl: null,
      authname: null,
    }
  },
  methods: {
    getConnection() {
      let url = "https://rubinity.stackit.gg/api/connection/test2/"+this.name ;
      axios
      .get(url)
      .then((response) => {
       // JSON responses are automatically parsed.
       this.code  = response.ok
       this.host = response.data.public_ip
       this.port = response.data.connection.port
       this.status = response.data
       console.log(this.host);
      //  console.log("request", this.rq);
      })
      .catch((e) => {
        this.res = {}
        console.log(e);
        if (e.response.status == "401"){
            this.status = e
            this.note = "Please login:"
            this.authurl = "/auth"
            this.authname = "Auth"
          }
          else{
            this.status = e.response.status
          }

        // console.log("request", this.rq);
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
