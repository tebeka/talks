#!/bin/bash
# Push application to AppEngine and update deployment tag

# Exit on error
set -e

/opt/google_appengine_go/appcfg.py update .
hg tag -f appengine
hg push
