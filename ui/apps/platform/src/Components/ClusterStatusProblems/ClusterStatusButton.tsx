/* eslint-disable react/jsx-no-bind */
import React, { ReactElement } from 'react';
import { Activity } from 'react-feather';
import { useHistory } from 'react-router-dom';
import { Tooltip, DetailedTooltipOverlay } from '@stackrox/ui-components';

import { clustersBasePath } from 'routePaths';

const fgColorDefault = 'text-base-600';

const fgColorUnhealthy = 'text-alert-700';
const bothColorsUnhealthy = `bg-alert-200 ${fgColorUnhealthy}`;

const fgColorDegraded = 'text-warning-700';

const trClassName = 'align-top leading-normal';
const thClassName = 'font-600 pl-0 pr-2 py-0 text-left';
const tdClassName = 'p-0 text-right';

type ClusterStatusButtonProps = {
    degraded?: number;
    unhealthy?: number;
};

/*
 * Visual indicator in top navigation whether any clusters have health problems.
 *
 * The right corners of the button display non-zero counts.
 * The tooltip body displays query results, including zero counts.
 * A button click opens the Clusters list with a search query.
 */
const ClusterStatusButton = ({
    degraded = 0,
    unhealthy = 0,
}: ClusterStatusButtonProps): ReactElement => {
    const history = useHistory();
    const hasDegradedClusters = degraded > 0;
    const hasUnhealthyClusters = unhealthy > 0;

    // Use table instead of TooltipFieldValue to align numbers.
    // Because of flex-col, tooltip body has full width of tooltip,
    // therefore a div wrapper in needed so that its child table
    // can have automatic width as little as its content needs.
    const resultsElement = (
        <div>
            <table>
                <tbody>
                    <tr
                        className={
                            hasUnhealthyClusters
                                ? `${trClassName} ${bothColorsUnhealthy}`
                                : trClassName
                        }
                        key="unhealthy"
                    >
                        <th className={thClassName} scope="row">
                            Unhealthy
                        </th>
                        <td className={tdClassName}>{unhealthy}</td>
                    </tr>
                    <tr
                        className={
                            hasDegradedClusters ? `${trClassName} ${fgColorDegraded}` : trClassName
                        }
                        key="degraded"
                    >
                        <th className={thClassName} scope="row">
                            Degraded
                        </th>
                        <td className={tdClassName}>{degraded}</td>
                    </tr>
                </tbody>
            </table>
        </div>
    );

    // Unhealthy at upper right to suggest higher severity.
    const unhealthyElement = hasUnhealthyClusters ? (
        <span
            aria-label="Number of clusters with Unhealthy status"
            className={`absolute top-0 right-0 p-1 rounded-bl ${bothColorsUnhealthy}`}
        >
            {unhealthy}
        </span>
    ) : null;

    // Degraded at lower right to suggest lower severity.
    const degradedElement = hasDegradedClusters ? (
        <span
            aria-label="Number of clusters with Degraded status"
            className={`absolute bottom-0 right-0 p-1 ${fgColorDegraded}`}
        >
            {degraded}
        </span>
    ) : null;

    let iconColor = fgColorDefault;
    // The color indicates the more severe health problem.
    if (hasUnhealthyClusters) {
        iconColor = fgColorUnhealthy;
    } else if (hasDegradedClusters) {
        iconColor = fgColorDegraded;
    }

    const onClick = () => {
        history.push({
            pathname: clustersBasePath,
            search: '',
            // TODO after ClustersPage sets search filter according to search query string in URL:
            // If any clusters have problems, then Clusters list has search filter.
            // search: hasUnhealthyClusters || hasDegradedClusters ? '?s[Cluster Health][0]=UNHEALTHY&s[Cluster Health][1]=DEGRADED' : '',
        });
    };

    // Using aria-label for accessibility instead of title to avoid two tooltips.
    // The tooltip has title and subtitle partly to limit its width to the minimum,
    // because the button is near the right edge of the top navigation bar.
    return (
        <Tooltip
            content={
                <DetailedTooltipOverlay
                    title="Cluster Status"
                    subtitle="Problems"
                    body={resultsElement}
                />
            }
        >
            <button
                aria-label="Cluster Status Problems"
                type="button"
                onClick={onClick}
                className="relative flex font-600 h-full items-center pt-1 px-4"
            >
                {unhealthyElement}
                {degradedElement}
                <span className={iconColor}>
                    <Activity className="h-4 w-4" />
                </span>
            </button>
        </Tooltip>
    );
};

export default ClusterStatusButton;
