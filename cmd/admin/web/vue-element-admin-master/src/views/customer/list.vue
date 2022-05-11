<template>
    <div>
        <el-row>
            <el-row class="tool-bar">
                <el-row class="left">
                    <el-input v-model="keyword" placeholder="关键字" style="width:200px;"></el-input>
                    <el-button @click="list" type="primary" icon="el-icon-search" style="margin-left:5px;">搜索</el-button>
                </el-row>
                <el-row class="right">
                    <el-button type="success" v-permission="['editor','admin']" @click="create">添加新客户</el-button>
                </el-row>
            </el-row>
        </el-row>

        <el-table :data="tableData" stripe style="width:100%">
            <el-table-column prop="id" label="ID" min-width="100" header-align="center" align="center" />
            <el-table-column prop="name" label="姓名" min-width="100" header-align="center" align="center" />
            <el-table-column prop="title" label="标题" min-width="100" header-align="center" align="center" />
            <el-table-column prop="url" label="URL" min-width="100" header-align="center" align="center" />
            <el-table-column prop="communication" label="联系方式" min-width="100" header-align="center" align="center" />
            <el-table-column prop="type" label="类型" min-width="100" header-align="center" align="center" />
            <el-table-column  label="操作" min-width="100" header-align="center" align="center" >
                <template slot-scope="scope">
                    <el-button type="warning" icon="el-icon-edit" v-permission="['editor','admin']" circle @click="update(scope.row)"></el-button>
                    <el-button type="danger" icon="el-icon-delete" v-permission="['admin']" circle @click="remove(scope.row.id)"></el-button>
                </template>
            </el-table-column>
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

        <el-dialog
            :title="dialogTitle"
            :visible.sync="dialogVisible"
            width="40%">
            
            <el-form ref="form" :model="Customer" label-width="80px">
                <el-form-item label="名称">
                    <el-input v-model="Customer.name" style="width:300px"></el-input>
                    <span style="margin-left: 10px">用户名不能为空</span>
                </el-form-item>
                <el-form-item label="标题">
                    <el-input v-model="Customer.title" style="width:300px;"></el-input>
                </el-form-item>
                <el-form-item label="URL">
                    <el-input v-model="Customer.url" style="width:300px;"></el-input>
                </el-form-item>
                <el-form-item label="类型">
                    <el-checkbox-group v-model="Customer.type">
                        <el-checkbox label="abc">abc</el-checkbox>
                        <el-checkbox label="123">123</el-checkbox>
                    </el-checkbox-group>
                </el-form-item>
                <el-form-item label="联系方式">
                    <el-input v-model="Customer.communication" type="textarea" rows="2" style="width:300px;"></el-input>
                </el-form-item>
            </el-form>

            <span slot="footer" class="dialog-footer">
                <el-button @click="dialogVisible = false">取 消</el-button>
                <el-button v-if="this.dialogTitle=='新增客户'" type="primary" @click="add">提 交</el-button>
                <el-button v-if="this.dialogTitle=='编辑'" type="primary" @click="edit">保 存</el-button>
            </span>
        </el-dialog>


    </div>
</template>

<script>

// import http from '@/utils/http'
import {listCustomer,createCustomer,updateCustomer,deleteCustomer} from '@/api/customer'

export default {
    data(){
        return {
            keyword: "",
            page: 1,
            size: 10,
            total: 0,
            tableData: [],
            dialogVisible: false,
            dialogTitle:"",
            Customer : {},

        }
    },
    created(){
        this.list()
    },
    methods:{
        // 新增客户 默认选择abc角色
        create(){
            this.dialogVisible = true;
            this.dialogTitle = "新增客户";
            this.Customer = {
                type: ["abc"]
            }
        },
        // 添加提交
        add(){
            // 调用
            createCustomer(this.Customer).then(res=>{
                this.$message({
                    message:'添加成功',
                    type:'success'    
                });
                this.list();
                this.dialogVisible = false;
            })
        },
        // 数据回显
        update(Customer){
            this.dialogVisible = true;
            this.dialogTitle = "编辑";
            this.Customer = Customer;
        },
        // 修改提交
        edit(){
            updateCustomer(this.Customer).then(res=>{
                this.$message({
                    message:'修改成功',
                    type:'success'    
                });
                this.list();
                this.dialogVisible = false;
            })
        },
        // 删除提交
        remove(customerId){
            this.Customer = {
                id :customerId
            }
            
            // 删除确认
            if (window.confirm ("你确定真的要删除吗")){
                deleteCustomer(this.Customer).then(res=>{
                    this.$message({
                        message:'删除成功',
                        type:'success'    
                    });
                    this.list();
                })
            }     
        },
        list(){
            let data = {"keyword":this.keyword, "page":this.page,"size":this.size}
            listCustomer(data).then(res=>{
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