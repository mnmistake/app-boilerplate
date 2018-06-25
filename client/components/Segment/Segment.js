// @flow
import React, { Fragment } from 'react';
import AceEditor from 'react-ace';
import 'brace/mode/javascript';
import 'brace/theme/tomorrow';

import type { SegmentType } from 'types/Segment.types';
import Field from 'components/Field';
import * as styles from 'components/Segment/Segment.scss';

type Props = SegmentType & {
    __ID__: number,
    value: ?string,
    onLabelChange: () => void,
    onSegmentChange: () => void,
    isEditor: boolean,
};

const Segment = (props: Props) => {
    const {
        __ID__,
        id,
        isEditor,

        label,
        content,
        //createdAt, TODO: use `createdAt`
        value,

        onLabelChange,
        onSegmentChange,
    } = props;

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

    const renderEditor = () => (
        <Fragment>
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
        </Fragment>
    );

    const renderStatic = () => (
        <Fragment>
            {label && <h3>{label}</h3>}
            <div className={styles.segment}>
                <AceEditor
                    {...editorProps}
                    className={styles.static}
                    readOnly
                    highlightActiveLine={false}
                    value={content}
                    name={`SEGMENT__${id}`}
                />
            </div>
        </Fragment>
    );

    return (
        <div className={styles.segmentWrapper}>
            {isEditor ? renderEditor() : renderStatic()}
        </div>
    );
};

export default Segment;
