
import request from '@/utils/request'

export default {
    post : function(vm, api, data){
        const url = api
        return new Promise((resolve, reject)=>{
            request.post(url, data).then(res=>{
                if(res.code !== 20000){
                    vm.$message.error(res.message)
                    reject(res)
                }else{
                    resolve(res)
                }
            })
        })
    }
}
