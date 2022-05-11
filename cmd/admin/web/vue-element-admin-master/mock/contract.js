module.exports = [
    
    {
      url: '/crm/contract/list',
      type: 'post',
      response: config => {

        return {
            code: 20000,
            data: "SUCCESS",
            data:{
                page:1,
                size:10,
                total:100,
                records:[
                    {id:"xxx1",title:"与xxx签署的合同",state:1,startDate:"2021-12-4T12:00",endDate:"2022-12-4T12:00",desc:"这里是合同简要的描述",customerId:"xxx"},
                    {id:"xxx2",title:"与xxx签署的合同",state:1,startDate:"2021-12-4T12:00",endDate:"2022-12-4T12:00",desc:"这里是合同简要的描述",customerId:"xxx"},
                    {id:"xxx3",title:"与xxx签署的合同",state:1,startDate:"2021-12-4T12:00",endDate:"2022-12-4T12:00",desc:"这里是合同简要的描述",customerId:"xxx"},
                    {id:"xxx4",title:"与xxx签署的合同",state:1,startDate:"2021-12-4T12:00",endDate:"2022-12-4T12:00",desc:"这里是合同简要的描述",customerId:"xxx"},
                    {id:"xxx5",title:"与xxx签署的合同",state:1,startDate:"2021-12-4T12:00",endDate:"2022-12-4T12:00",desc:"这里是合同简要的描述",customerId:"xxx"},
                ]
            }
        }
      }
    },
]