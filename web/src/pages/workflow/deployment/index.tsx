import { ActionType, PageContainer, ProColumns, ProColumnType, ProTable, useToken } from '@ant-design/pro-components';
import { useRef } from 'react';
import { useQuery } from '@knockout-js/ice-urql/runtime';
import { gql } from '@knockout-js/ice-urql/request';
import { Deployment } from '@/generated/workflow/graphql';

const deploymentQuery = gql`query deployment(
  $first: Int,$orderBy:DeploymentOrder,$where:DeploymentWhereInput){
  deployments (first:$first,orderBy:$orderBy,where:$where){
    totalCount,pageInfo{ hasNextPage,hasPreviousPage,startCursor,endCursor }
    edges{
      node{
        id,name,source,tenantID,deployTime,procDefs {
          id,name,version,resourceKey,resourceID
        }
      }
    }
  }
}
`;

export default () => {
  const { token } = useToken(),
    proTableRef = useRef<ActionType>();
  return (
    <PageContainer header={
      {
        title: '部署',
        breadcrumb: {
          items: [
            {
              title: '工作流管理',
            },
            {
              title: '流程部署',
            },
          ],
        },
      }
    }
    >
      <ProTable
        actionRef={proTableRef}
        rowKey={'id'}
        search={{
          searchText: '搜索',
          resetText: '重置',
          labelWidth: 'auto',
        }}
        request={async (params, sort, filter) => {
          const result = useQuery({ query: deploymentQuery });
          console.log(params, sort, filter);
          return {
            data: [],
            success: true,
            total: 0,
          };
        }}
        columns={[
          {
            title: 'ID',
            dataIndex: 'id',
            width: 120,
            search: false,
          },
          {
            title: '部署时间',
            dataIndex: 'code',
            width: 120,
          },
        ] as ProColumns<Deployment>[]}
      />
    </PageContainer>
  );
};
