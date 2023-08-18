import i18n from 'i18next';
import { initReactI18next } from 'react-i18next';
import enUS from './locales/en-US';
import zhCN from './locales/zh-CN';
import { LocaleType } from '@knockout-js/layout';

// 多语言文件
const resources = {
  [LocaleType.enUS]: enUS,
  [LocaleType.zhCN]: zhCN,
};

i18n
  .use(initReactI18next)
  .init({
    resources,
    lng: LocaleType.zhCN,
    interpolation: {
      escapeValue: false,
    },
  });

export default i18n;
