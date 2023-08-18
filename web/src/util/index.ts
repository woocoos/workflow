import { TreeAction } from '@knockout-js/api';
import dayjs from 'dayjs';
import timezone from 'dayjs/plugin/timezone';
import utc from 'dayjs/plugin/utc';
import { ReactNode } from 'react';

export type TreeDataState<T> = {
  key: string;
  title: string | ReactNode;
  children?: TreeDataState<T>[];
  parentId: string | number;
  node?: T;
};

dayjs.extend(utc);
dayjs.extend(timezone);

/**
 * 首字母大写
 * @param str
 * @returns
 */
export const firstUpper = (str: string) => {
  return str.slice(0, 1).toUpperCase() + str.slice(1);
};

// zh-CN
export const browserLanguage = () => {
  let locale = '';
  switch (navigator.language) {
    case 'zh':
      locale = 'zh-CN';
      break;
    case 'en':
      locale = 'en-US';
      break;
    default:
      locale = 'zh-CN';
      break;
  }
  return locale;
};

/**
 * tree数据结构的形成
 * @param allList
 * @param parentList
 * @param defineKey
 * @param parentId
 * @returns
 */
export const formatTreeData = <T>(
  allList: Array<T>,
  parentList?: Array<T>,
  defineKey?: {
    key?: string;
    parentId?: string;
    children?: string;
  },
  parentId?: string | number) => {
  const dataKey = { key: 'key', parentId: 'parentId', children: 'children' },
    pid = parentId === undefined ? '0' : parentId;

  if (defineKey) {
    for (let key in defineKey) {
      dataKey[key] = defineKey[key];
    }
  }

  if (!parentList) {
    parentList = allList.filter(item => item[dataKey.parentId] == pid);
  }
  parentList.forEach((pItem) => {
    let children = allList.filter(
      (allItem) => allItem[dataKey.parentId] == pItem[dataKey.key],
    );
    if (children && children.length) {
      pItem[dataKey.children] = formatTreeData(allList, children, dataKey);
    }
  });
  return parentList;
};

/**
 * 循环处理tree结构数据
 * @param data
 * @param key
 * @param callback
 * @returns
 */
export const loopTreeData = <T extends { key: string; children?: Array<T> }>(
  data: Array<T>,
  key: string,
  callback: (rData: T, idx: number, parentData: Array<T>) => void,
): void => {
  for (let idx = 0; idx < data.length; idx++) {
    const item = data[idx];
    if (item.key === key) {
      return callback(item, idx, data);
    }
    if (item.children) {
      loopTreeData(item.children, key, callback);
    }
  }
};


/**
 * 格式化日期
 * @param date
 * @param format  YYYY-MM-DD HH:mm:ss
 * @param tz  时区
 * @param isTzSet  true将当前时间设置成这个时区，false 将当前时间根据时区转换
 * 例子 isTzSet=true
 *      dayjs.tz("2022-07-07 16:30:00", "America/New_York").format("YYYY-MM-DDTHH:mm:ssZ")
 *      = "2022-07-07T16:30:00-04:00"
 * 例子 isTzSet=false
 *      dayjs("2022-07-07T20:30:00Z").tz("America/New_York").format("YYYY-MM-DD HH:mm:ss")
 *      = "2022-07-07 16:30:00"
 * @returns
 */
export const getDate = (
  date: dayjs.ConfigType,
  format = 'YYYY-MM-DD',
  tz?: string,
  isTzSet?: boolean,
) => {
  if (date) {
    if (tz) {
      if (isTzSet) {
        return dayjs.tz(date, tz).format(format);
      } else {
        return dayjs(date).tz(tz).format(format);
      }
    }
    return dayjs(date).format(format);
  } else {
    return null
  }
};

/**
 * 日期数据范围转换成文本表示
 * @param date
 * @param format
 * @returns  '5h6m'
 */
export const dateRangeTurnDuration = (date: [dayjs.ConfigType, dayjs.ConfigType]) => {
  let duration: string[] = [], startDate = dayjs(date[0]);
  ["y", "M", "d", "h", "m", "s"].forEach((unit) => {
    const diffValue = dayjs(date[1]).diff(startDate, unit as dayjs.UnitType);
    if (diffValue) {
      startDate = startDate.add(diffValue, unit as dayjs.ManipulateType)
      duration.push(`${diffValue}${unit}`)
    }
  })
  return duration.join('');
}

/**
 * 日期数据范围转换成文本表示
 * @param startDate
 * @param duration '5h6m'
 * @param format
 * @returns
 */
export const durationTurnEndDate = (
  startDate: dayjs.ConfigType,
  duration: string,
  format = 'YYYY-MM-DD',
) => {
  const keys = ['y', "M", 'd', 'h', 'm', 's'], dl = duration.length;
  let addValue = '', endDate = dayjs(startDate);
  for (let i = 0; i < dl; i++) {
    const item = duration[i];
    if (keys.includes(item)) {
      endDate = endDate.add(Number(addValue), item as dayjs.ManipulateType);
      addValue = '';
    } else {
      addValue += item;
    }
  }
  return getDate(endDate, format);
}

/**
  * 生成随机字符串
  * @param len  几位
  * @returns {string}
  */
export const randomId = (len: number) => {
  let rdmString = '';
  for (; rdmString.length < len; rdmString += Math.random().toString(36).substr(2));
  return rdmString.substr(0, len);
};

/**
 * 根据tree拖拽控件结果处理返回需要的数据
 * @param treeData
 * @param dragInfo
 * @returns
 */
export const getTreeDropData = <T>(treeData: TreeDataState<T>[], dragInfo: any) => {
  const dragTreeData = JSON.parse(JSON.stringify(treeData));

  let dragObj: TreeDataState<T>,
    action: TreeAction = TreeAction.Child,
    sourceId = '',
    targetId = '';
  loopTreeData<TreeDataState<T>>(dragTreeData, dragInfo.dragNode.key, (item, idx, pArr) => {
    pArr.splice(idx, 1);
    sourceId = item.key;
    dragObj = item;
  });
  if (dragInfo.dropToGap) {
    loopTreeData(dragTreeData, dragInfo.node.key, (_item, idx, arr) => {
      if (dragInfo.dropPosition === -1) {
        targetId = arr[0].key;
        action = TreeAction.Up;
        arr.splice(0, 0, dragObj);
      } else {
        targetId = arr[idx].key;
        action = TreeAction.Down;
        arr.splice(dragInfo.dropPosition, 0, dragObj);
      }
    });
  } else {
    // 直接插入第一个子节点
    loopTreeData<TreeDataState<T>>(dragTreeData, dragInfo.node.key, (item) => {
      item.children = item.children || [];
      if (item.children.length) {
        targetId = item.children[0].key;
        action = TreeAction.Up;
      } else {
        targetId = item.key;
        action = TreeAction.Child;
      }
      item.children.unshift(dragObj);
    });
  }

  return {
    sourceId, targetId, action, newTreeData: dragTreeData,
  };
};

/**
 * 更新数据的时候使用  处理增加clear和只提交数值变化的数据
 * @param target   update对象
 * @param original 对比目标的原始数据
 */
export const updateFormat = <T>(target: T, original: Record<string, any>) => {
  const ud: Record<string, any> = {};
  for (const key in target) {
    const tValue = target[key];
    if (tValue !== original[key]) {
      ud[key] = tValue;
      if (typeof tValue != 'boolean' && !tValue) {
        const clearKey = `clear${firstUpper(key)}`.replace('ID', '');
        ud[clearKey] = true;
      }
    }
  }
  return ud as T;
};
