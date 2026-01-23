PORT_FROM=2002
PORT_TO=22
DEST=172.24.4.16
iptables -t nat -A PREROUTING  -p tcp --dport $PORT_FROM -j DNAT --to $DEST:$PORT_TO
iptables -t nat -A POSTROUTING -p tcp -d $DEST --dport $PORT_TO -j MASQUERADE