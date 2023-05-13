import request from '@/utils/request'

// 查询FyStatus列表
export function listFyStatus(query) {
    return request({
        url: '/api/v1/fy-status',
        method: 'get',
        params: query
    })
}

// 查询FyStatus详细
export function getFyStatus (id) {
    return request({
        url: '/api/v1/fy-status/' + id,
        method: 'get'
    })
}


// 新增FyStatus
export function addFyStatus(data) {
    return request({
        url: '/api/v1/fy-status',
        method: 'post',
        data: data
    })
}

// 修改FyStatus
export function updateFyStatus(data) {
    return request({
        url: '/api/v1/fy-status/'+data.id,
        method: 'put',
        data: data
    })
}

// 删除FyStatus
export function delFyStatus(data) {
    return request({
        url: '/api/v1/fy-status',
        method: 'delete',
        data: data
    })
}

