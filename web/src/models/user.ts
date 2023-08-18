import { createModel } from 'ice';
import { LoginRes } from '@/services/auth';
import { setItem, removeItem, getItem } from '@/pkg/localStore';
import { User } from '@knockout-js/api';

type UserState = {
  id: string;
  displayName: string;
  avatarFileId?: string;
};

type ModelState = {
  refreshToken: string;
  token: string;
  tenantId: string;
  user: UserState | null;
};


export default createModel({
  state: {
    token: '',
    refreshToken: '',
    tenantId: '',
    user: null,
    darkMode: false,
    compactMode: false,
  } as ModelState,
  reducers: {
    updateToken(prevState: ModelState, payload: string) {
      if (payload) {
        setItem('token', payload);
      } else {
        removeItem('token');
      }
      prevState.token = payload;
    },
    updateRefreshToken(prevState: ModelState, payload: string) {
      if (payload) {
        setItem('refreshToken', payload);
      } else {
        removeItem('refreshToken');
      }
      prevState.refreshToken = payload;
    },
    updateTenantId(prevState: ModelState, payload: string) {
      if (payload) {
        setItem('tenantId', payload);
      } else {
        removeItem('tenantId');
      }
      prevState.tenantId = payload;
    },
    updateUser(prevState: ModelState, payload: UserState | null) {
      if (payload) {
        setItem('user', payload);
      } else {
        removeItem('user');
      }
      prevState.user = payload;
    },
  },
  effects: () => ({
    /**
     * 登录
     * @param payload
     */
    async loginAfter(payload: LoginRes) {
      if (payload.accessToken) {
        this.updateToken(payload.accessToken);
        if (payload.user) {
          this.saveUser({
            id: payload.user.id,
            displayName: payload.user.displayName,
            avatarFileID: payload.user?.avatarFileId || '',
          } as User)
          if (payload.user.domains?.length) {
            const tenantId = getItem<string>('tenantId')
            if (!payload.user.domains.find(item => item.id == tenantId)) {
              this.updateTenantId(payload.user.domains[0].id);
            }
          } else {
            this.updateTenantId('');
          }
        }
        this.updateRefreshToken(payload.refreshToken || '');
      }
    },
    /**
     * 退出
     */
    async logout() {
      this.updateToken('');
      this.updateUser(null);
    },
    /**
     * 更新用户信息
     * @param user
     */
    async saveUser(user: User) {
      this.updateUser({
        id: user.id,
        displayName: user.displayName,
        avatarFileId: user.avatarFileID || undefined,
      });
    },
    /**
     * 更新租户id
     * @param tenantId
     */
    async saveTenantId(tenantId: string) {
      this.updateTenantId(tenantId);
    },
  }),
});
