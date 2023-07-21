<template>
  <article class="type-img-left">
    <router-link :to="'/article/' + id">
      <div class="cover">
        <div class="left">
          <img src="https://shilight.cn/images/background/8.jpg" alt="">
        </div>
        <div class="right">
          <div class="top">
            <span class="date">
              <Icon iconName="icon-xuanzeshijian" iconSize="1.5"></Icon>
              {{ $getDate(created) }}
            </span>
            <span class="num-word">
              <Icon iconName="icon-gangbi" iconSize="1.5"></Icon>
              {{ $getWordCount(content) }} 字
            </span>
            <span class="time-consum">
              <Icon iconName="icon-yanjingliulankeshi" iconSize="1.3"></Icon>
              {{ (content.length / 200 + 0.5).toFixed(0) }} 分钟
            </span>
          </div>
          <p class="title">
            <a href="">{{ title }}</a>
          </p>
          <p class="text">{{ content }}</p>
          <div class="bottom">
            <span>
              <a href="">安卓开发</a>
            </span>
            <a href="" class="more">more...</a>
          </div>
        </div>
      </div>
    </router-link>
  </article>
</template>

<script>
export default {
  name: 'Article',
  props: {
    id: {
      type: Number,
      require: true
    },
    title: {
      type: String,
      require: true
    },
    content: {
      type: String,
      require: true
    },
    created: {
      type: Number
    }
  },
  methods: {
  },
}
</script>

<style lang="less" scoped>
.cover {
  display: flex;
  width: 100%;
  height: 224px;
  border-radius: 1rem;
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
    padding: 1rem;

    div {
      width: 100%;
    }

    .top {
      display: flex;
      justify-content: flex-end;
      font-size: .875rem;

      span {
        margin-left: .625rem;
      }
    }

    .title {
      margin: 12px 0;
      // text-align: center;
      font-weight: 700;

      a {
        font-size: 1.25rem;
        color: rgb(233, 84, 107);
      }
    }

    // 正文描述
    .text {
      font-size: .875rem;
      line-height: 2;
      max-height: 128px;
      display: -webkit-box;
      -webkit-box-orient: vertical;
      -webkit-line-clamp: 3;
      text-overflow: ellipsis;
      overflow: hidden;
    }

    .bottom {
      padding-left: 1.25rem;
      position: absolute;
      left: 0;
      bottom: 0;
      display: flex;
      line-height: 2.625rem;
      justify-content: space-between;

      span {
        a {
          font-size: .875rem;
          color: var(--grey-5);
        }
      }

      .more {
        width: 4.5rem;
        height: 2.625rem;
        text-align: center;
        color: rgba(255, 255, 255, .6);
        border-radius: 1rem 0;
        background-image: linear-gradient(to right, var(--color-pink) 0, var(--color-orange) 100%);
      }
    }
  }
}
</style>