HERO_ID=1
relatives=$(curl -s https://learn.zone01oujda.ma/assets/superhero/all.json | jq -r ".[] | select(.id==$HERO_ID) | .connections.relatives")
echo $relatives