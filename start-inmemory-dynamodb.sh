#!/bin/sh

if ! type /usr/local/bin/dynamodb-local &>/dev/null; then
    brew cask install dynamodb-local
fi

exec java -Djava.library.path=/usr/local/Caskroom/dynamodb-local/latest/DynamoDBLocal_lib \
          -jar /usr/local/Caskroom/dynamodb-local/latest/DynamoDBLocal.jar \
          -delayTransientStatuses \
          -inMemory \
          -port 8123 \
          -sharedDb
