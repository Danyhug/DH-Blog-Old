<template>
  <div id="app">
    <Header></Header>
    <Banner>
      <h1>我的个人纪录</h1>
      <h2>DH-BLOG</h2>
    </Banner>
    <Main>
      <HomeChild></HomeChild>
    </Main>
    <Footer></Footer>
  </div>
</template>

<script>
import Header from '@/components/Header/Header'
import Banner from '@/components/Banner/Banner'
import Main from '@/components/Main/Main'
import HomeChild from '@/components/Main/Children/HomeChild'
import Footer from '@/components/Footer/Footer'

export default {
  name: 'HomeView',
  components: {
    Header,
    Banner,
    Main,
    HomeChild,
    Footer
  }
}
</script>
