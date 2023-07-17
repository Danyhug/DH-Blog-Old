<template>
  <div class="inner">
    <div class="left">
      <div class="info">
        <img src="https://shilight.cn/images/avatar.jpg" alt="">
        <div class="info-text">
          <p class="title">Danyhug</p>
          <p class="sub-title">我们原神玩家是这样的</p>
          <ul class="links">
            <li>
              <a href="">
                <p class="num">46</p>
                文章
              </a>
            </li>
            <span></span>
            <li>
              <a href="">
                <p class="num">17</p>
                分类
              </a>
            </li>
            <span></span>
            <li>
              <a href="">
                <p class="num">53</p>
                标签
              </a>
            </li>
          </ul>

          <ul class="external-links">
            <li>
              <a href="https://github.com/Danyhug" target="_blank">
                <Icon iconName="icon-github" iconSize="1.5"></Icon>
              </a>
            </li>
            <li>
              <a href="https://space.bilibili.com/38630127" target="_blank">
                <Icon iconName="icon-bilibili" iconSize="1.5" style="fill: rgb(250,116,153)"></Icon>
              </a>
            </li>
          </ul>
        </div>
      </div>
    </div>
    <div class="right">
      <slot></slot>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Main',
  components: {
  }
}
</script>

<style lang="less" scoped>
.sub-title {
  color: var(--grey-5);
}

.left,
.right {
  padding: 2.5rem 1.5625rem;
}

.left {
  display: flex;
  flex-direction: column;
  align-items: center;
  width: 30%;
  text-align: center;

  .links {
    li {
      color: var(--grey-6);

      .num {
        font-weight: bold;
        font-size: 22px;
      }
    }

    span {
      display: block;
      border-left: 1px solid var(--grey-4);
    }
  }

  img {
    width: 9.375rem;
    height: 9.375rem;
    border-radius: 50%;
  }

  .info-text {
    line-height: 2.2em;
  }

  .links {
    margin: 1.25rem 0;
    display: flex;
    justify-content: space-between;
  }

  .external-links {
    text-align: center;

    li {
      display: inline-block;
      padding: 0 .9375rem;

    }
  }
}

// 右侧部分
.right {
  width: 70%;
  box-shadow: 0 0 .625rem .125rem var(--grey-2);
}
</style>