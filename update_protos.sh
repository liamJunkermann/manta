#!/bin/bash
DIR=dota
rm -rf $DIR
svn export https://github.com/SteamDatabase/GameTracking-Dota2.git/trunk/Protobufs $DIR
# removing files from update protobufs
rm -rf $DIR/gametoolevents.proto $DIR/dota_messages_mlbot.proto $DIR/dota_gcmessages_common_bot_script.proto $DIR/steammessages_base.proto $DIR/steammessages_clientserver_login.proto $DIR/*steamworks*.proto $DIR/tensorflow $DIR/netowkr_connection.proto
for f in $DIR/*.proto; do mv "$f" "$(echo "$f" | sed s/.proto$/_m.proto/)"; done
sed -i '/import "google\/protobuf\//! s/.proto";/_m.proto";/g' $DIR/*.proto
sed -i 's/^\(\s*\)\(optional\|repeated\|required\|extend\)\s*\./\1\2 /' $DIR/*.proto
sed -i 's!^\s*rpc\s*\(\S*\)\s*(\.\([^)]*\))\s*returns\s*(\.\([^)]*\))\s*{!rpc \1 (\2) returns (\3) {!' $DIR/*.proto
sed -i '1isyntax = "proto2";\n\npackage dota;\noption go_package = "github.com/liamJunkermann/manta/dota;dota";\n' $DIR/*.proto
sed -i '/^import "google\/protobuf\/valve_extensions\.proto"/d' $DIR/*.proto
sed -i '/^option (/d' $DIR/*.proto
sed -i 's/\s\[.*\]//g' $DIR/*.proto
protoc -I $DIR --go_out=paths=source_relative:$DIR  $DIR/*.proto