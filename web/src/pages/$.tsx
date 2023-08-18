import { history } from 'ice';
import { Button, Result } from 'antd';
import { useTranslation } from 'react-i18next';

export default () => {
  const { t } = useTranslation();

  return (
    <Result
      status="404"
      title="404"
      subTitle={t('page_404_title')}
      extra={
        <Button type="primary" onClick={() => history?.go(-1)}>
          {t('back')}
        </Button>
      }
    />
  );
};
