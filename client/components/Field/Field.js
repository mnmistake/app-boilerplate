import React from 'react';
import classNames from 'classnames';

const Field = ({ input, name, placeholder, type, required, large, onChange, onBlur, autoFocus }) => (
    <div className={classNames('inputWrapper', large && 'large')}>
        <input
            {...input}
            type={type}
            onChange={onChange}
            onBlur={onBlur}
            className="input"
            required={required}
            autoFocus={autoFocus}
        />
        <label htmlFor={name}>{placeholder}</label>
        <span className="inputBar" />
    </div>
);

export default Field;
