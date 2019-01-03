<template>
  <section id="main">
    <div class="inner">

      <!-- One -->

      <section id="one" class="wrapper style1" v-for="a in articles">

        <div class="image fit flush">
          <img src="static/pic02.jpg" alt="">
        </div>
        <header class="special">
          <h2>{{a.Title}}</h2>
          <!--<p>vehicula urna sed justo bibendum</p>-->
        </header>
        <div class="content">
          {{a.Content}}
          <!--<p>Curabitur eget semper ante. Morbi eleifend ultricies est, a blandit diam vehicula vel. Vestibulum porttitor nisi quis viverra hendrerit. Suspendisse vel volutpat nibh, vel elementum lacus. Maecenas commodo pulvinar dui, at cursus metus imperdiet vel. Vestibulum et efficitur urna. Duis velit nulla, interdum sed felis sit amet, facilisis auctor nunc. Cras luctus urna at risus feugiat scelerisque nec sed diam. </p>-->
          <!--<p>Nunc fringilla metus odio, at rutrum augue tristique vel. Nulla ac tellus a neque ullamcorper porta a vitae ipsum. Morbi est sapien, hendrerit quis mi in, aliquam bibendum orci. Vestibulum imperdiet nibh vitae maximus posuere. Aenean sit amet nibh feugiat, condimentum tellus eu, malesuada nunc. Mauris ac pulvinar turpis, sit amet pharetra est. Ut bibendum justo condimentum, vehicula ex vel, venenatis libero. Etiam vehicula urna sed justo bibendum, ac lacinia nunc pulvinar. Integer nec velit orci. Vestibulum pellentesque varius dapibus. Morbi ullamcorper augue est, sed luctus orci fermentum dictum. Nunc tincidunt, nisl consequat convallis viverra, metus nisl ultricies dui, vitae dapibus ligula urna sit amet nibh. Duis ut venenatis enim.</p>-->
        </div>
      </section>

      <!-- Two -->


    </div>
  </section>
</template>

<script>
  import config from "../../config"
  import axios from "axios"
export default {
  name: 'HelloWorld',
  data: ()=>{
    return {
      msg: 'Welcome to Your Vue.js App',
      articles:[],
      start: 0,
      end:0,
  }},
  mounted() {
    // console.log("data",data.articles)
    let that = this
    console.log(this)
    axios({url:config.dev.requesturi + '/article/Pages',
    data:JSON.stringify({
      start:0,
      end:10
    }),
      method:"post"
      }
    ).then((res)=>{
      that.articles = res.data.articles
    })
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>
