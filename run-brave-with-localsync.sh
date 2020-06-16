#!/usr/bin/env bash
# https://peter.sh/experiments/chromium-command-line-switches/

THIS_DIR="$(realpath "$(dirname "${BASH_SOURCE[0]}")")"
BROWSER_NUM=1
SYNC_URL="http://localhost:8295"

if ! type multiview &>/dev/null; then
    npm i -g multiview
fi

multiview \
    [ \
        "/Applications/Brave Browser Nightly.app/Contents/MacOS/Brave Browser Nightly" \
            --sync-url="$SYNC_URL" \
            --profile-directory="$THIS_DIR/.local/browser-profile-$(( BROWSER_NUM++ ))" \
            --managed-user-id=$BROWSER_NUM \
            --suppress-message-center-popups \
            --log-level=0 \
    ] [ \
        "/Applications/Brave Browser Nightly.app/Contents/MacOS/Brave Browser Nightly" \
            --sync-url="$SYNC_URL" \
            --profile-directory="$THIS_DIR/.local/browser-profile-$(( BROWSER_NUM++ ))" \
            --managed-user-id=$BROWSER_NUM \
            --suppress-message-center-popups \
            --log-level=0 \
    ] \
    --autoexit

echo "Done."
exit 0
