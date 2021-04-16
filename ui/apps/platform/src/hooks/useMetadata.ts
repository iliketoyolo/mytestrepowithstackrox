import { useSelector } from 'react-redux';
import { createSelector } from 'reselect';
import { selectors } from 'reducers';

export type Metadata = {
    version: string;
    buildFlavor: string;
    releaseBuild: boolean;
    licenseStatus: string;
    versionString?: string;
};

const selectMetadata = createSelector([selectors.getMetadata], (metadata: Metadata) => metadata);

function useMetadata(): Metadata {
    const metadata: Metadata = useSelector(selectMetadata);

    const versionSuffix = metadata.releaseBuild === false ? ' [DEV BUILD]' : '';
    const versionString = `v${metadata.version}${versionSuffix}`;

    metadata.versionString = versionString;

    return metadata;
}

export default useMetadata;
