import React from 'react';

import * as styles from './Spinner.scss';

const Spinner = () => (
    <div className={styles.spinner}>
        <div className={styles.bounce1} />
        <div className={styles.bounce2} />
        <div className={styles.bounce3} />
    </div>
);

export default Spinner;