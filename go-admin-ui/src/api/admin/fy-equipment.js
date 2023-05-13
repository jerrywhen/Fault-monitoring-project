import request from '@/utils/request'

// 查询FyEquipment列表
export function listFyEquipment(query) {
    return request({
        url: '/api/v1/fy-equipment',
        method: 'get',
        params: query
    })
}

// 查询FyEquipment详细
export function getFyEquipment (id) {
    return request({
        url: '/api/v1/fy-equipment/' + id,
        method: 'get'
    })
}


// 新增FyEquipment
export function addFyEquipment(data) {
    return request({
        url: '/api/v1/fy-equipment',
        method: 'post',
        data: data
    })
}

// 修改FyEquipment
export function updateFyEquipment(data) {
    return request({
        url: '/api/v1/fy-equipment/'+data.id,
        method: 'put',
        data: data
    })
}

// 删除FyEquipment
export function delFyEquipment(data) {
    return request({
        url: '/api/v1/fy-equipment',
        method: 'delete',
        data: data
    })
}

