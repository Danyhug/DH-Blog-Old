<template>
  <article class="type-img-left">
    <div class="left">
      <img src="https://shilight.cn/images/background/8.jpg" alt="">
    </div>
    <div class="right">
      <div class="top">
        <span class="date">2023-07-11</span>
        <span class="num-word">6.5k 字</span>
        <span class="time-consum">2023-07-11</span>
      </div>
      <p class="title">
        <a href="">安卓大作业——飞鸽传书（高仿微信）</a>
      </p>
      <p class="text"># 产品介绍 # 项目背景 随着网络的出现，即时通讯软件便出现在了人们的生活中，即时通讯app可以帮助人们实现快速沟通，节省时间，提高效率。例如，团队成员可以在办公室或出差途中随</p>
      <div class="bottom">
        <span>
          <a href="">安卓开发</a>
        </span>
        <a href="" class="more">more...</a>
      </div>
    </div>
  </article>
</template>

<script>
export default {
  name: 'Article'
}
</script>

<style lang="less" scoped>
article {
  display: flex;
  width: 100%;
  height: 260px;
  border-radius: 16px;
  background: var(--grey-0);
  overflow: hidden;

  div {
    width: 50%;
  }

  .left {
    position: relative;
    overflow: hidden;
    clip-path: polygon(0 0, 92% 0%, 100% 100%, 0% 100%);

    img {
      display: block;
      width: 100%;
      height: 100%;
      object-fit: cover;
    }
  }

  .right {
    position: relative;
    padding: 20px;

    div {
      width: 100%;
    }

    .top {
      display: flex;
      justify-content: flex-end;
      font-size: 14px;

      span {
        margin-left: 10px;
      }
    }

    .title {
      margin: 15px 0;
      // text-align: center;
      font-weight: 700;

      a {
        font-size: 20px;
        color: rgb(233, 84, 107);
      }
    }

    // 正文描述
    .text {
      font-size: 14px;
      line-height: 2;
    }

    .bottom {
      padding-left: 20px;
      position: absolute;
      left: 0;
      bottom: 0;
      display: flex;
      line-height: 42px;
      justify-content: space-between;

      span {
        a {
          font-size: 14px;
          color: var(--grey-5);
        }
      }

      .more {
        width: 72px;
        height: 42px;
        text-align: center;
        color: rgba(255, 255, 255, .6);
        border-radius: 16px 0;
        background-image: linear-gradient(to right, var(--color-pink) 0, var(--color-orange) 100%);
      }
    }
  }
}
</style>