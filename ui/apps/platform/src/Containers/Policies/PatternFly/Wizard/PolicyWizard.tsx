import React, { ReactElement, useState } from 'react';
import { useHistory } from 'react-router-dom';
import { Wizard, Breadcrumb, Title, BreadcrumbItem, Divider } from '@patternfly/react-core';

import { Policy } from 'types/policy.proto';
import { policiesBasePathPatternFly as policiesBasePath } from 'routePaths';
import BreadcrumbItemLink from 'Components/BreadcrumbItemLink';

import { PageAction } from '../policies.utils';

type PolicyWizardProps = {
    pageAction: PageAction;
    policy: Policy;
};

function PolicyWizard({ pageAction, policy }: PolicyWizardProps): ReactElement {
    const history = useHistory();
    const [stepIdReached, setStepIdReached] = useState(1);

    const steps = [
        {
            id: 1,
            name: 'Policy details',
            component: <div>PolicyDetailsForm</div>,
            canJumpTo: stepIdReached >= 1,
        },
        {
            id: 2,
            name: 'Policy behavior',
            component: <div>PolicyBehaviorForm</div>,
            canJumpTo: stepIdReached >= 2,
        },
        {
            id: 3,
            name: 'Policy criteria',
            component: <div>PolicyCriteriaForm</div>,
            canJumpTo: stepIdReached >= 3,
        },
        {
            id: 4,
            name: 'Policy scope',
            component: <div>PolicyScopeForm</div>,
            canJumpTo: stepIdReached >= 4,
        },
        {
            id: 5,
            name: 'Review policy',
            component: <div>ReviewPolicyForm</div>,
            nextButtonText: 'Finish',
            canJumpTo: stepIdReached >= 5,
        },
    ];

    function closeWizard(): void {
        history.goBack();
    }

    function onNext(newStep): void {
        const { id } = newStep;
        setStepIdReached(stepIdReached < id ? id : stepIdReached);
    }

    return (
        <>
            <Breadcrumb className="pf-u-mb-md">
                <BreadcrumbItemLink to={policiesBasePath}>Policies</BreadcrumbItemLink>
                <BreadcrumbItem isActive>{policy?.id || 'Create policy'}</BreadcrumbItem>
            </Breadcrumb>
            <Title headingLevel="h1">{policy?.name || 'Create policy'}</Title>
            <div className="pf-u-mb-md pf-u-mt-sm">
                Design custom security policies for your environment
            </div>
            <Divider component="div" />
            <Wizard
                navAriaLabel={`${pageAction} policy steps`}
                mainAriaLabel={`${pageAction} policy content`}
                onClose={closeWizard}
                steps={steps}
                onNext={onNext}
            />
        </>
    );
}

export default PolicyWizard;