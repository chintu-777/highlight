import React, { useEffect } from 'react';

import Button from '../../../../components/Button/Button/Button';
import ElevatedCard from '../../../../components/ElevatedCard/ElevatedCard';
import SvgSearchIcon from '../../../../static/SearchIcon';
import { usePlayerUIContext } from '../../context/PlayerUIContext';
import usePlayerConfiguration from '../../PlayerHook/utils/usePlayerConfiguration';
import styles from './NoActiveSessionCard.module.scss';

const NoActiveSessionCard = () => {
    const { setShowLeftPanel } = usePlayerConfiguration();
    const { searchBarRef } = usePlayerUIContext();

    useEffect(() => {
        setShowLeftPanel(true);
    }, [setShowLeftPanel]);

    return (
        <ElevatedCard
            className={styles.card}
            title="Ready to see what's happening in your app?"
            actions={
                <Button
                    trackingId="NoActiveSessionCardPerformASearch"
                    type="primary"
                    className={styles.buttonWrapper}
                    onClick={() => {
                        if (searchBarRef) {
                            searchBarRef.focus();
                        }
                    }}
                >
                    <SvgSearchIcon />
                    Perform a Search
                </Button>
            }
        >
            <p>
                View a recent session or find a specific user, event, or
                segment.
            </p>
        </ElevatedCard>
    );
};

export default NoActiveSessionCard;
