import request from '@/utils/request'

// 查询FyUser列表
export function listFyUser(query) {
    return request({
        url: '/api/v1/fy-user',
        method: 'get',
        params: query
    })
}

// 查询FyUser详细
export function getFyUser (id) {
    return request({
        url: '/api/v1/fy-user/' + id,
        method: 'get'
    })
}


// 新增FyUser
export function addFyUser(data) {
    return request({
        url: '/api/v1/fy-user',
        method: 'post',
        data: data
    })
}

// 修改FyUser
export function updateFyUser(data) {
    return request({
        url: '/api/v1/fy-user/'+data.id,
        method: 'put',
        data: data
    })
}

// 删除FyUser
export function delFyUser(data) {
    return request({
        url: '/api/v1/fy-user',
        method: 'delete',
        data: data
    })
}

