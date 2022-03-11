<template>
    <div>
        <div ref='myChart' class="chart"></div>
        <h3 class="chart-title">历年温度曲线</h3>
        <div class="total-container">
            <div>2021供暖季的温度总和为: {{totalV4}}</div>
            <div>前三个年度平均温度的温度总和为:{{averTotal}}</div>
            <div>2018供暖季的温度总和为: {{totalV1}}</div>
            <div>2019供暖季的温度总和为: {{totalV2}}</div>
            <div>2020供暖季的温度总和为: {{totalV3}}</div>
        </div>
    </div>
</template>
<script>


import axios from 'axios';

let echarts = require('echarts');

export default {
    data(){
        return {
            mychart: null,
            totalV1: 0,
            totalV2: 0,
            totalV3: 0,
            totalV4: 0,
            averTotal: 0
        };
    },
    mounted: function(){
        // 指定图表的配置项和数据
        let option = {
            // title: {
            //     text: '历年温度曲线',
            //     // textAlign: 'right',
            //     textVerticalAlign: 'top',
            // },
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
                },
                {
                    name: '前三年平均',
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

                // console.log(res.data);
                this.totalV1= res.data.TotalV1;
                this.totalV2= res.data.TotalV2;
                this.totalV3= res.data.TotalV3;
                this.totalV4= res.data.TotalV4;
                this.averTotal = res.data.AverTotal;

                let option = this.mychart.getOption();
                res.data.V1.forEach((v)=>{
                    // console.log(v.Dt);
                    let averTemp = (v.Temp.Max + v.Temp.Min)/2;
                    // option.xAxis[0].data.push(this.formatDate(v.Dt));
                    option.series[0].data.push(averTemp);
                });
                res.data.V2.forEach((v)=>{
                    let averTemp = (v.Temp.Max + v.Temp.Min)/2;
                    option.xAxis[0].data.push(this.formatDate(v.Dt));
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
                 res.data.Aver.forEach((v)=>{
                    let averTemp = v.Temp.Max;
                    // option.xAxis[0].data.push(this.formatDate(v.Dt));
                    option.series[4].data.push(averTemp);
                });
                this.mychart.setOption(option);
            });
            
        },
        formatDate:function(dt) {
            let date = new Date(dt*1000);
            console.log(dt);
            return date.getMonth()+1 + "-" + date.getDate();
        },
    }
}
</script>
<style>
.chart {
    height: 500px;
}
.chart-title {
    text-align: center;
    line-height: 24px;
}
.total-container {
    text-align: center;
    font-size: 15px;
}
.total-container div {
    float: left;
    width:20%;
    font-weight: bold;
}
</style>
