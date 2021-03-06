import React, { useState } from 'react';
import PropTypes from 'prop-types';
import pluralize from 'pluralize';

import { createStructuredSelector } from 'reselect';
import { connect } from 'react-redux';
import { selectors } from 'reducers';

import { getNetworkFlows } from 'utils/networkUtils/getNetworkFlows';
import { filterModes, filterLabels } from 'constants/networkFilterModes';
import { PanelNew, PanelBody, PanelHead, PanelHeadEnd, PanelTitle } from 'Components/Panel';
import TablePagination from 'Components/TablePagination';
import NoResultsMessage from 'Components/NoResultsMessage';
import useSearchFilteredData from 'hooks/useSearchFilteredData';
import NetworkFlowsSearch from './NetworkFlowsSearch';
import NetworkFlowsTable from './NetworkFlowsTable';
import getNetworkFlowValueByCategory from './networkFlowUtils/getNetworkFlowValueByCategory';

const NetworkFlows = ({ edges, filterState, onNavigateToDeploymentById }) => {
    const { networkFlows } = getNetworkFlows(edges, filterState);

    const [page, setPage] = useState(0);
    const [searchOptions, setSearchOptions] = useState([]);

    const filteredNetworkFlows = useSearchFilteredData(
        networkFlows,
        searchOptions,
        getNetworkFlowValueByCategory
    );

    const filterStateString = filterState !== filterModes.all ? filterLabels[filterState] : '';

    if (!filteredNetworkFlows.length) {
        return <NoResultsMessage message={`No ${filterStateString} network flows`} />;
    }

    const subHeaderText = `${filteredNetworkFlows.length} ${filterStateString} ${pluralize(
        'Flow',
        filteredNetworkFlows.length
    )}`;

    return (
        <PanelNew testid="panel">
            <PanelHead>
                <PanelTitle testid="panel-header" text={subHeaderText} />
                <PanelHeadEnd>
                    <div className="flex flex-1">
                        <NetworkFlowsSearch
                            networkFlows={networkFlows}
                            searchOptions={searchOptions}
                            setSearchOptions={setSearchOptions}
                        />
                    </div>
                    <TablePagination
                        page={page}
                        dataLength={filteredNetworkFlows.length}
                        setPage={setPage}
                    />
                </PanelHeadEnd>
            </PanelHead>
            <PanelBody>
                <NetworkFlowsTable
                    networkFlows={filteredNetworkFlows}
                    page={page}
                    filterState={filterState}
                    onNavigateToDeploymentById={onNavigateToDeploymentById}
                />
            </PanelBody>
        </PanelNew>
    );
};

NetworkFlows.propTypes = {
    edges: PropTypes.arrayOf(PropTypes.shape({})),
    filterState: PropTypes.number.isRequired,
    onNavigateToDeploymentById: PropTypes.func.isRequired,
};

NetworkFlows.defaultProps = {
    edges: [],
};

const mapStateToProps = createStructuredSelector({
    filterState: selectors.getNetworkGraphFilterMode,
});

export default connect(mapStateToProps, null)(NetworkFlows);
