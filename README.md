# dnd_pregen
DnD character generator

---------------------------------------------------------------------------------------------------------------------
Req: \
1- docker MongoDB 6 + DB dump + creds.env in folder ../db/ \
2- golang 1.20+

Build: \
cd ../cmd/ \
CGO_ENABLED=0 go build && ./cmd
