name=$(curl -s https://learn.zone01oujda.ma/assets/superhero/all.json | jq -r '.[] | select(.id ==170 ) | .name')
power=$(curl -s https://learn.zone01oujda.ma/assets/superhero/all.json | jq '.[] | select(.id ==170 ) | .powerstats.power')
gender=$(curl -s https://learn.zone01oujda.ma/assets/superhero/all.json | jq -r '.[] | select(.id ==170 ) | .appearance.gender')
printf "%s\n%s\n%s\n" "$name" "$power" "$gender"