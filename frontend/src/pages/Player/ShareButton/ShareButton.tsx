import React, { useContext } from 'react';
import PrimaryButton from '../../../components/Button/PrimaryButton/PrimaryButton';
import ReplayerContext from '../ReplayerContext';
import { message } from 'antd';
import { PlayerSearchParameters } from '../PlayerHook/utils';
import styles from './ShareButton.module.scss';

const ShareButton = () => {
    const { time } = useContext(ReplayerContext);

    /**
     * Copies the current session URL with a search parameter "ts" with the player's current time in seconds.
     */
    const onGetLinkWithTimestamp = () => {
        const currentUrl = new URL(window.location.href);
        currentUrl.searchParams.set(
            PlayerSearchParameters.ts,
            (time / 1000).toString()
        );

        message.success('Copied link!');
        navigator.clipboard.writeText(currentUrl.href);
    };
    return (
        <PrimaryButton
            onClick={onGetLinkWithTimestamp}
            className={styles.shareButton}
        >
            Share
        </PrimaryButton>
    );
};

export default ShareButton;
