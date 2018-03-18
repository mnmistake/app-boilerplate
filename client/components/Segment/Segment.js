import React from 'react';
import AceEditor from 'react-ace';

import 'brace/mode/javascript';
import 'brace/theme/tomorrow';

import Field from '../Field/renderField';
import * as styles from './Segment.scss';

const Segment = props => {
    const { isCreator, label, content, createdAt } = props;
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
                <AceEditor
                    {...editorProps}
                    value={value}
                    defaultValue="// Write something..."
                    onChange={onSegmentChange}
                    name={`SEGMENT__${__ID__}`}
                />
            </React.Fragment>
        );
    };

    const renderStatic = () => (
        <AceEditor
            {...editorProps}
            readOnly
            highlightActiveLine={false}
            value={content}
            name={`SEGMENT__${props.id}`}
        />
    );

    return (
        <div className={styles.segment}>
            {isCreator ? renderEditor() : renderStatic()}
        </div>
    );
};

export default Segment;
