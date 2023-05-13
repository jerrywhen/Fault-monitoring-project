import request from '@/utils/request'

// 查询FyHandle列表
export function listFyHandle(query) {
    return request({
        url: '/api/v1/fy-handle',
        method: 'get',
        params: query
    })
}

// 查询FyHandle详细
export function getFyHandle (id) {
    return request({
        url: '/api/v1/fy-handle/' + id,
        method: 'get'
    })
}


// 新增FyHandle
export function addFyHandle(data) {
    return request({
        url: '/api/v1/fy-handle',
        method: 'post',
        data: data
    })
}

// 修改FyHandle
export function updateFyHandle(data) {
    return request({
        url: '/api/v1/fy-handle/'+data.id,
        method: 'put',
        data: data
    })
}

// 删除FyHandle
export function delFyHandle(data) {
    return request({
        url: '/api/v1/fy-handle',
        method: 'delete',
        data: data
    })
}

