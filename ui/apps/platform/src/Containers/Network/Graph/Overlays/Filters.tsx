import React, { ReactElement } from 'react';
import { connect } from 'react-redux';
import { createStructuredSelector } from 'reselect';

import { selectors } from 'reducers';
import { actions as graphActions } from 'reducers/network/graph';
import { filterModes, filterLabels } from 'constants/networkFilterModes';
import NamespaceEdgeFilter, { NamespaceEdgeFilterState } from './NamespaceEdgeFilter';

const baseButtonClassName =
    'flex-shrink-0 px-2 py-px border-2 rounded-sm uppercase text-xs font-700';
const buttonClassName = `${baseButtonClassName} border-base-400 hover:bg-primary-200 text-base-600`;
const activeButtonClassName = `${baseButtonClassName} bg-primary-300 border-primary-400 hover:bg-primary-200 text-primary-700 border-l-2 border-r-2`;

type FiltersProps = {
    setFilterMode: (mode) => void;
    sidePanelOpen: boolean;
    filterMode: number;
    showNamespaceFlows: NamespaceEdgeFilterState;
    setShowNamespaceFlows: (value) => void;
};

function Filters({
    setFilterMode,
    sidePanelOpen,
    filterMode,
    showNamespaceFlows,
    setShowNamespaceFlows,
}: FiltersProps): ReactElement {
    function handleChange(mode) {
        return () => {
            setFilterMode(mode);
        };
    }

    return (
        <div
            className={`flex absolute top-0 left-0 mt-2 ${
                sidePanelOpen ? 'flex-col' : ''
            } ml-2 absolute z-1`}
        >
            <div className="p-2 bg-primary-100 flex items-center text-sm border-base-400 border-2">
                <span className="text-base-500 font-700 mr-2">Flows:</span>
                <div className="flex items-center">
                    <button
                        type="button"
                        value={filterMode}
                        className={`${
                            filterMode === filterModes.active
                                ? activeButtonClassName
                                : buttonClassName
                        }
                ${filterMode === filterModes.allowed ? 'border-r-0' : ''}`}
                        onClick={handleChange(filterModes.active)}
                        data-testid="network-connections-filter-active"
                    >
                        {`${filterLabels[filterModes.active] as string}`}
                    </button>
                    <button
                        type="button"
                        value={filterMode}
                        className={`${
                            filterMode === filterModes.allowed
                                ? activeButtonClassName
                                : `${buttonClassName} border-l-0 border-r-0`
                        }`}
                        onClick={handleChange(filterModes.allowed)}
                        data-testid="network-connections-filter-allowed"
                    >
                        {`${filterLabels[filterModes.allowed] as string}`}
                    </button>
                    <button
                        type="button"
                        value={filterMode}
                        className={`${
                            filterMode === filterModes.all ? activeButtonClassName : buttonClassName
                        }
                ${filterMode === filterModes.allowed ? 'border-l-0' : ''}`}
                        onClick={handleChange(filterModes.all)}
                        data-testid="network-connections-filter-all"
                    >
                        {`${filterLabels[filterModes.all] as string}`}
                    </button>
                </div>
            </div>
            <div
                className={`px-2 py-1 bg-primary-100 flex items-center text-sm border-base-400 border-2 ${
                    sidePanelOpen ? 'mt-1' : 'ml-1'
                }`}
            >
                <NamespaceEdgeFilter
                    selectedState={showNamespaceFlows}
                    setFilter={setShowNamespaceFlows}
                />
            </div>
        </div>
    );
}

const mapStateToProps = createStructuredSelector({
    sidePanelOpen: selectors.getSidePanelOpen,
    filterMode: selectors.getNetworkGraphFilterMode,
});

const mapDispatchToProps = {
    setFilterMode: graphActions.setNetworkGraphFilterMode,
};

export default connect(mapStateToProps, mapDispatchToProps)(Filters);
