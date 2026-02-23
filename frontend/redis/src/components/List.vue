
<template>
  <div class="listRepl">
    <h1>{{ title }}</h1>
    <form v-on:submit.prevent="listRepl">
    <p>Namespace: <input type="text" placeholder={{ namespace }} required v-model="namespace"></p>
    <button type="submit"> List </button>
    </form>
    <!-- kind": "RedisReplication", "name": "repl2", "status": { "masterNode": "repl2-0", "connectionInfo": 
     { "host": "repl2-master.test2.svc.cluster.local", "port": 6379 } } }, -->
    <div style="text-align: center;" >
      <table>
        <tbody>
        <tr v-for="item in list" :key="item">
          <td>{{ item.kind }}</td>
          <td>{{ item.name}}</td>
        </tr>
      </tbody>
      </table>
    </div>
  <p>{{ note }} <router-link to={{aurthurl}}>{{ authname }}</router-link></p>
 
  </div>
</template>


<script>
import axios from '/Users/mariia.rubina13/Projects/cloud/week5/vue/node_modules/axios'
export default {
  name: 'listRepl',
  props: {
    title: String,
  },
  data() {
    return {
      name: null,
      namespace: "test2",
      status: null,
      code: 0,
      response: null,
      list: null,
      note: null,
      authurl: null,
      authname: null,
    }
  },
  methods: {
    listRepl() {
        this.rq = {"namespace": "test2", "name": this.name}
      axios
      .get("https://rubinity.stackit.gg/api/list/" + this.namespace)
      .then((response) => {
    
       this.code  = response.ok
       this.list = response.data
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
  display: block;
  margin: 0 10px;
}
// .tcontainer{
//   text-align: center;

// }

table {
  margin: 0 auto; /* top/bottom 0, left/right auto */
}

a {
  color: #42b983;
}
</style>
