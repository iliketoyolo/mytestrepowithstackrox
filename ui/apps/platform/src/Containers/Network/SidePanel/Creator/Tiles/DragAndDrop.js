import React, { useCallback } from 'react';
import { connect } from 'react-redux';
import PropTypes from 'prop-types';
import * as Icon from 'react-feather';
import { useDropzone } from 'react-dropzone';

import { fileUploadColors } from 'constants/visuals/colors';
import { actions as notificationActions } from 'reducers/notifications';
import { actions as sidepanelActions } from 'reducers/network/sidepanel';
import wizardStages from 'Containers/Network/SidePanel/wizardStages';

const DragAndDrop = (props) => {
    const showToast = useCallback(() => {
        const errorMessage = 'Invalid file type. Try again.';
        props.addToast(errorMessage);
        setTimeout(props.removeToast, 500);
    }, [props]);

    const onDrop = useCallback(
        (acceptedFiles) => {
            props.setNetworkPolicyModificationState('REQUEST');
            acceptedFiles.forEach((file) => {
                // check file type.
                if (file && !file.name.includes('.yaml')) {
                    showToast();
                    return;
                }

                props.setNetworkPolicyModificationName(file.name);
                const reader = new FileReader();
                reader.onload = () => {
                    const fileAsBinaryString = reader.result;
                    props.setNetworkPolicyModification({ applyYaml: fileAsBinaryString });
                    props.setNetworkPolicyModificationState('SUCCESS');
                };
                reader.onerror = () => {
                    props.setNetworkPolicyModificationState('ERROR');
                };
                reader.readAsBinaryString(file);
                props.setNetworkPolicyModificationSource('UPLOAD');
                props.setWizardStage(wizardStages.simulator);
            });
        },
        [props, showToast]
    );

    const { getRootProps, getInputProps } = useDropzone({ onDrop });

    return (
        <div className="flex flex-col bg-base-100 rounded-sm shadow flex-grow flex-shrink-0 mb-4">
            <div className="flex text-accent-700 p-3 border-b border-base-300 mb-2 items-center flex-shrink-0">
                <Icon.Upload size="20px" strokeWidth="1.5px" />

                <div className="pl-3 font-700 text-lg">Upload a network policy YAML</div>
            </div>
            <div className="mb-3 px-3 font-600 text-lg leading-loose text-base-600">
                Upload your network policies to quickly preview your environment under different
                policy configurations and time windows. When ready, apply the network policies
                directly or share them with your team.
            </div>
            <div
                {...getRootProps()}
                className="bg-warning-100 border border-dashed border-warning-500 cursor-pointer flex flex-col h-full hover:bg-warning-200 justify-center min-h-32 mt-3 outline-none py-3 self-center uppercase w-full"
            >
                <input {...getInputProps()} />
                <div className="flex flex-shrink-0 flex-col">
                    <div
                        className="mt-3 h-18 w-18 self-center rounded-full flex items-center justify-center flex-shrink-0"
                        style={{
                            background: fileUploadColors.BACKGROUND_COLOR,
                            color: fileUploadColors.ICON_COLOR,
                        }}
                    >
                        <Icon.Upload className="h-8 w-8" strokeWidth="1.5px" />
                    </div>
                    <span className="font-700 mt-3 text-center text-warning-800">
                        Upload and simulate network policy YAML
                    </span>
                </div>
            </div>
        </div>
    );
};

DragAndDrop.propTypes = {
    setNetworkPolicyModification: PropTypes.func.isRequired,
    setNetworkPolicyModificationState: PropTypes.func.isRequired,
    setNetworkPolicyModificationSource: PropTypes.func.isRequired,
    setNetworkPolicyModificationName: PropTypes.func.isRequired,
    setWizardStage: PropTypes.func.isRequired,

    addToast: PropTypes.func.isRequired,
    removeToast: PropTypes.func.isRequired,
};

const mapDispatchToProps = {
    setNetworkPolicyModification: sidepanelActions.setNetworkPolicyModification,
    setNetworkPolicyModificationState: sidepanelActions.setNetworkPolicyModificationState,
    setNetworkPolicyModificationSource: sidepanelActions.setNetworkPolicyModificationSource,
    setNetworkPolicyModificationName: sidepanelActions.setNetworkPolicyModificationName,

    setWizardStage: sidepanelActions.setNetworkWizardStage,

    addToast: notificationActions.addNotification,
    removeToast: notificationActions.removeOldestNotification,
};

export default connect(null, mapDispatchToProps)(DragAndDrop);