
import Vue from 'vue';


Vue.component("custom-button2", {
    template: '<button>自定义按钮2</button>',
});

Vue.component("positive-numbers", {
    template: '<div>有{{positiveNumbers.length}}个整数</div>',
    data() {
        return {
            numbers: [1, -2, 3,-5, 9],
        };
    },
    computed: {
        positiveNumbers() {
            return this.numbers.filter(num => num>=0);
        }
    }
});

Vue.component("hello", {
    template: '<div>Hello, {{name}} {{age}}</div>',
    data() {
        return {

        };
    },
    props: {
        name: {
            type: String,
            default: 'World'
        },
        age: {
            type: Number,
            required: true,
            validator(v){
                return Number(v) > 10;
            }
        }
    }
});
Vue.component("display-number", {
    template: '<span>当前的数字是{{number}}</span>',
    props: ['number'],
});