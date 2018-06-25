// @flow
import React, { PureComponent } from 'react';
import classNames from 'classnames';
import { graphql } from 'react-apollo';

import * as styles from 'components/CreateSheet/CreateSheet.scss';
import Field from 'components/Field';
import ToggleButton from 'components/ToggleButton';
import SegmentCreator from 'components/CreateSheet/SegmentCreator';
import { createSheet } from 'graphql/mutations/Sheets.graphql';

type InternalSegmentType = {
    label: string,
    content: string,
};

type Args = {
    name: string,
    segments: Array<InternalSegmentType>,
};

type Props = {
    createSheet: Args => Object,
};

type State = {
    segments: Array<Object>,
    name: string,
};

@graphql(createSheet, {
    props: ({ mutate }) => ({
        createSheet: ({ name, segments }: Args) => (
            mutate({ variables: { name, segments } })
        ),
    }),
})
export default class CreateSheet extends PureComponent<Props, State> {
    state = {
        // __ID__ is used solely for React element keys.
        // This does not represent the actual ID of the segment.
        segments: [{ __ID__: 0, label: '', content: '' }],
        name: '',
    };

    setField = (__ID__: number, field: string, key: string) => {
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

    setContent = (id: number, content: string) => this.setField(id, content, 'content');
    setLabel = (id: number, label: string) => this.setField(id, label, 'label');

    addSegment = () => this.setState(prevState => ({
        segments: [...prevState.segments, { __ID__: this.state.segments.pop().__ID__ + 1 }]
    }));

    createSheet = async (e: Event) => {
        e.preventDefault();

        const { segments, name } = this.state;
        const { createSheet } = this.props;
        const mappedSegments = segments.map(s => ({ label: s.label, content: s.content }));

        if (mappedSegments.length) { // TODO: errors when no segments are made.
            try {
                const res = await createSheet({ name, segments: mappedSegments });
                const { id } = res.data.createSheet;

                if (res.data) {
                    this.props.history.push(`/sheet/${id}`);
                }
            } catch (err) {
                console.error(err);
            }
        }
    };

    render() {
        const { segments } = this.state;

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
                    {segments && segments.map(s => (
                        <SegmentCreator
                            key={s.__ID__}
                            __ID__={s.__ID__}
                            // $FlowFixMe
                            value={s.content}
                            setContent={this.setContent}
                            setLabel={this.setLabel}
                        />
                    ))}
                    <div className={styles.addSegmentBtnWrapper}>
                        <button type="button" className="circleBtn" onClick={this.addSegment}>
                            +
                        </button>
                    </div>
                    <button
                        style={{ marginTop: '16px' }}
                        className={classNames('fullWidthBtn', styles.createSheetBtn)}
                    >
                        Create sheet
                    </button>
                </div>
            </form>
        );
    }
}
