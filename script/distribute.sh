set -e
cd "$(dirname $0)/.."

DEST=/tmp/birthday-greetings.zip

rm $DEST 2> /dev/null || true

cd ..
zip -r $DEST birthday-greetings -x '*/lib*' -x '*.DS_Store*' -x '*target*' -x '*bin*' -x '*.git*'
echo
echo "See your distribution file here:"
ls -l $DEST


