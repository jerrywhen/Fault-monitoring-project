import request from '@/utils/request'

// 查询ArticleA列表
export function listArticleA(query) {
    return request({
        url: '/api/v1/article-a',
        method: 'get',
        params: query
    })
}

// 查询ArticleA详细
export function getArticleA (id) {
    return request({
        url: '/api/v1/article-a/' + id,
        method: 'get'
    })
}


// 新增ArticleA
export function addArticleA(data) {
    return request({
        url: '/api/v1/article-a',
        method: 'post',
        data: data
    })
}

// 修改ArticleA
export function updateArticleA(data) {
    return request({
        url: '/api/v1/article-a/'+data.id,
        method: 'put',
        data: data
    })
}

// 删除ArticleA
export function delArticleA(data) {
    return request({
        url: '/api/v1/article-a',
        method: 'delete',
        data: data
    })
}

