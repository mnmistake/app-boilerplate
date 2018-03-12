import React from 'react';

const Segment = ({ id, label, content, createdAt }) => (
    <div className="gridItem">
        <h1>{label}</h1>
        <p>{content}</p>
    </div>
);

export default Segment;
