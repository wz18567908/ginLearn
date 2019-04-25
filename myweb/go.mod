module ginLearn/myweb

go 1.12

require (
	github.com/alecthomas/template v0.0.0-20160405071501-a0175ee3bccc
	github.com/gin-gonic/gin v1.3.0
	github.com/jinzhu/gorm v1.9.4
	github.com/swaggo/gin-swagger v1.1.0
	github.com/swaggo/swag v1.5.0
)

replace (
	golang.org/x/sys v0.0.0-20190215142949-d0b11bdaac8a => github.com/golang/sys v0.0.0-20190215142949-d0b11bdaac8a
	golang.org/x/sys v0.0.0-20190222072716-a9d3bda3a223 => github.com/golang/sys v0.0.0-20190222072716-a9d3bda3a223
	golang.org/x/sys v0.0.0-20190322080309-f49334f85ddc => github.com/golang/sys v0.0.0-20190322080309-f49334f85ddc
)

replace golang.org/x/text v0.3.0 => github.com/golang/text v0.3.0

replace (
	golang.org/x/tools v0.0.0-20190110015856-aa033095749b => github.com/golang/tools v0.0.0-20190110015856-aa033095749b
	golang.org/x/tools v0.0.0-20190322203728-c1a832b0ad89 => github.com/golang/tools v0.0.0-20190322203728-c1a832b0ad89
)

replace (
	golang.org/x/net v0.0.0-20181005035420-146acd28ed58 => github.com/golang/net v0.0.0-20181005035420-146acd28ed58
	golang.org/x/net v0.0.0-20190108225652-1e06a53dbb7e => github.com/golang/net v0.0.0-20190108225652-1e06a53dbb7e
	golang.org/x/net v0.0.0-20190311183353-d8887717615a => github.com/golang/net v0.0.0-20190311183353-d8887717615a
	golang.org/x/net v0.0.0-20190322120337-addf6b3196f6 => github.com/golang/net v0.0.0-20190322120337-addf6b3196f6
)

replace (
	golang.org/x/crypto v0.0.0-20190308221718-c2843e01d9a2 => github.com/golang/crypto v0.0.0-20190308221718-c2843e01d9a2
	golang.org/x/crypto v0.0.0-20190325154230-a5d413f7728c => github.com/golang/crypto v0.0.0-20190325154230-a5d413f7728c
)

replace cloud.google.com/go v0.37.2 => github.com/golang/go v0.0.0-20190425030032-44343c777ca8
