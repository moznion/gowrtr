#!/bin/bash

CURRENT_PATH="$(cd "$(dirname "$0")" || exit; pwd)"
PATH="$CURRENT_PATH/bin:$PATH"

(cd "$CURRENT_PATH/.." && go generate ./...)

