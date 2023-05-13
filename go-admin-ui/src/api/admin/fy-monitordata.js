import request from '@/utils/request'

// 查询FyMonitordata列表
export function listFyMonitordata(query) {
    return request({
        url: '/api/v1/fy-monitordata',
        method: 'get',
        params: query
    })
}

// 查询FyMonitordata详细
export function getFyMonitordata (id) {
    return request({
        url: '/api/v1/fy-monitordata/' + id,
        method: 'get'
    })
}


// 新增FyMonitordata
export function addFyMonitordata(data) {
    return request({
        url: '/api/v1/fy-monitordata',
        method: 'post',
        data: data
    })
}

// 修改FyMonitordata
export function updateFyMonitordata(data) {
    return request({
        url: '/api/v1/fy-monitordata/'+data.id,
        method: 'put',
        data: data
    })
}

// 删除FyMonitordata
export function delFyMonitordata(data) {
    return request({
        url: '/api/v1/fy-monitordata',
        method: 'delete',
        data: data
    })
}

