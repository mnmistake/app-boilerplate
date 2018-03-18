import React from 'react';

import * as styles from './Avatar.scss';

const Avatar = ({ username, size }) => {
    const avatarSize = {
        width: size || '35px',
        height: size || '35px',
    };

    return (
        <div className={styles.avatar} style={avatarSize}>
            <span>{username[0].toUpperCase()}</span>
        </div>
    );
};

export default Avatar;