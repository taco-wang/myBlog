<template>
  <section id="main">
    <div class="inner">

      <!-- One -->

      <section id="one" class="wrapper style1" v-for="a in articles">

        <div class="image fit flush">
          <!--<img src="https://api.yingjoy.cn/pic/?t=random&w=1920" alt="">-->
          <!--<img src="{{'https://source.unsplash.com/random/800*600?i=' + a.ID}}" alt="">-->
          <!--获取随机图片-->
          <img v-bind:src="getPic">
        </div>
        <header class="special">
          <h2>{{a.Title}}</h2>
        </header>
        <div class="content">
          {{a.Content}}
        </div>
        <div> >>>>>>>>>>>>>>>>>>>>>>>分割线<<<<<<<<<<<<<<<<<<<<<</div>
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
  methods:{
    getPic:()=>{
      axios.get("https://tools.67cc.cn/img/?type=json").then((res)=>{
        return res.data.url
      })
    },
  },
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
