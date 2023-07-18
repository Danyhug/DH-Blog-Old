<template>
    <el-container class="main-container">
        <el-header>
            <Header></Header>
        </el-header>
        <el-container>
            <el-aside width="13.75rem">
                <Nav></Nav>
            </el-aside>
            <el-container>
                <el-main>
                    <Main></Main>
                </el-main>
                <el-footer>
                    <Footer></Footer>
                </el-footer>
            </el-container>
        </el-container>
    </el-container>
</template>
  
<script>
import Header from '@/components/Admin/Header/Header.vue'
import Footer from '@/components/Admin/Footer/Footer.vue'
import Nav from '@/components/Admin/Nav/Nav'
import Main from '@/components/Admin/Main/Main'

export default {
    name: 'AdminView',
    components: { Header, Footer, Nav, Main }
}
</script>

<style lang="less" scoped>
.main-container {
    min-height: 100vh;
}

.el-header,
.el-footer {
    background-color: #409EFF;
    color: #fff;
    text-align: center;
    line-height: 3.75rem;
}

.el-aside {
    background-color: rgb(245, 245, 245);
    color: #333;
    text-align: center;
    line-height: 12.5rem;

    .el-menu {
        color: #303133;
        font-size: 16px;
        background-color: inherit;

        .el-menu-item:hover {
            color: #ffe;
            background-color: #909399;
            border-left: .5rem solid #F56C6C;
        }
    }
}

.el-main {
    padding: 1.25rem 1.875rem;
    background-color: #FFF;
    color: #333;
}



.el-container:nth-child(5) .el-aside,
.el-container:nth-child(6) .el-aside {
    line-height: 16.25rem;
}

.el-container:nth-child(7) .el-aside {
    line-height: 20rem;
}
</style>