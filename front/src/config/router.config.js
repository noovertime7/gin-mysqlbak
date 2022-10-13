// eslint-disable-next-line
import { UserLayout, BasicLayout, BlankLayout, PageView } from '@/layouts'
// import { bxAnaalyse } from '@/core/icons'

const RouteView = {
  name: 'RouteView',
  render: h => h('router-view')
}

export const asyncRouterMap = [
  {
    path: '/',
    name: 'index',
    component: BasicLayout,
    meta: { title: 'menu.home' },
    redirect: '/dashboard/analysis',
    children: [
      // dashboard
      {
        path: '/dashboard',
        name: 'dashboard',
        redirect: '/dashboard/analysis',
        component: RouteView,
        meta: { title: 'menu.dashboard', keepAlive: true, icon: 'dashboard', permission: ['dashboard'] },
        children: [
          {
            path: '/dashboard/analysis/:pageNo([1-9]\\d*)?',
            name: 'Analysis',
            component: () => import('@/views/dashboard/Analysis'),
            meta: { title: 'menu.dashboard.analysis', keepAlive: false, permission: ['dashboard'] }
          }
          // 外部链接
          // {
          //   path: 'https://www.baidu.com/',
          //   name: 'Monitor',
          //   meta: { title: 'menu.dashboard.monitor', target: '_blank' }
          // },
          // {
          //   path: '/dashboard/workplace',
          //   name: 'Workplace',
          //   component: () => import('@/views/dashboard/Workplace'),
          //   meta: { title: 'menu.dashboard.workplace', keepAlive: true, permission: ['dashboard'] }
          // }
        ]
      },
      // forms
      // {
      //   path: '/form',
      //   redirect: '/form/base-form',
      //   component: RouteView,
      //   meta: { title: 'menu.form', icon: 'form', permission: ['form'] },
      //   children: [
      //     {
      //       path: '/form/base-form',
      //       name: 'BaseForm',
      //       component: () => import('@/views/form/basicForm'),
      //       meta: { title: 'menu.form.basic-form', keepAlive: true, permission: ['form'] }
      //     },
      //     {
      //       path: '/form/step-form',
      //       name: 'StepForm',
      //       component: () => import('@/views/form/stepForm/StepForm'),
      //       meta: { title: 'menu.form.step-form', keepAlive: true, permission: ['form'] }
      //     },
      //     {
      //       path: '/form/advanced-form',
      //       name: 'AdvanceForm',
      //       component: () => import('@/views/form/advancedForm/AdvancedForm'),
      //       meta: { title: 'menu.form.advanced-form', keepAlive: true, permission: ['form'] }
      //     }
      //   ]
      // },
      // list
      // {
      //   path: '/list',
      //   name: 'list',
      //   component: RouteView,
      //   redirect: '/list/table-list',
      //   meta: { title: 'menu.list', icon: 'table', permission: ['table'] },
      //   children: [
      //     {
      //       path: '/list/table-list/:pageNo([1-9]\\d*)?',
      //       name: 'TableListWrapper',
      //       hideChildrenInMenu: true, // 强制显示 MenuItem 而不是 SubMenu
      //       component: () => import('@/views/list/TableList'),
      //       meta: { title: 'menu.list.table-list', keepAlive: true, permission: ['table'] }
      //     },
      //     {
      //       path: '/list/basic-list',
      //       name: 'BasicList',
      //       component: () => import('@/views/list/BasicList'),
      //       meta: { title: 'menu.list.basic-list', keepAlive: true, permission: ['table'] }
      //     },
      //     {
      //       path: '/list/card',
      //       name: 'CardList',
      //       component: () => import('@/views/list/CardList'),
      //       meta: { title: 'menu.list.card-list', keepAlive: true, permission: ['table'] }
      //     },
      //     {
      //       path: '/list/search',
      //       name: 'SearchList',
      //       component: () => import('@/views/list/search/SearchLayout'),
      //       redirect: '/list/search/article',
      //       meta: { title: 'menu.list.search-list', keepAlive: true, permission: ['table'] },
      //       children: [
      //         {
      //           path: '/list/search/article',
      //           name: 'SearchArticles',
      //           component: () => import('../views/list/search/Article'),
      //           meta: { title: 'menu.list.search-list.articles', permission: ['table'] }
      //         },
      //         {
      //           path: '/list/search/project',
      //           name: 'SearchProjects',
      //           component: () => import('../views/list/search/Projects'),
      //           meta: { title: 'menu.list.search-list.projects', permission: ['table'] }
      //         },
      //         {
      //           path: '/list/search/application',
      //           name: 'SearchApplications',
      //           component: () => import('../views/list/search/Applications'),
      //           meta: { title: 'menu.list.search-list.applications', permission: ['table'] }
      //         }
      //       ]
      //     }
      //   ]
      // },

      // profile
      // {
      //   path: '/profile',
      //   name: 'profile',
      //   component: RouteView,
      //   redirect: '/profile/basic',
      //   meta: { title: 'menu.profile', icon: 'profile', permission: ['profile'] },
      //   children: [
      //     {
      //       path: '/profile/basic',
      //       name: 'ProfileBasic',
      //       component: () => import('@/views/profile/basic'),
      //       meta: { title: 'menu.profile.basic', permission: ['profile'] }
      //     },
      //     {
      //       path: '/profile/advanced',
      //       name: 'ProfileAdvanced',
      //       component: () => import('@/views/profile/advanced/Advanced'),
      //       meta: { title: 'menu.profile.advanced', permission: ['profile'] }
      //     }
      //   ]
      // },

      // result
      // {
      //   path: '/result',
      //   name: 'result',
      //   component: RouteView,
      //   redirect: '/result/success',
      //   meta: { title: 'menu.result', icon: 'check-circle-o', permission: ['result'] },
      //   children: [
      //     {
      //       path: '/result/success',
      //       name: 'ResultSuccess',
      //       component: () => import(/* webpackChunkName: "result" */ '@/views/result/Success'),
      //       meta: { title: 'menu.result.success', keepAlive: false, hiddenHeaderContent: true, permission: ['result'] }
      //     },
      //     {
      //       path: '/result/fail',
      //       name: 'ResultFail',
      //       component: () => import(/* webpackChunkName: "result" */ '@/views/result/Error'),
      //       meta: { title: 'menu.result.fail', keepAlive: false, hiddenHeaderContent: true, permission: ['result'] }
      //     }
      //   ]
      // },

      // // Exception
      {
        path: '/local',
        name: 'local',
        component: RouteView,
        redirect: '/local/app',
        meta: { title: '本地备份', icon: 'project', permission: ['local'] },
        children: [
          {
            path: '/local/app',
            name: '应用管理',
            component: () => import(/* webpackChunkName: "fail" */ '@/views/local/app/app'),
            meta: { title: '应用管理', permission: ['local'] }
          },
          {
            path: '/local/app/task-list/:hostID(\\d+)',
            hidden: true,
            name: '主机任务列表',
            hideChildrenInMenu: true, // 强制显示 MenuItem 而不是 SubMenu
            component: () => import('@/views/local/task/taskList'),
            meta: { title: '主机任务列表', keepAlive: false, hiddenHeaderContent: true, permission: ['local'] }
          },
          {
            path: '/local/task',
            name: '任务列表',
            component: () => import('@/views/local/taskList/taskList-List'),
            meta: { title: '任务列表', permission: ['local'] }
          },
          {
            path: '/local/history',
            name: '备份历史',
            component: () => import(/* webpackChunkName: "fail" */ '@/views/local/history/mysql/MysqlBasicList'),
            meta: { title: '备份历史', permission: ['local'] }
          }
        ]
      },
      // cluster
      {
        path: '/cluster',
        name: 'cluster',
        component: RouteView,
        redirect: '/cluster/service',
        meta: { title: '集群备份', icon: 'project', permission: ['cluster'] },
        children: [
          {
            path: '/cluster/service',
            name: '服务列表',
            component: () => import('@/views/cluster/service/serviceList'),
            meta: { title: '服务列表', permission: ['cluster'] }
          },
          {
            path: '/cluster/app',
            name: '集群应用',
            component: () => import('@/views/cluster/app/appList'),
            meta: { title: '应用管理', permission: ['cluster'] }
          },
          {
            path: '/cluster/task_overview',
            name: '任务总览',
            component: () => import('@/views/cluster/taskoverview/taskOverview'),
            meta: { title: '任务总览', permission: ['cluster'] }
          },
          {
            path: '/cluster/app/task-list/:service_name',
            hidden: true,
            name: '集群任务列表',
            hideChildrenInMenu: true, // 强制显示 MenuItem 而不是 SubMenu
            component: () => import('@/views/cluster/task/model'),
            meta: { title: '集群任务列表', keepAlive: false, hiddenHeaderContent: true, permission: ['cluster'] }
          },
          {
            path: '/cluster/history',
            name: '历史记录',
            component: () => import('@/views/cluster/history/model'),
            meta: { title: '备份历史', permission: ['cluster'] }
          },
          {
            path: '/cluster/history/detail/:id(\\d+)/:service_name',
            name: '详情',
            hidden: true,
            component: () => import('@/views/cluster/history/elastic/detail/esHistoryDetail'),
            meta: { title: '详情', permission: ['cluster'] }
          }
        ]
      },
      // account
      // {
      //   path: '/account',
      //   component: RouteView,
      //   redirect: '/account/center',
      //   name: 'account',
      //   meta: { title: 'menu.account', icon: 'user', keepAlive: true, permission: ['user'] },
      //   children: [
      //     {
      //       path: '/account/center',
      //       name: 'center',
      //       component: () => import('@/views/account/center'),
      //       meta: { title: 'menu.account.center', keepAlive: true, permission: ['user'] }
      //     },
      //     {
      //       path: '/account/settings',
      //       name: 'settings',
      //       component: () => import('@/views/account/settings/Index'),
      //       meta: { title: 'menu.account.settings', hideHeader: true, permission: ['user'] },
      //       redirect: '/account/settings/basic',
      //       hideChildrenInMenu: true,
      //       children: [
      //         {
      //           path: '/account/settings/basic',
      //           name: 'BasicSettings',
      //           component: () => import('@/views/account/settings/BasicSetting'),
      //           meta: { title: 'account.settings.menuMap.basic', hidden: true, permission: ['user'] }
      //         },
      //         {
      //           path: '/account/settings/security',
      //           name: 'SecuritySettings',
      //           component: () => import('@/views/account/settings/Security'),
      //           meta: {
      //             title: 'account.settings.menuMap.security',
      //             hidden: true,
      //             keepAlive: true,
      //             permission: ['user']
      //           }
      //         },
      //         {
      //           path: '/account/settings/custom',
      //           name: 'CustomSettings',
      //           component: () => import('@/views/account/settings/Custom'),
      //           meta: { title: 'account.settings.menuMap.custom', hidden: true, keepAlive: true, permission: ['user'] }
      //         },
      //         {
      //           path: '/account/settings/binding',
      //           name: 'BindingSettings',
      //           component: () => import('@/views/account/settings/Binding'),
      //           meta: { title: 'account.settings.menuMap.binding', hidden: true, keepAlive: true, permission: ['user'] }
      //         },
      //         {
      //           path: '/account/settings/notification',
      //           name: 'NotificationSettings',
      //           component: () => import('@/views/account/settings/Notification'),
      //           meta: {
      //             title: 'account.settings.menuMap.notification',
      //             hidden: true,
      //             keepAlive: true,
      //             permission: ['user']
      //           }
      //         }
      //       ]
      //     }
      //   ]
      // },

      // settings
      {
        path: '/settings',
        name: 'settingsPage',
        component: PageView,
        meta: { title: '系统管理', icon: 'setting', permission: ['settings'] },
        redirect: '/settings/icon-selector',
        children: [
          {
            path: '/settings/icon-selector',
            name: 'TestIconSelect',
            component: () => import('@/views/settings/IconSelectorView'),
            meta: { title: 'IconSelector', icon: 'tool', keepAlive: true, permission: ['settings'] }
          },
          {
            path: '/settings/list/user-list',
            name: 'UserList',
            component: () => import('@/views/settings/TreeList'),
            meta: { title: '用户管理', keepAlive: true }
          }
          // {
          //   path: '/settings/list/permission-list',
          //   name: 'PermissionList',
          //   component: () => import('@/views/settings/PermissionList'),
          //   meta: { title: '权限列表', keepAlive: true }
          // }
        ]
      }
    ]
  },
  {
    path: '*',
    redirect: '/404',
    hidden: true
  }
]

/**
 * 基础路由
 * @type { *[] }
 */
export const constantRouterMap = [
  {
    path: '/user',
    component: UserLayout,
    redirect: '/user/login',
    hidden: true,
    children: [
      {
        path: 'login',
        name: 'login',
        component: () => import(/* webpackChunkName: "user" */ '@/views/user/Login')
      },
      {
        path: 'register',
        name: 'register',
        component: () => import(/* webpackChunkName: "user" */ '@/views/user/Register')
      },
      {
        path: 'register-result',
        name: 'registerResult',
        component: () => import(/* webpackChunkName: "user" */ '@/views/user/RegisterResult')
      },
      {
        path: 'recover',
        name: 'recover',
        component: undefined
      }
    ]
  },

  {
    path: '/404',
    component: () => import(/* webpackChunkName: "fail" */ '@/views/exception/404')
  }
]
