import request from '@/utils/request'

// 查询FyData列表
export function listFyData(query) {
    return request({
        url: '/api/v1/fy-data',
        method: 'get',
        params: query
    })
}

// 查询FyData详细
export function getFyData (id) {
    return request({
        url: '/api/v1/fy-data/' + id,
        method: 'get'
    })
}


// 新增FyData
export function addFyData(data) {
    return request({
        url: '/api/v1/fy-data',
        method: 'post',
        data: data
    })
}

// 修改FyData
export function updateFyData(data) {
    return request({
        url: '/api/v1/fy-data/'+data.id,
        method: 'put',
        data: data
    })
}

// 删除FyData
export function delFyData(data) {
    return request({
        url: '/api/v1/fy-data',
        method: 'delete',
        data: data
    })
}

