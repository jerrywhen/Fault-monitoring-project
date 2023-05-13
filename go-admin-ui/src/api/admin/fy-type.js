import request from '@/utils/request'

// 查询FyType列表
export function listFyType(query) {
    return request({
        url: '/api/v1/fy-type',
        method: 'get',
        params: query
    })
}

// 查询FyType详细
export function getFyType (id) {
    return request({
        url: '/api/v1/fy-type/' + id,
        method: 'get'
    })
}


// 新增FyType
export function addFyType(data) {
    return request({
        url: '/api/v1/fy-type',
        method: 'post',
        data: data
    })
}

// 修改FyType
export function updateFyType(data) {
    return request({
        url: '/api/v1/fy-type/'+data.id,
        method: 'put',
        data: data
    })
}

// 删除FyType
export function delFyType(data) {
    return request({
        url: '/api/v1/fy-type',
        method: 'delete',
        data: data
    })
}

