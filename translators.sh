#!/bin/bash
# Translators script for umotd

if [ -z "$1" ]; then
    echo "Run the command like this: \`$0 <your-language-code>\`"
    echo ""
    echo "> Example: \`$0 fr\`"
    exit 1
fi

if [ -d "locales/$1" ]; then
    echo "Language $1 already exists. Updating..."
    ~/go/bin/xgotext -in . -out locales/temp
    msgmerge --update locales/$1/LC_MESSAGES/default.po locales/temp/default.pot
    rm -rf locales/temp
    rm -f locales/$1/LC_MESSAGES/default.po~
    echo "Translations for $1 updated!"
    exit 0
fi

echo "Creating new language $1..."
mkdir -p locales/$1/LC_MESSAGES
~/go/bin/xgotext -in . -out locales/temp
cp locales/temp/default.pot locales/$1/LC_MESSAGES/default.po
rm -rf locales/temp
echo "Translations for $1 generated. Edit locales/$1/LC_MESSAGES/default.po"
