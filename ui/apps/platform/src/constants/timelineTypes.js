// these show the types of objects we'll display in the timeline graph
export const graphObjectTypes = {
    EVENT: 'EVENT',
};

// these show the types of root Entities in the timeline view
// @TODO: Use this to keep track of what level the timeline view is in
export const rootTypes = {
    DEPLOYMENT: 'DEPLOYMENT',
    POD: 'POD',
};

// these show the types of the list of Entities we'll show in the left-hand section of the timeline view
export const graphTypes = {
    POD: 'POD',
    CONTAINER: 'CONTAINER',
};

export const eventTypes = {
    POLICY_VIOLATION: 'PolicyViolationEvent',
    PROCESS_ACTIVITY: 'ProcessActivityEvent',
    RESTART: 'ContainerRestartEvent',
    TERMINATION: 'ContainerTerminationEvent',
};

/**
 * @typedef {string} ClusteredEventType
 * @enum {ClusteredEventType}
 */
export const clusteredEventTypes = {
    ...eventTypes,
    PROCESS_IN_BASELINE_ACTIVITY: 'processInBaselineActivityEvent',
    GENERIC: 'Generic',
};

export const selectOptionEventTypes = {
    ALL: 'All',
    ...eventTypes,
};
