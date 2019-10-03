import React from 'react';
import PropTypes from 'prop-types';

const renderSubHeader = subHeader => {
    if (!subHeader) return null;
    return <div className="mt-1 italic capitalize opacity-75">{subHeader}</div>;
};

const PageHeader = ({ header, subHeader, classes, bgStyle, children }) => (
    <div
        className={`flex h-18 px-4 bg-base-100 w-full flex-no-shrink z-10 border-b border-base-400 ${classes}`}
        style={bgStyle}
        data-test-id="page-header"
    >
        <div className="min-w-max pr-4 self-center">
            <div className="uppercase text-lg tracking-widest font-700 pt-1">{header}</div>
            {renderSubHeader(subHeader)}
        </div>
        <div className="flex w-full items-center">{children}</div>
    </div>
);

PageHeader.propTypes = {
    header: PropTypes.string.isRequired,
    subHeader: PropTypes.string,
    classes: PropTypes.string,
    bgStyle: PropTypes.shape({}),
    children: PropTypes.oneOfType([PropTypes.element, PropTypes.arrayOf(PropTypes.element)])
};

PageHeader.defaultProps = {
    children: null,
    subHeader: null,
    classes: '',
    bgStyle: null
};

export default PageHeader;
