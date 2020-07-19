# lpr

## gen struct via github.com/smallnest/gen

~/go/bin/gen --sqltype=postgres --connstr="host=192.168.5.11 port=5432 user=remote dbname=fzzl password=my032003 sslmode=disable" --database fzzl --module github.com/lekai63/lpr/model --verbose --overwrite --out ~/lpr/ --json --db --api=apis --model=models --gorm  --guregu  --rest --mod --server  --makefile --run-gofmt --host=localhost --port=5430