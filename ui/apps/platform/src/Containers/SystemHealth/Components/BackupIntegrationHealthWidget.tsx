import React, { useState, useEffect, ReactElement } from 'react';

import { fetchBackupIntegrationsHealth } from 'services/IntegrationHealthService';
import { fetchIntegration } from 'services/IntegrationsService';
import integrationsList from 'Containers/Integrations/integrationsList';
import IntegrationHealthWidgetVisual from './IntegrationHealthWidgetVisual';
import { mergeIntegrationResponses, IntegrationMergedItem } from '../utils/integrations';

type WidgetProps = {
    pollingCount: number;
};

const BackupIntegrationHealthWidget = ({ pollingCount }: WidgetProps): ReactElement => {
    const [backupsMerged, setBackupsMerged] = useState([] as IntegrationMergedItem[]);
    const [backupsRequestHasError, setBackupsRequestHasError] = useState(false);

    useEffect(() => {
        Promise.all([fetchBackupIntegrationsHealth(), fetchIntegration('backups')])
            .then(([integrationsHealth, { response }]) => {
                setBackupsMerged(
                    mergeIntegrationResponses(
                        integrationsHealth,
                        response.externalBackups,
                        integrationsList.backups
                    )
                );
                setBackupsRequestHasError(false);
            })
            .catch(() => {
                setBackupsMerged([]);
                setBackupsRequestHasError(true);
            });
    }, [pollingCount]);

    return (
        <IntegrationHealthWidgetVisual
            id="backup-integrations"
            integrationText="Backup Integrations"
            integrationsMerged={backupsMerged}
            requestHasError={backupsRequestHasError}
        />
    );
};

export default BackupIntegrationHealthWidget;