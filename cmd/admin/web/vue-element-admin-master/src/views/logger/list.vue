<template>
    <div>
        <el-row>
            <el-row class="tool-bar">
                <el-row class="left">
                    <el-input v-model="keyword" placeholder="关键字" style="width:200px;"></el-input>
                    <el-button @click="list" type="primary" icon="el-icon-search" style="margin-left:5px;" plain>搜索</el-button>
                </el-row>
            </el-row>
        </el-row>

        <el-table :data="tableData" stripe style="width:100%">
            <el-table-column prop="id" label="ID" min-width="100" header-align="center" align="center" />
            <el-table-column prop="time" label="时间" min-width="100" header-align="center" align="center" />
            <el-table-column prop="client_ip" label="IP" min-width="100" header-align="center" align="center" />
            <el-table-column prop="status_code" label="状态码" min-width="100" header-align="center" align="center" />
            <el-table-column prop="req_uri" label="URL" min-width="100" header-align="center" align="center" />
        </el-table>

        <el-row>
            <el-pagination
                @size-change="handleSizeChange"
                @current-change="handleCurrentChange"
                :current-page.sync="page"
                :page-sizes="[1, 10, 20, 50, 100]"
                :page-size="size"
                layout="total, sizes, prev, pager, next, jumper"
                :total="total">
            </el-pagination>
        </el-row>

    </div>
</template>

<script>

// import http from '@/utils/http'
import {listLogger} from '@/api/logger'

export default {
    data(){
        return {
            keyword: "",
            page: 1,
            size: 10,
            total: 0,
            tableData: [],
            logger : {}
        }
    },
    created(){
        this.list()
    },
    methods:{
       
        list(){
            let data = {"keyword":this.keyword, "page":this.page,"size":this.size}
            listLogger(data).then(res=>{
                this.tableData = res.data.records;
                this.total = res.data.total;
            })
        },
        handleSizeChange(size){
            this.page = 1;
            this.size = size;
            this.list();
        },
        handleCurrentChange(page){
            this.page = page;
            this.list();
        },
    }
}
</script>