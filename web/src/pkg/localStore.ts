const module = 'workflow';

/**
 * 获取持久化数据
 * @param key
 * @returns
 */
export const getItem = <T>(key: string): T | null => {
  const dataStr = localStorage.getItem(module);
  if (dataStr) {
    const data = JSON.parse(dataStr);
    return data[key];
  }
  return null;
};

/**
 * 数据设置到持久化
 * @param key
 * @param value
 */
export const setItem = <T>(key: string, value: T) => {
  const dataStr = localStorage.getItem(module);
  if (dataStr) {
    const data = JSON.parse(dataStr);
    data[key] = value;
    localStorage.setItem(module, JSON.stringify(data));
  } else {
    const data = {};
    data[key] = value;
    localStorage.setItem(module, JSON.stringify(data));
  }
};

/**
 * 数据持久化移除
 * @param key
 */
export const removeItem = (key: string) => {
  const dataStr = localStorage.getItem(module);
  if (dataStr) {
    const data = JSON.parse(dataStr);
    delete data[key];
    localStorage.setItem(module, JSON.stringify(data));
  }
};

/**
 * 监听持久化数据变化通知更新
 * @param keys
 */
export const monitorKeyChange = (keys: { key: string; onChange: (value: any) => void }[]) => {
  window.addEventListener('storage', (ev) => {
    if (ev.key === module) {
      if (ev.oldValue && ev.newValue) {
        const od = JSON.parse(ev.oldValue),
          nd = JSON.parse(ev.newValue);
        keys.forEach(key => {
          const kk = key.key;
          if (typeof od[kk] === 'object') {
            if (JSON.stringify(od[kk]) !== JSON.stringify(nd[kk])) {
              key.onChange(nd[kk]);
            }
          } else {
            if (od[kk] !== nd[kk]) {
              key.onChange(nd[kk]);
            }
          }
        });
      }
    }
  });
};
