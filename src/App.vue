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
            legend: {
                show: true,
            },
            tooltip: {
                 trigger: 'axis',
                position: function (pt) {
                    return [pt[0], '10%'];
                }
            },
            dataZoom: [
                {
                    type: 'inside',
                    start: 0,
                    end: 100
                },
                // {
                //     start: 0,
                //     end: 20
                // }
            ],
            xAxis: {
                type: 'category',
                boundaryGap: false,
                data:[],
            },
            yAxis: {
                type: 'value',
                min: -10, 
                max: 30,
            },
            series: [
                {
                    name: '2018供暖季',
                    type: 'line',
                    smooth: true,
                    symbol: 'none',
                    data: [],
                },
                {
                    name: '2019供暖季',
                    type: 'line',
                    smooth: true,
                    symbol: 'none',
                    data: [],
                },
                {
                    name: '2020供暖季',
                    type: 'line',
                    smooth: true,
                    symbol: 'none',
                    data: [],
                },
                {
                    name: '2021供暖季',
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
                    option.xAxis[0].data.push(this.formatDate(v.Dt));
                    option.series[0].data.push(averTemp);
                });
                res.data.V2.forEach((v)=>{
                    let averTemp = (v.Temp.Max + v.Temp.Min)/2;
                    // option.xAxis[0].data.push(this.formatDate(v.Dt));
                    option.series[1].data.push(averTemp);
                });
                res.data.V3.forEach((v)=>{
                    let averTemp = (v.Temp.Max + v.Temp.Min)/2;
                    // option.xAxis[0].data.push(this.formatDate(v.Dt));
                    option.series[2].data.push(averTemp);
                });
                res.data.V4.forEach((v)=>{
                    let averTemp = (v.Temp.Max + v.Temp.Min)/2;
                    // option.xAxis[0].data.push(this.formatDate(v.Dt));
                    option.series[3].data.push(averTemp);
                });
                this.mychart.setOption(option);
            });
            
        },
        formatDate:function(dt) {
            let date = new Date(dt*1000);
            return date.getMonth()+1 + "-" + date.getDate();
        },
    }
}
</script>
<style>
.chart {
    height: 500px;
}
</style>
