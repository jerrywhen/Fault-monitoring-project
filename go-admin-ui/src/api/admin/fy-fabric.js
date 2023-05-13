import request from '@/utils/request'

// 查询FyFabric列表
export function listFyFabric(query) {
    return request({
        url: '/api/v1/fy-fabric',
        method: 'get',
        params: query
    })
}

// 查询FyFabric详细
export function getFyFabric (id) {
    return request({
        url: '/api/v1/fy-fabric/' + id,
        method: 'get'
    })
}


// 新增FyFabric
export function addFyFabric(data) {
    return request({
        url: '/api/v1/fy-fabric',
        method: 'post',
        data: data
    })
}

// 修改FyFabric
export function updateFyFabric(data) {
    return request({
        url: '/api/v1/fy-fabric/'+data.id,
        method: 'put',
        data: data
    })
}

// 删除FyFabric
export function delFyFabric(data) {
    return request({
        url: '/api/v1/fy-fabric',
        method: 'delete',
        data: data
    })
}

