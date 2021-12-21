s/---.*---/{beacons: []*beacon{/g
s/^$/}},/g
s/\(^-\?[0-9].*\)/\t{c:coord{\1},oc:coord{\1}},/g
