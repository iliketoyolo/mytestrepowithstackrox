import table from '../selectors/table';
import toast from '../selectors/toast';
import tooltip from '../selectors/tooltip';
import navigationSelectors from '../selectors/navigation';

export const url = '/main/integrations';

export const selectors = {
    configure: `${navigationSelectors.navExpandable}:contains("Platform Configuration")`,
    navLink: `${navigationSelectors.navLinks}:contains("Integrations")`,
    kubernetesTile: 'div[role="button"]:contains("Kubernetes")',
    dockerRegistryTile: 'div[role="button"]:contains("Generic Docker Registry")',
    tiles: 'div[role="button"]',
    clairTile: 'div[role="button"]:contains("CoreOS Clair")',
    clairifyTile: 'div[role="button"]:contains("Clairify")',
    slackTile: 'div[role="button"]:contains("Slack")',
    dockerTrustedRegistryTile: 'div[role="button"]:contains("Docker Trusted Registry")',
    quayTile: 'div[role="button"]:contains("Quay.io")',
    amazonECRTile: 'div[role="button"]:contains("Amazon ECR")',
    tenableTile: 'div[role="button"]:contains("Tenable.io")',
    googleContainerRegistryTile: 'div[role="button"]:contains("Google Container Registry")',
    anchoreScannerTile: 'div[role="button"]:contains("Anchore Scanner")',
    ibmCloudTile: 'div[role="button"]:contains("IBM Cloud")',
    microsoftACRTile: 'div[role="button"]:contains("Microsoft ACR")',
    jFrogArtifactoryTile: 'div[role="button"]:contains("JFrog Artifactory")',
    sonatypeNexusTile: 'div[role="button"]:contains("Sonatype Nexus")',
    redHatTile: 'div[role="button"]:contains("Red Hat")',
    googleCloudSCCTile: 'div[role="button"]:contains("Google Cloud SCC")',
    jiraTile: 'div[role="button"]:contains("Jira")',
    emailTile: 'div[role="button"]:contains("Email")',
    splunkTile: 'div[role="button"]:contains("Splunk")',
    pagerDutyTile: 'div[role="button"]:contains("PagerDuty")',
    awsSecurityHubTile: 'div[role="button"]:contains("AWS Security Hub")',
    syslogTile: 'div[role="button"]:contains("Syslog")',
    genericWebhookTile: 'div[role="button"]:contains("Generic Webhook")',
    amazonS3Tile: 'div[role="button"]:contains("Amazon S3")',
    googleCloudStorageTile: 'div[role="button"]:contains("Google Cloud Storage")',
    scopedAccessPluginTile: 'div[role="button"]:contains("Scoped Access Plugin")',
    apiTokenTile: 'div[role="button"]:contains("API Token")',
    clusters: {
        k8sCluster0: 'div.rt-td:contains("Kubernetes Cluster 0")',
    },
    buttons: {
        new: 'button:contains("New")',
        next: 'button:contains("Next")',
        downloadYAML: 'button:contains("Download YAML")',
        delete: 'button:contains("Delete")',
        test: 'button:contains("Test")',
        create: 'button:contains("Create")',
        save: 'button:contains("Save")',
        confirm: 'button:contains("Confirm")',
        generate: 'button:contains("Generate"):not([disabled])',
        revoke: 'button:contains("Revoke")',
        closePanel: 'button[data-testid="cancel"]',
    },
    apiTokenForm: {
        nameInput: 'form[data-testid="api-token-form"] input[name="name"]',
        roleSelect: 'form[data-testid="api-token-form"] .react-select__control',
    },
    apiTokenBox: 'span:contains("eyJ")', // all API tokens start with eyJ
    apiTokenDetailsDiv: 'div[data-testid="api-token-details"]',
    clusterForm: {
        nameInput: 'form[data-testid="cluster-form"] input[name="name"]',
        imageInput: 'form[data-testid="cluster-form"] input[name="mainImage"]',
        endpointInput: 'form[data-testid="cluster-form"] input[name="centralApiEndpoint"]',
    },
    dockerRegistryForm: {
        nameInput: "form input[name='name']",
        typesSelect: 'form .react-select__control',
        endpointInput: "form input[name='docker.endpoint']",
    },
    slackForm: {
        nameInput: "form input[name='name']",
        defaultWebhook: "form input[name='labelDefault']",
        labelAnnotationKey: "form input[name='labelKey']",
    },
    awsSecurityHubForm: {
        nameInput: "form input[name='name']",
        awsAccountNumber: "form input[name='awsSecurityHub.accountId']",
        awsRegion: 'form .react-select__control',
        awsRegionListItems: '.react-select__menu-list > div',
        awsAccessKeyId: "form input[name='awsSecurityHub.credentials.accessKeyId']",
        awsSecretAccessKey: "form input[name='awsSecurityHub.credentials.secretAccessKey']",
    },
    syslogForm: {
        nameInput: "form input[name='name']",
        localFacility: 'form .react-select__control',
        localFacilityListItems: '.react-select__menu-list > div',
        receiverHost: "form input[name='syslog.tcpConfig.hostname']",
        receiverPort: 'form .react-numeric-input input',
        useTls: "form input[name='syslog.tcpConfig.useTls']",
        disableTlsValidation: "form input[name='syslog.tcpConfig.skipTlsVerify']",
    },
    modalHeader: '.ReactModal__Content header',
    formSaveButton: 'button[data-testid="save-integration"]',
    resultsSection: '[data-testid="results-message"]',
    labeledValue: '[data-testid="labeled-value"]',
    plugins: '.mb-6:first div[role="button"]',
    dialog: '.dialog',
    checkboxes: 'input',
    table,
    toast,
    tooltip,
};
