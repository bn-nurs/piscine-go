url=https://raw.githubusercontent.com/kigiri/superhero-api/master/api/all.json

curl -s ${url} | jq ' .[] | select( .id == '$HERO_ID') | .connections .relatives' | cut -d '"' -f2
