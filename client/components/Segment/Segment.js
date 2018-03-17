import React from 'react';
import AceEditor from 'react-ace';

import 'brace/mode/javascript';
import 'brace/theme/tomorrow';

import Field from '../Field/renderField';
import * as styles from './Segment.scss';

const Segment = ({ __ID__, label, value, content, createdAt, isCreator, onSegmentChange, onLabelChange }) => (
    <div className={styles.segment}>
        <Field
            type="text"
            name="label"
            placeholder="Label"
            onChange={onLabelChange}
        />
        <AceEditor
            value={value}
            defaultValue="// Write something..."
            wrapEnabled
            onChange={onSegmentChange}
            width="100%"
            height="300px"
            fontSize={14}
            showPrintMargin={false}
            showGutter={false}
            mode="javascript"
            theme="tomorrow"
            name={`SEGMENT__${__ID__}`}
            editorProps={{ $blockScrolling: true }}
        />
    </div>
);

export default Segment;
