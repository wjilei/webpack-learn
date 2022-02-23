import Vue from 'vue';

let hours = new Date().getHours();
new Vue({
    el: "#app",
    data: {
        isMorning: hours <= 12,
        isAfternoon: hours > 12,
        inputType: "checkbox",
        product1Price: 1234,
        product2Price: 2345,
        product3Price: 32456,
    },
    filters: {
      formatPrice(value) {
        return (value/100).toFixed(2);
      }  
    },
    created(){
        console.log("hahahah");
    }
});