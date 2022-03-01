<template>
    <div>
        <div ref='myChart' class="chart"></div>
    </div>
</template>
<script>


import axios from 'axios';

let echarts = require('echarts');

export default {
    data(){
        return {
            mychart: null,
            dataV1:[],
            dataV2:[], 
            dataV3:[],
            dataV4:[] 
        };
    },
    mounted: function(){
        // 指定图表的配置项和数据
        let option = {
            title: {
                text: '历年温度曲线'
            },
            tooltip: {},
            xAxis: {
                type: 'time',
                boundaryGap: false,
                
            },
            yAxis: {
                type: 'value',
                min: 0, 
                max: 30,
            },
            series: [
                {
                    name: '2018供暖季',
                    type: 'line',
                    smooth: true,
                    symbol: 'none',
                    data: [],
                }
            ]
        };
        let that = this;
        this.$nextTick(function () {
            that.mychart = echarts.init(that.$refs.myChart);
            that.mychart.setOption(option);
            that.getWeatherData();
        });
    },
    methods: {
        getWeatherData:function() {
            axios.get("/api/get-weather").then(res=>{

                console.log(res.data);
                let option = this.mychart.getOption();
                res.data.V1.forEach((v)=>{
                    // console.log(v.Dt);
                    let averTemp = (v.Temp.Max + v.Temp.Min)/2;
                    option.series[0].data.push([new Date(v.Dt*1000),averTemp]);
                });
                this.mychart.setOption(option);
            });
            
        },
    }
}
</script>
<style>
.chart {
    width: 900px;
    height: 500px;
}
</style>
