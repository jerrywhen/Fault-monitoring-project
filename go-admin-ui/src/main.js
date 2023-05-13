// 这段代码是一个Vue.js应用程序的入口文件，主要包括以下内容：

// 引入并使用Vue.js相关的依赖库，如Element UI、Vue-Codemirror、Viser、Vue-DND等。
// 引入并使用自定义的Vue.js插件permission，用于实现权限控制。
// 引入并使用全局方法和组件，挂载到Vue.prototype和Vue.component上，以便在应用程序中全局使用。
// 注册全局过滤器，使其在应用程序中全局可用。
// 关闭生产环境警告提示。
// 创建Vue.js实例，将其渲染到指定的DOM元素上。
// 总的来说，这段代码是一个Vue.js应用程序的配置文件，用于引入和配置应用程序所需的依赖库、插件、组件和过滤器等，并创建Vue.js实例以启动应用程序。

import Vue from 'vue'

import Cookies from 'js-cookie'

import 'normalize.css/normalize.css' // a modern alternative to CSS resets

import Element from 'element-ui'
import './styles/element-variables.scss'

import '@/styles/index.scss' // global css
import '@/styles/admin.scss'

import VueCodemirror from 'vue-codemirror'
import 'codemirror/lib/codemirror.css'
Vue.use(VueCodemirror)

import App from './App'
import store from './store'
import router from './router'
import permission from './directive/permission'

import { getDicts } from '@/api/admin/dict/data'
import { getItems, setItems } from '@/api/table'
import { getConfigKey } from '@/api/admin/sys-config'
import { parseTime, resetForm, addDateRange, selectDictLabel, /* download,*/ selectItemsLabel ,convertHexToRgba} from '@/utils/costum'
//自己写的css
import "@/api/equipment/status-color.css";

import './icons' // icon
import './permission' // permission control
import './utils/error-log' // error log

import Viser from 'viser-vue'
Vue.use(Viser)

import * as filters from './filters' // global filters

import Pagination from '@/components/Pagination'
import BasicLayout from '@/layout/BasicLayout'

import VueParticles from 'vue-particles'
Vue.use(VueParticles)

import '@/utils/dialog'

// 全局方法挂载
Vue.prototype.getDicts = getDicts
Vue.prototype.getItems = getItems
Vue.prototype.setItems = setItems
Vue.prototype.getConfigKey = getConfigKey
Vue.prototype.parseTime = parseTime
Vue.prototype.resetForm = resetForm
Vue.prototype.addDateRange = addDateRange
Vue.prototype.selectDictLabel = selectDictLabel
Vue.prototype.selectItemsLabel = selectItemsLabel
// Vue.prototype.download = download

// 全局组件挂载
Vue.component('Pagination', Pagination)
Vue.component('BasicLayout', BasicLayout)

Vue.prototype.msgSuccess = function (msg) {
  this.$message({ showClose: true, message: msg, type: 'success' })
}

Vue.prototype.msgError = function (msg) {
  this.$message({ showClose: true, message: msg, type: 'error' })
}

Vue.prototype.msgInfo = function (msg) {
  this.$message.info(msg)
}

Vue.use(permission)

Vue.use(Element, {
  size: Cookies.get('size') || 'medium' // set element-ui default size
})

import VueDND from 'awe-dnd'
Vue.use(VueDND)

import 'remixicon/fonts/remixicon.css'

// console.info(`欢迎使用go-admin，谢谢您对我们的支持，在使用过程中如果有什么问题，
// 请访问https://github.com/go-admin-team/go-admin 或者
//  https://github.com/go-admin-team/go-admin-ui 向我们反馈，
//  谢谢！`)

// register global utility filters
Object.keys(filters).forEach(key => {
  Vue.filter(key, filters[key])
})

Vue.config.productionTip = false

new Vue({
  el: '#app',
  router,
  store,
  render: h => h(App)
})
