import Vue from 'vue';

let hours = new Date().getHours();
new Vue({
    el: "#app",
    data: {
        isMorning: hours <= 12,
        isAfternoon: hours > 12,
        inputType: "checkbox",
    },
    created(){
        console.log("hahahah");
    }
});