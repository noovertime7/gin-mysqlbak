module.exports = {
    plugins: ['@vuepress/back-to-top', '@vuepress/search', {
        searchMaxSuggestions: 10
    }],
    title: 'TBak System', // 显示在左上角的网页名称以及首页在浏览器标签显示的title名称
    description: 'TBak数据库备份平台', // meta 中的描述文字，用于SEO
    // 注入到当前页面的 HTML <head> 中的标签
    head: [
        ['link', {rel: 'icon', href: '/logo.png'}],  //浏览器的标签栏的网页图标
    ],
    markdown: {
        lineNumbers: true
    },
    serviceWorker: true,
    themeConfig: {
        logo: '/logo.png',
        lastUpdated: 'lastUpdate', // string | boolean
        nav: [
            {text: '首页', link: '/'},
            {
                text: '分类',
                ariaLabel: '分类',
                items: [
                    {text: '帮助文档', link: '/pages/help/start.md'},
                    {text: '产品介绍', link: '/pages/product/product.md'},
                ]
            },
            {text: '功能演示', link: '/pages/show/show.md'},
            {text: 'Github', link: 'https://github.com/noovertime7/gin-mysqlbak'},
        ],
        sidebar: {
            '/pages/help/': [
                {
                    title: '现在开始',   // 必要的
                    collapsable: false, // 可选的, 默认值是 true,
                    sidebarDepth: 1,    // 可选的, 默认值是 1
                    children: [
                        ['start.md', '开始部署'],
                        ['use.md', '开始使用']
                    ]
                },
                {
                    title: '高级选项',
                    collapsable: false, // 可选的, 默认值是 true,
                    children: [
                        ['advance.md', '高级选项']
                    ]
                }
            ],
        }
    }
}