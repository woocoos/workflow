import Layout from '@/components/layout';
import LayoutStark from '@/components/layout/stark';
import { isInIcestark } from '@ice/stark-app';

export default () => {
  return isInIcestark() ? <LayoutStark /> : <Layout />
}
