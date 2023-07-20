<template>
    <div>
        <ul class="pages">
            <template v-for="index of total">
                <!-- 当前页 -->
                <router-link :to="'/page/' + index">
                    <li v-if="index == nowPage" class="current">{{ index }}</li>
                    <li v-else>{{ index }}</li>
                </router-link>
            </template>

            <!-- <li class="current">1</li>
            <li>2</li>
            <li class="no">...</li>
            <li>100</li> -->
        </ul>
    </div>
</template>
<script>
export default {
    props: {
        // 总页数
        total: {
            type: Number,
            require: true,
            default: 1
        },
        nowPage: {

        }
    },
    data() {
        return {
        }
    }
}
</script>
<style lang="less">
.pages {
    display: flex;
    justify-content: center;

    .current,
    li:hover:not(.no) {
        color: var(--grey-0);
        background-image: linear-gradient(to right, var(--color-pink) 0, var(--color-orange) 100%);
        box-shadow: 0 0 .75rem var(--color-pink-a3);
    }

    li {
        margin: 0 10px;
        padding: 8px 12px;
        border-radius: 5px;
        transition: color .4s;

        &:not(.no) {
            cursor: pointer;
        }
    }

}
</style>