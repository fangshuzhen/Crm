<template>
    <div>
        <el-row>
            <el-row class="tool-bar">
                <el-row class="left">
                    <el-input v-model="keyword" placeholder="关键字" style="width:200px;"></el-input>
                    <el-button @click="list" type="primary" icon="el-icon-search" style="margin-left:5px;" plain>搜索</el-button>
                </el-row>
                <el-row class="right">
                    <el-button type="success" v-permission="['editor','admin']" @click="create" plain>添加新用户</el-button>
                </el-row>
            </el-row>
        </el-row>

        <el-table :data="tableData" stripe style="width:100%">
            <el-table-column prop="id" label="ID" min-width="100" header-align="center" align="center" />
            <el-table-column prop="name" label="姓名" min-width="100" header-align="center" align="center" />
            <el-table-column prop="username" label="用户名" min-width="100" header-align="center" align="center" />
            <el-table-column prop="avatar" label="头像" min-width="100" header-align="center" align="center" />
            <el-table-column prop="introduction" label="描述" min-width="100" header-align="center" align="center" />
            <el-table-column prop="roles" label="权限" min-width="100" header-align="center" align="center" />
            <el-table-column  label="操作" min-width="150" header-align="center" align="center" >
                <template slot-scope="scope">
                    <el-button type="primary" icon="el-icon-edit" v-permission="['editor','admin']"  @click="update(scope.row)" plain>编辑</el-button>
                    <el-button type="danger" icon="el-icon-delete" v-permission="['admin']"  @click="remove(scope.row.id,scope.row.username)" plain>删除</el-button>
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
            :close-on-click-modal="false"
            :show-close="false"
            width="40%">
            
            <el-form ref="form" :model="userInfo" label-width="80px">
                <el-form-item label="姓名">
                    <el-input v-model="userInfo.name" style="width:300px"></el-input>
                    <span style="margin-left: 10px">用户名不能为空</span>
                </el-form-item>
                <el-form-item label="用户名">
                    <el-input v-model="userInfo.username" style="width:300px;"></el-input>
                </el-form-item>
                <el-form-item label="密码">
                    <el-input v-model="userInfo.password" type="password" style="width:300px;"></el-input>
                </el-form-item>
                <el-form-item label="权限">
                    <el-checkbox-group v-model="userInfo.roles">
                        <el-checkbox label="admin">管理员</el-checkbox>
                        <el-checkbox label="guest">游客</el-checkbox>
                        <el-checkbox label="editor">编辑</el-checkbox>
                    </el-checkbox-group>
                </el-form-item>
                <el-form-item label="用户介绍">
                    <el-input v-model="userInfo.introduction" type="textarea" rows="2" style="width:300px;"></el-input>
                </el-form-item>
            </el-form>

            <span slot="footer" class="dialog-footer">
                <el-button @click="cancel">取 消</el-button>
                <el-button v-if="this.dialogTitle=='新增用户'" type="primary" @click="add">提 交</el-button>
                <el-button v-if="this.dialogTitle=='编辑'" type="primary" @click="edit">保 存</el-button>
            </span>
        </el-dialog>

    </div>
</template>

<script>

// import http from '@/utils/http'
import {listUser,createUser,updateUser,deleteUser} from '@/api/user'
import Cookies from "js-cookie"

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
            userInfo : {},
            checkUser:"false"
        }
    },
    created(){
        this.list()
    },
    methods:{
        cancel(){
            this.list()
            this.dialogVisible = false
        },
        create(){
            this.dialogVisible = true;
            this.dialogTitle = "新增用户";
            this.userInfo = {
                roles: ["guest"]
            }
        },
        // 添加提交
        add(){
            //新增用户
            createUser(this.userInfo).then(res=>{
                this.$message({
                    message:'添加成功',
                    type:'success'    
                });
                this.list();
                this.dialogVisible = false;
            })
        },
        // 数据回显
        update(userInfo){
            this.dialogVisible = true;
            this.dialogTitle = "编辑";
            this.userInfo = userInfo;
            this.userInfo.password = "";
        },
        // 修改提交
        edit(){
            updateUser(this.userInfo).then(res=>{
                this.$message({
                    message:'修改成功',
                    type:'success'    
                });
                this.list();
                this.dialogVisible = false;
            })
        },
        // 删除提交
        remove(userId,userName){
            this.userInfo = {
                id :userId,
                username:userName
            }
            // 删除确认
            if (window.confirm ("你确定真的要删除吗")){
                // 判断将要删除用户是否当前在线
                if(this.userInfo.username == Cookies.get("username")){
                    this.$message({
                        message:'不可删除当前已登录用户: '+ Cookies.get("username") +" !",
                        type:'error'    
                    });
                }    
                else{
                    deleteUser(this.userInfo).then(res=>{
                        this.$message({
                            message:'删除成功',
                            type:'success'    
                        });
                        this.list();
                    })   
                }
            }
        },
        list(){
            let data = {"keyword":this.keyword, "page":this.page,"size":this.size}
            listUser(data).then(res=>{
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