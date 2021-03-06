import React, { useContext } from 'react';
import PropTypes from 'prop-types';
import entityTypes from 'constants/entityTypes';
import { gql, useQuery } from '@apollo/client';
import queryService from 'utils/queryService';

import Loader from 'Components/Loader';
import Widget from 'Components/Widget';
import ViewAllButton from 'Components/ViewAllButton';
import Sunburst from 'Components/visuals/Sunburst';
import NoComponentVulnMessage from 'Components/NoComponentVulnMessage';
import workflowStateContext from 'Containers/workflowStateContext';
import {
    cvssSeverityColorMap,
    cvssSeverityTextColorMap,
    cvssSeverityColorLegend,
} from 'constants/severityColors';
import { getScopeQuery } from 'Containers/VulnMgmt/Entity/VulnMgmtPolicyQueryUtil';

const CVES_QUERY = gql`
    query getCvesByCVSS($query: String, $scopeQuery: String) {
        results: vulnerabilities(query: $query, scopeQuery: $scopeQuery) {
            cve
            cvss
            severity
            summary
        }
    }
`;

const vulnerabilitySeveritySuffix = '_VULNERABILITY_SEVERITY';

const CvesByCvssScore = ({ entityContext, parentContext }) => {
    const { loading, data = {} } = useQuery(CVES_QUERY, {
        variables: {
            query: queryService.entityContextToQueryString(entityContext),
            scopeQuery: getScopeQuery(parentContext),
        },
    });

    let content = <Loader />;
    let header;

    const workflowState = useContext(workflowStateContext);
    const viewAllURL = workflowState
        .pushList(entityTypes.CVE)
        .setSort([{ id: 'cvss', desc: true }])
        .toUrl();

    function getChildren(vulns, severity) {
        return vulns
            .filter(
                (vuln) =>
                    vuln.severity === `${severity.toUpperCase()}${vulnerabilitySeveritySuffix}`
            )
            .map(({ cve, cvss, summary }) => {
                const severityString = `${severity.toUpperCase()}${vulnerabilitySeveritySuffix}`;
                return {
                    severity,
                    name: `${cve} -- ${summary}`,
                    color: cvssSeverityColorMap[severityString],
                    labelColor: cvssSeverityTextColorMap[severityString],
                    textColor: cvssSeverityTextColorMap[severityString],
                    value: cvss,
                    link: workflowState.pushRelatedEntity(entityTypes.CVE, cve).toUrl(),
                };
            });
    }

    function getSunburstData(vulns) {
        return cvssSeverityColorLegend.map(({ title, color, textColor }) => {
            const severity = title.toUpperCase();
            return {
                name: title,
                color,
                children: getChildren(vulns, severity),
                textColor,
                value: 0,
            };
        });
    }

    function getSidePanelData(vulns) {
        return cvssSeverityColorLegend.map(({ title, textColor }) => {
            const severity = `${title.toUpperCase()}${vulnerabilitySeveritySuffix}`;
            const category = vulns.filter((vuln) => vuln.severity === severity);
            const text = `${category.length} rated as ${title}`;
            return {
                text,
                textColor,
            };
        });
    }
    if (!loading) {
        if (!data || !data.results) {
            content = (
                <div className="flex mx-auto items-center">No scanner setup for this registry.</div>
            );
        } else if (!data.results.length) {
            content = <NoComponentVulnMessage />;
        } else {
            const sunburstData = getSunburstData(data.results);
            const sidePanelData = getSidePanelData(data.results).reverse();
            header = <ViewAllButton url={viewAllURL} />;
            content = (
                <Sunburst
                    data={sunburstData}
                    rootData={sidePanelData}
                    totalValue={data.results.length}
                    units="value"
                    small
                />
            );
        }
    }

    return (
        <Widget className="h-full pdf-page" header="CVEs by CVSS Score" headerComponents={header}>
            {content}
        </Widget>
    );
};

CvesByCvssScore.propTypes = {
    entityContext: PropTypes.shape({}),
    parentContext: PropTypes.shape({}),
};

CvesByCvssScore.defaultProps = {
    entityContext: {},
    parentContext: {},
};

export default CvesByCvssScore;
