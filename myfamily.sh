allRelatives=$(curl -s https://learn.zone01oujda.ma/assets/superhero/all.json | jq -r ".[] | select(.id==$HERO_ID) | .connections.relatives ")

formattedRelatives=$(echo $allRelatives | sed 's/"//g')

echo $formattedRelatives