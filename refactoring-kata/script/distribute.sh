set -e 
cd "$(dirname $0)/.."

DEST=/tmp/birthday-greetings.zip

rm $DEST 2> /dev/null || true
mvn clean
cd ..
cp -rp refactoring-kata birthday-greetings
zip -r $DEST birthday-greetings
rm -rf birthday-greetings
ls -l $DEST


