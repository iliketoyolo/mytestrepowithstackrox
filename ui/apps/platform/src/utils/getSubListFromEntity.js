import pluralize from 'pluralize';
import entityTypes from 'constants/entityTypes';

function getSubfieldName(entityType) {
    if (entityType === entityTypes.SERVICE_ACCOUNT) {
        return 'serviceAccounts';
    }
    if (entityType === entityTypes.ROLE) {
        return 'k8sRoles';
    }
    return pluralize(entityType.toLowerCase());
}

function getSubListFromEntity(entityData, entityType) {
    if (!entityData || !entityType) {
        return [];
    }
    const field = getSubfieldName(entityType);
    return entityData[field] || [];
}

export default getSubListFromEntity;
