import request from '@/utils/request'

// 查询FyArea列表
export function listFyArea(query) {
    return request({
        url: '/api/v1/fy-area',
        method: 'get',
        params: query
    })
}

// 查询FyArea详细
export function getFyArea (id) {
    return request({
        url: '/api/v1/fy-area/' + id,
        method: 'get'
    })
}


// 新增FyArea
export function addFyArea(data) {
    return request({
        url: '/api/v1/fy-area',
        method: 'post',
        data: data
    })
}

// 修改FyArea
export function updateFyArea(data) {
    return request({
        url: '/api/v1/fy-area/'+data.id,
        method: 'put',
        data: data
    })
}

// 删除FyArea
export function delFyArea(data) {
    return request({
        url: '/api/v1/fy-area',
        method: 'delete',
        data: data
    })
}

