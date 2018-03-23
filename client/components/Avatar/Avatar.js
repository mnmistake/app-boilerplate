// @flow
import React from 'react';

import * as styles from './Avatar.scss';

type Props = {
    username: string,
    size: ?string,
};

const Avatar = ({ username, size }: Props) => {
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
