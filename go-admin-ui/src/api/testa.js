import request from '@/utils/request'

// 查询Testa列表
export function TestaList(query) {
    return request({
        url: '/api/v1/TestaList',
        method: 'post',
        data:data
    })
}


