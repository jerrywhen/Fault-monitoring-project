import request from '@/utils/request'

// 查询FyFault列表
export function listFyFault(query) {
    return request({
        url: '/api/v1/fy-fault',
        method: 'get',
        params: query
    })
}

// 查询FyFault详细
export function getFyFault (id) {
    return request({
        url: '/api/v1/fy-fault/' + id,
        method: 'get'
    })
}


// 新增FyFault
export function addFyFault(data) {
    return request({
        url: '/api/v1/fy-fault',
        method: 'post',
        data: data
    })
}

// 修改FyFault
export function updateFyFault(data) {
    return request({
        url: '/api/v1/fy-fault/'+data.id,
        method: 'put',
        data: data
    })
}

// 删除FyFault
export function delFyFault(data) {
    return request({
        url: '/api/v1/fy-fault',
        method: 'delete',
        data: data
    })
}

