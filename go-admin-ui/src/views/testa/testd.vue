<template>
  <div>
    <div id="charts"></div>
  </div>
</template>

<script>
import echarts from "echarts";

export default {
  name: "testd",
  components: {},

  mounted() {
    let chartDom = document.getElementById("charts");
    let mychart = echarts.init(chartDom);
    let option = {
      title: {
        text: "ECharts 入门示例",
        textStyle: {
          color: "red",
        },
      },
      tooltip: {},
      legend: {
        data: ["销量"],
      },
      xAxis: {
        data: ["衬衫", "羊毛衫", "雪纺衫", "裤子", "高跟鞋", "袜子"],
      },
      yAxis: {},
      series: [
        {
          name: "销量",
          type: "bar",
          data: [5, 20, 36, 10, 10, 20],
        },
      ],
    };
    mychart.setOption(option);
  },
  data() {
    return {};
  },
  methods: {},
};
</script>
<style>
#charts {
  width: 500px;
  height: 500px;
}
</style>