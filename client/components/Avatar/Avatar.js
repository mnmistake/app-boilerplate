import React from 'react';

import * as styles from './Avatar.scss';

const Avatar = ({ username }) => (
    <div className={styles.avatar}>
        <span>{username[0].toUpperCase()}</span>
    </div>
);

export default Avatar;