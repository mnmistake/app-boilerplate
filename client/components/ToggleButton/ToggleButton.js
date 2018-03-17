import React from 'react';

import * as styles from './ToggleButton.scss';

export default () => (
    <label className={styles.toggleButton}>
        <input type="checkbox" />
        <i></i>
    </label>
);
