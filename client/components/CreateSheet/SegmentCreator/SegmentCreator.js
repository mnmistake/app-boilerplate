// @flow
import React, { PureComponent } from 'react';
import Segment from 'components/Segment';

type Props = {
    value: ?string,
    __ID__: number,
    setContent: (number, string) => void,
    setLabel: (number, string) => void,
};

export default class SegmentCreator extends PureComponent<Props> {
    render() {
        const { value, __ID__, setContent, setLabel } = this.props;

        return (
            <Segment
                value={value && value}
                __ID__={__ID__}
                onSegmentChange={content => setContent(__ID__, content)}
                onLabelChange={e => setLabel(__ID__, e.target.value)}
                isEditor
            />
        );
    }
}
