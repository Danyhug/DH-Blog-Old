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
              <Icon iconName="icon-github" iconSize="1.5"></Icon>
            </li>
            <li>
              <Icon iconName="icon-bilibili" iconSize="1.5" style="fill: rgb(250,116,153)"></Icon>
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
  padding: 40px 25px;
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
        font-size: 1.375rem;
      }
    }

    span {
      display: block;
      border-left: .0625rem solid var(--grey-4);
    }
  }

  img {
    width: 150px;
    height: 150px;
    border-radius: 50%;
  }

  .info-text {
    line-height: 2.2em;
  }

  .links {
    margin: 20px 0;
    display: flex;
    justify-content: space-between;
  }

  .external-links {
    text-align: center;

    li {
      display: inline-block;
      padding: 0 15px;

    }
  }
}

// 右侧部分
.right {
  width: 70%;
  box-shadow: 0 0 10px 2px var(--grey-2);
}
</style>