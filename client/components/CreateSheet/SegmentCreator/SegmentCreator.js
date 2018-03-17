import React from 'react';

import Segment from '../../Segment/index';

export default class SegmentCreator extends React.PureComponent {
    render() {
        const { value, __ID__, setContent, setLabel } = this.props;

        return (
            <Segment
                value={value && value}
                __ID__={__ID__}
                onSegmentChange={content => this.props.setContent(__ID__, content)}
                onLabelChange={e => setLabel(__ID__, e.target.value)}
                isCreator
            />
        );
    }
}