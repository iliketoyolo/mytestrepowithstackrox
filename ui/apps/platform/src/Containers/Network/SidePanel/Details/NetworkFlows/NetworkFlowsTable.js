import React from 'react';
import * as Icon from 'react-feather';
import uniqBy from 'lodash/uniqBy';
import uniqWith from 'lodash/uniqWith';

import { filterModes, filterLabels } from 'constants/networkFilterModes';
import { networkProtocolLabels } from 'messages/network';
import Table, {
    Expander,
    rtTrActionsClassName,
    defaultHeaderClassName,
    defaultColumnClassName,
} from 'Components/Table';
import RowActionButton from 'Components/RowActionButton';
import PortsAndProtocolsTable from './PortsAndProtocolsTable';

function renderPortsAndProtocols({ original }) {
    const { portsAndProtocols } = original;
    const uniqProtocols = uniqBy(portsAndProtocols, (datum) => datum.protocol);
    const uniqPorts = uniqBy(portsAndProtocols, (datum) => datum.port);
    if (uniqProtocols.length > 1 || uniqPorts.length > 1) {
        return <PortsAndProtocolsTable portsAndProtocols={portsAndProtocols} />;
    }
    return null;
}

const NetworkFlowsTable = ({ networkFlows, page, filterState, onNavigateToDeploymentById }) => {
    const filterStateString = filterState !== filterModes.all ? filterLabels[filterState] : '';
    const columns = [
        {
            headerClassName: `${defaultHeaderClassName} max-w-10`,
            className: `${defaultColumnClassName} max-w-10 break-all`,
            expander: true,
            Expander: ({ isExpanded, original }) => {
                const uniquePortsAndProtocols = uniqWith(
                    original.portsAndProtocols,
                    (a, b) => a.port === b.port && a.protocol === b.protocol
                );
                if (uniquePortsAndProtocols.length <= 1) {
                    return null;
                }
                return <Expander isExpanded={isExpanded} />;
            },
        },
        {
            headerClassName: `${defaultHeaderClassName} w-2`,
            className: `${defaultColumnClassName} w-2 break-all`,
            Header: 'Traffic',
            accessor: 'traffic',
        },
        {
            headerClassName: `${defaultHeaderClassName} w-10`,
            className: `${defaultColumnClassName} w-10 break-normal`,
            Header: 'Entity',
            accessor: 'entityName',
        },
        {
            headerClassName: `${defaultHeaderClassName} w-3`,
            className: `${defaultColumnClassName} w-3 break-all`,
            Header: 'Type',
            accessor: 'type',
        },
        {
            headerClassName: `${defaultHeaderClassName} w-10`,
            className: `${defaultColumnClassName} w-10 break-all`,
            Header: 'Namespace',
            accessor: 'namespace',
        },
        {
            headerClassName: `${defaultHeaderClassName} w-2`,
            className: `${defaultColumnClassName} w-2 break-all`,
            Header: 'Ports',
            accessor: 'portsAndProtocols',
            Cell: ({ value }) => {
                if (value.length === 0) {
                    return '-';
                }
                const uniquePorts = uniqBy(value, (datum) => datum.port);
                if (uniquePorts.length > 1) {
                    return 'Multiple';
                }
                return uniquePorts[0].port;
            },
        },
        {
            headerClassName: `${defaultHeaderClassName} w-2`,
            className: `${defaultColumnClassName} w-2 break-all`,
            Header: 'Protocols',
            accessor: 'portsAndProtocols',
            Cell: ({ value }) => {
                if (value.length === 0) {
                    return '-';
                }
                const protocols = uniqBy(value, (datum) => datum.protocol)
                    .map((datum) => networkProtocolLabels[datum.protocol])
                    .join(', ');
                return protocols;
            },
        },
        {
            headerClassName: `${defaultHeaderClassName} w-2`,
            className: `${defaultColumnClassName} w-2 break-all`,
            Header: 'Connection',
            accessor: 'connection',
        },
        {
            headerClassName: `${defaultHeaderClassName} hidden`,
            className: `${rtTrActionsClassName} w-4 break-all`,
            accessor: 'deploymentId',
            Cell: ({ original }) => {
                const { deploymentId: id, entityType } = original;
                function onClickHandler() {
                    onNavigateToDeploymentById(id, entityType);
                }
                return (
                    <div className="border-2 border-r-2 border-base-400 bg-base-100 flex">
                        <RowActionButton
                            text="Navigate to Deployment"
                            onClick={onClickHandler}
                            icon={<Icon.ArrowUpRight className="my-1 h-4 w-4" />}
                        />
                    </div>
                );
            },
        },
    ];

    return (
        <Table
            rows={networkFlows}
            columns={columns}
            noDataText={`No ${filterStateString} deployment flows`}
            page={page}
            idAttribute="deploymentId"
            SubComponent={renderPortsAndProtocols}
            noHorizontalPadding
        />
    );
};

export default NetworkFlowsTable;
