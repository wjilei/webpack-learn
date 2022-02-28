import Vue from 'vue';
import App from './App.vue'
import  './utils.js';

let hours = new Date().getHours();

const CustomButton = {
    template: '<button class="custom-button"><slot>默认按钮</slot></button>'
};

new Vue({
    el: "#app",
    data: {
        isMorning: hours <= 12,
        isAfternoon: hours > 12,
        inputType: "checkbox",
        product1Price: 1234,
        product2Price: 2345,
        product3Price: 32456,
        number: 0,
    },
    components: {
        App,
    },
    filters: {
      formatPrice(value) {
        return (value/100).toFixed(2);
      }  
    },
    created(){
        console.log("hahahah");
    },
    mounted() {
        setInterval(() => {
            this.number = this.number +1; 
        }, 1000);
    }

});