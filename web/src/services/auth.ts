import store from '@/store';
import { request } from 'ice';
import jwtDcode, { JwtPayload } from 'jwt-decode';

export interface LoginRes {
  accessToken?: string;
  expiresIn?: number;
  refreshToken?: string;
  stateToken?: string;
  callbackUrl?: string;
  user?: {
    id: string;
    displayName: string;
    avatarFileId: string;
    domains: {
      id: string;
      name: string;
    }[];
  };
}

/**
 * 退出登录
 * @returns
 */
export async function logout() {
  return await request.post(`${process.env.ICE_LOGOUT_URL}`);
}

/**
 * 前往登录
 */
export const goLogin = () => {
  const loginUrl = process.env.ICE_LOGIN_URL
  if (loginUrl)  {
    if (loginUrl.toLowerCase().startsWith("http")) {
      const url = new URL(loginUrl)
      if (url.host !== location.host && url.pathname !== location.pathname) {
        location.href = `${loginUrl}?redirect=${encodeURIComponent(location.href)}`
      }
    } else {
      if (loginUrl !== location.pathname) {
        location.href = `${loginUrl}?redirect=${encodeURIComponent(location.href)}`
      }
    }
  }
}

let refreshTokenFn: NodeJS.Timeout;

export function refreshToken() {
  clearTimeout(refreshTokenFn);
  refreshTokenFn = setTimeout(async () => {
    const userState = store.getModelState('user');
    if (userState.token && userState.refreshToken) {
      const jwt = jwtDcode<JwtPayload>(userState.token);
      if ((jwt.exp || 0) * 1000 - Date.now() < 30 * 60 * 1000) {
        // 小于30分钟的时候需要刷新token
        const tr = await request.post(`${process.env.ICE_REFRESH_URL}`, {
          refreshToken: userState.refreshToken,
        });
        if (tr.accessToken) {
          store.dispatch.user.updateToken(tr.accessToken);
        }
      }
    }
  }, 2000);
}
