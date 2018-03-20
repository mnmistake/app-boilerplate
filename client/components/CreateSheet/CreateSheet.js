import React from 'react';
import classNames from 'classnames';
import { graphql } from 'react-apollo';

import history from '../../history';
import * as styles from './CreateSheet.scss';
import Field from '../Field/renderField';
import ToggleButton from '../ToggleButton';
import SegmentCreator from './SegmentCreator';

import { createSheet } from '../../graphql/mutations/Sheets.graphql';

@graphql(createSheet, {
    props: ({ mutate }) => ({
        createSheet: ({ name, segments }) => mutate({ variables: { name, segments } }),
    }),
})
export default class CreateSheet extends React.PureComponent {
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

    setContent = (id, content) => {
        console.log(this.segmentCreator)
        this.setField(id, content, 'content');
    };
    setLabel = (id, label) => this.setField(id, label, 'label');

    addSegmentCreator = () => this.setState({
        segmentCreators: [
            ...this.state.segmentCreators,
            { __ID__: this.state.segmentCreators.pop().__ID__ + 1 },
        ],
    });

    createSheet = async (e) => {
        e.preventDefault();

        const { segments, name } = this.state;
        const { createSheet } = this.props;
        const mappedSegments = segments.map(s => ({ label: s.label, content: s.content }));

        if (mappedSegments.length) { // TODO: errors when no segments are made.
            try {
                const res = await createSheet({ name, segments: mappedSegments });
                const { id } = res.data.createSheet;

                if (res.data) {
                    history.push(`/sheet/${id}`);
                }
            } catch (err) {
                console.error(err);
            }
        }
    };

    render() {
        const { segmentCreators } = this.state;

        return (
            <form className={classNames('container', styles.createSheetWrapper)} onSubmit={e => this.createSheet(e)}>
                <div className={styles.header}>
                    <Field
                        required
                        autoFocus
                        large
                        type="text"
                        name="name"
                        placeholder="Sheet name"
                        onChange={e => this.setState({ name: e.target.value })}
                    />
                    <div className={styles.toggle}>
                        <span className="note">Private</span>
                        <ToggleButton />
                    </div>
                </div>
                <div className="segmentsWrapper">
                    {segmentCreators && segmentCreators.map(s => (
                        <SegmentCreator
                            ref={x => this.segmentCreator = x}
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
                        className="fullWidthBtn"
                    >
                        Create sheet
                    </button>
                </div>
            </form>
        );
    }
}