// @flow
import React from 'react';
import AceEditor from 'react-ace';
import 'brace/mode/javascript';
import 'brace/theme/tomorrow';

import type { SegmentType } from '../../types/Segment.types';
import Field from '../Field/renderField';
import * as styles from './Segment.scss';

type Props = SegmentType & {
    __ID__: number,
    value: ?string,
    onLabelChange: () => void,
    onSegmentChange: () => void,
};

const Segment = (props: Props) => {
    const { isCreator, label, content, createdAt } = props; // TODO: use `createdAt`
    const editorProps = {
        width: '100%',
        height: '300px',
        mode: 'javascript',
        theme: 'tomorrow',
        showPrintMargin: false,
        showGutter: false,
        fontSize: 14,
        wrapEnabled: true,
    };

    const renderEditor = () => {
        const { __ID__, value, onLabelChange, onSegmentChange } = props;
        return (
            <React.Fragment>
                <Field
                    type="text"
                    name="label"
                    placeholder="Label"
                    onChange={onLabelChange}
                />
                <div className={styles.segment}>
                    <AceEditor
                        {...editorProps}
                        value={value}
                        defaultValue="// Write something..."
                        onChange={onSegmentChange}
                        name={`SEGMENT__${__ID__}`}
                    />
                </div>
            </React.Fragment>
        );
    };

    const renderStatic = () => (
        <React.Fragment>
            {label && <h3>{label}</h3>}
            <div className={styles.segment}>
                <AceEditor
                    {...editorProps}
                    className={styles.static}
                    readOnly
                    highlightActiveLine={false}
                    value={content}
                    name={`SEGMENT__${props.id}`}
                />
            </div>
        </React.Fragment>
    );

    return (
        <div className={styles.segmentWrapper}>
            {isCreator ? renderEditor() : renderStatic()}
        </div>
    );
};

export default Segment;
