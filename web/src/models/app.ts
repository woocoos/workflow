import { createModel } from 'ice';
import { setItem } from '@/pkg/localStore';
import { LocaleType } from '@knockout-js/layout';

type ModelState = {
  locale: LocaleType;
  darkMode: boolean;
  compactMode: boolean;
};

export default createModel({
  state: {
    locale: LocaleType.zhCN,
    darkMode: false,
    compactMode: false,
  } as ModelState,
  reducers: {
    updateLocale(prevState: ModelState, payload: LocaleType) {
      setItem('locale', payload);
      prevState.locale = payload;
    },
    updateDarkMode(prevState: ModelState, payload: boolean) {
      setItem('darkMode', payload);
      prevState.darkMode = payload;
    },
  },
  effects: () => ({}),
});
