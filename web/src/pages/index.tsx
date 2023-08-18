import { PageContainer, ProCard, useToken } from '@ant-design/pro-components';
import { useTranslation } from 'react-i18next';


export default () => {
  const { token } = useToken(),
    { t } = useTranslation();

  return (
    <PageContainer
      header={{
        title: t('worktable'),
        style: { background: token.colorBgContainer },
      }}
    >
      hello
    </PageContainer>
  );
};

