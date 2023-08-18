import { defineAppConfig, defineDataLoader } from 'ice';
import { defineAuthConfig } from '@ice/plugin-auth/esm/types';
import { defineStoreConfig } from '@ice/plugin-store/esm/types';
import { defineRequestConfig } from '@ice/plugin-request/esm/types';
import { defineUrqlConfig, requestInterceptor } from "@knockout-js/ice-urql/types";
import store from '@/store';
import '@/assets/styles/index.css';
import { getItem, removeItem, setItem } from '@/pkg/localStore';
import { browserLanguage } from './util';
import { goLogin } from './services/auth';
import jwtDcode, { JwtPayload } from 'jwt-decode';
import { defineChildConfig } from '@ice/plugin-icestark/types';
import { isInIcestark } from '@ice/stark-app';
import { User, userPermissions } from '@knockout-js/api';

const APPCODE = process.env.ICE_APP_CODE || 'workflow',
  ADMIN_URL = process.env.ICE_ADMIN_URL || '',
  UCENTER_URL = process.env.ICE_UCENTER_URL || '',
  LOGIN_URL = process.env.ICE_LOGIN_URL || '',
  REFRESH_URL = process.env.ICE_REFRESH_URL || '',
  // SPM
  SIGN_CID = process.env.ICE_SIGN_CID || 'sign_cid=knockout';

export const icestark = defineChildConfig(() => ({
  mount: () => {
    // 在微应用挂载前执行
  },
  unmount: () => {
    // 在微应用卸载后执行
  },
}));

if (process.env.ICE_CORE_MODE === 'development') {
  // 无登录项目增加前端缓存内容 方便开发和展示
  setItem('token', process.env.ICE_DEV_TOKEN)
  setItem('tenantId', process.env.ICE_DEV_TENANT_ID)
  setItem('user', {
    id: 1,
    displayName: 'admin',
  })
}

// App config, see https://v3.ice.work/docs/guide/basic/app
export default defineAppConfig(() => ({
  // Set your configs here.
  app: {
    rootId: 'app',
  },
}));

// 用来做初始化数据
export const dataLoader = defineDataLoader(async () => {
  if (!isInIcestark()) {
    const sign = SIGN_CID
    if (document.cookie.indexOf(sign) === -1) {
      removeItem('token')
      removeItem('refreshToken')
    }
    document.cookie = sign
  }
  let locale = getItem<string>('locale'),
    token = getItem<string>('token'),
    refreshToken = getItem<string>('refreshToken'),
    darkMode = getItem<string>('darkMode'),
    compactMode = getItem<string>('compactMode'),
    tenantId = getItem<string>('tenantId'),
    user = getItem<User>('user');
  if (token) {
    // 增加jwt判断token过期的处理
    try {
      const jwt = jwtDcode<JwtPayload>(token);
      if ((jwt.exp || 0) * 1000 < Date.now()) {
        token = '';
      }
    } catch (err) {
      token = '';
    }
  }
  if (!locale) {
    locale = browserLanguage();
  }

  return {
    app: {
      locale,
      darkMode,
      compactMode,

    },
    user: {
      token,
      refreshToken,
      tenantId,
      user,
    }
  };
});

// urql
export const urqlConfig = defineUrqlConfig([
  {
    instanceName: 'default',
    url: ADMIN_URL,
    exchangeOpt: {
      authOpts: {
        store: {
          getState: () => {
            const {token, tenantId, refreshToken} = store.getModelState('user')
            return {
              token, tenantId, refreshToken
            }
          },
          setStateToken: (newToken) => {
            store.dispatch.user.updateToken(newToken)
          }
        },
        login: LOGIN_URL,
        refreshApi: REFRESH_URL,
      }
    }
  },
  {
    instanceName: 'ucenter',
    url: UCENTER_URL,
  },
])

// 权限
export const authConfig = defineAuthConfig(async (appData) => {
  const {user, app} = appData,
    initialAuth = {};
  // 判断路由权限
  if (user.token) {
    const result = await userPermissions(APPCODE, {
      Authorization: `Bearer ${user.token}`,
      'X-Tenant-ID': user.tenantId,
    });
    if (result) {
      result.forEach(item => {
        if (item) {
          initialAuth[item.name] = true;
        }
      });
    }
  } else {
    store.dispatch.user.logout();
    goLogin();
  }
  return {
    initialAuth,
  };
});

// store数据项
export const storeConfig = defineStoreConfig(async (appData) => {
  const {user, app} = appData;
  return {
    initialStates: {
      user,
      app,
    },
  };
});

// 请求配置
export const requestConfig = defineRequestConfig(() => {
  return [
    {
      interceptors: requestInterceptor({
        store: {
          getState: () => {
            const {token, tenantId} = store.getModelState('user')
            return {
              token, tenantId
            }
          },
        },
        login: LOGIN_URL,
      })
    },
  ];
});

