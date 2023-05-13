<template>
  <div>
    <div id="charts"></div>
  </div>
</template>

<script>
import echarts from "echarts";

export default {
  name: "teste",
  components: {},

  mounted() {
    var chartDom = document.getElementById("charts");//初始化
    var myChart = echarts.init(chartDom);
    var option;

    let data = [];
    let now = new Date(1997, 9, 3);
    let oneDay = 24 * 3600 * 1000;
    let value = Math.random() * 1000;//生成小于1000的随机数
    for (var i = 0; i < 1000; i++) {
      data.push(randomData());//将randomdata加入（push）到data数组中
    }

    console.log(data);
    // console.log(now,oneDay,now+oneDay);

    function randomData() {
      now = new Date(+now + oneDay);
      //console.log(now,oneDay,Date(+now+oneDay))
      value = value + Math.random() * 21 - 10;//value+一个-10到10间的数
      //console.log( Math.round(Math.random() * 21 - 10));
      return {
        name: now.toString(),
        value: [
          [now.getFullYear(), now.getMonth() + 1, now.getDate()].join("/"),//join：用"/"连接年月日
          Math.round(value),//四舍五入
        ],
      };
    };

    option = {
      title: {
        text: "Dynamic Data & Time Axis",
      },
      tooltip: {
        trigger: "axis",
        formatter: function (params) {
          params = params[0];
          var date = new Date(params.name);
          return (
            date.getDate() +
            "/" +
            (date.getMonth() + 1) +
            "/" +
            date.getFullYear() +
            " : " +
            params.value[1]
          );
        },

        axisPointer: {
          animation: false,//是否开启动画
        },
      },

      //aplitLine：网格线
      xAxis: {
        type: "time",
        splitLine: {
          show: false,
        },
      },

      yAxis: {
        type: "value",
        boundaryGap: [0, "100%"],
        splitLine: {
          show: false,
        },
      },

      series: [
        {
          name: "Fake Data",
          type: "line",
          showSymbol: false,
          data: data,
        },
      ],
    };

    //https://blog.csdn.net/qq_46003166/article/details/108301157
    //循环执行（setInterval）
    //用法是setInterval（“方法名或方法”，“延时的时间”）， 第一个参数为方法名或者方法，注意为方法名的时候不要加括号,第二个参数为时间间隔
    setInterval(function () {
      for (var i = 0; i < 5; i++) {
        data.shift();//去除数组data第一个值
        data.push(randomData());//加入一个randomdata
      }
      myChart.setOption({
        series: [
          {
            data: data,
          },
        ],
      });
    }, 1000);

    option && myChart.setOption(option);
  },
  data() {
    return {};
  },
  methods: {
    
  },
};
</script>

<style>
#charts {
  width: 500px;
  height: 500px;
}
</style>