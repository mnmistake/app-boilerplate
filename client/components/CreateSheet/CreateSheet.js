import React from 'react';
import classNames from 'classnames';
import { graphql } from 'react-apollo';

import * as styles from './CreateSheet.scss';
import Field from '../Field/renderField';
import ToggleButton from '../ToggleButton';
import SegmentCreator from './SegmentCreator';

import userQuery from '../../graphql/queries/user.graphql';
import { createSheet } from '../../graphql/mutations/sheets.graphql';

@graphql(createSheet, {
    props: ({ mutate }) => ({
        createSheet: ({ userId, name, segments }) => mutate({ variables: { userId, name, segments } }),
    }),
})
@graphql(userQuery, {
    props: ({ data: { user } }) => ({
        user,
    }),
})
export default class CreateSheet extends React.Component {
    state = {
        // __ID__ is used solely for React element keys.
        // This does not represent the actual ID of the segment.
        segmentCreators: [{ __ID__: 0 }],
        segments: [],
        name: '',
    };

    setField = (__ID__, field, key) => {
        const segment = this.state.segments.find(s => s.__ID__ === __ID__);

        if (segment) {
            const segments = [...this.state.segments].map(s => (
                s.__ID__ === __ID__ ?
                    {
                        ...s,
                        [key]: field,
                    } : s
            ));

            this.setState({ segments });
        } else {
            this.setState({
                segments: [...this.state.segments, {
                    __ID__,
                    [key]: field,
                }],
            });
        }
    };

    setContent = (id, content) => this.setField(id, content, 'content');
    setLabel = (id, label) => this.setField(id, label, 'label');

    addSegmentCreator = () => this.setState({
        segmentCreators: [
            ...this.state.segmentCreators,
            { __ID__: this.state.segmentCreators.pop().__ID__ + 1 },
        ],
    });

    createSheet = () => {
        const { segments, name } = this.state;
        const { createSheet, user: { id: userId } } = this.props;

        if (segments) {
            createSheet({ userId, name, segments });
        }
    };

    render() {
        const { segmentCreators } = this.state;

        return (
            <form className={classNames('container', styles.createSheetWrapper)}>
                <div className={styles.header}>
                    <Field
                        required
                        autoFocus
                        large
                        type="text"
                        name="name"
                        placeholder="Name your sheet..."
                        onLabelChange={e => this.setState({ name: e.target.value })}
                    />
                    <div className={styles.toggle}>
                        <span className="note">Private</span>
                        <ToggleButton />
                    </div>
                </div>
                <div className={styles.segmentCreators}>
                    {segmentCreators && segmentCreators.map(s => (
                        <SegmentCreator
                            key={s.__ID__}
                            __ID__={s.__ID__}
                            value={s.content}
                            setContent={this.setContent}
                            setLabel={this.setLabel}
                        />
                    ))}
                    <div className={styles.addSegmentBtnWrapper}>
                        <button type="button" className="circleBtn" onClick={this.addSegmentCreator}>
                            +
                        </button>
                    </div>
                    <button
                        style={{ marginTop: '16px' }}
                        type="button"
                        className="fullWidthBtn"
                        onClick={this.createSheet}
                    >
                        Create sheet
                    </button>
                </div>
            </form>
        );
    }
}